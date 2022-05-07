/*
 * Copyright (c) 2019, Baidu XPU Authors.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

var passDeviceSpecs = flag.Bool("pass-device-specs", true, "pass the list of DeviceSpecs to the kubelet on Allocate()")
var renameDev = flag.Bool("rename-dev", false, "rename dev ?")

// XPUDevicePlugin implements the Kubernetes device plugin API
type XPUDevicePlugin struct {
	ResourceManager
	resourceName   string
	allocateEnvvar string
	socket         string

	server        *grpc.Server
	cachedDevices []*Device
	health        chan *Device
	stop          chan interface{}
}

// NewXPUDevicePlugin returns an initialized XPUDevicePlugin
func NewXPUDevicePlugin(resourceName string, resourceManager ResourceManager, allocateEnvvar string, socket string) *XPUDevicePlugin {
	return &XPUDevicePlugin{
		ResourceManager: resourceManager,
		resourceName:    resourceName,
		allocateEnvvar:  allocateEnvvar,
		socket:          socket,

		// These will be reinitialized every
		// time the plugin server is restarted.
		cachedDevices: nil,
		server:        nil,
		health:        nil,
		stop:          nil,
	}
}

func (m *XPUDevicePlugin) initialize() {
	m.cachedDevices = m.Devices()
	m.server = grpc.NewServer([]grpc.ServerOption{}...)
	m.health = make(chan *Device)
	m.stop = make(chan interface{})
}

func (m *XPUDevicePlugin) cleanup() {
	close(m.stop)
	m.cachedDevices = nil
	m.server = nil
	m.health = nil
	m.stop = nil
}

// Start starts the gRPC server, registers the device plugin with the Kubelet,
// and starts the device healthchecks.
func (m *XPUDevicePlugin) Start() error {
	//Delete the existing socket file when start
	if err := os.Remove(m.socket); err != nil && !os.IsNotExist(err) {
		log.Printf("Could not clean up when start device plugin for '%s': %s", m.resourceName, err)
		return err
	}
	//Make sure the directory exists
	if err := os.MkdirAll(pluginapi.DevicePluginPath, os.ModePerm); err != nil {
		log.Printf("Could not initialize device plugin path when start device plugin for '%s': %s", m.resourceName, err)
		return err
	}
	m.initialize()

	err := m.Serve()
	if err != nil {
		log.Printf("Could not start device plugin for '%s': %s", m.resourceName, err)
		m.cleanup()
		return err
	}
	log.Printf("Starting to serve '%s' on %s", m.resourceName, m.socket)

	err = m.Register()
	if err != nil {
		log.Printf("Could not register device plugin: %s", err)
		m.Stop()
		return err
	}
	log.Printf("Registered device plugin for '%s' with Kubelet", m.resourceName)

	go m.CheckHealth(m.stop, m.cachedDevices, m.health)

	return nil
}

// Stop stops the gRPC server.
func (m *XPUDevicePlugin) Stop() error {
	if m == nil || m.server == nil {
		return nil
	}
	log.Printf("Stopping to serve '%s' on %s", m.resourceName, m.socket)
	m.server.Stop()
	if err := os.Remove(m.socket); err != nil && !os.IsNotExist(err) {
		return err
	}
	m.cleanup()
	return nil
}

// Serve starts the gRPC server of the device plugin.
func (m *XPUDevicePlugin) Serve() error {
	sock, err := net.Listen("unix", m.socket)
	if err != nil {
		return err
	}

	pluginapi.RegisterDevicePluginServer(m.server, m)

	go func() {
		lastCrashTime := time.Now()
		restartCount := 0
		for {
			log.Printf("Starting GRPC server for '%s'", m.resourceName)
			err := m.server.Serve(sock)
			if err == nil {
				break
			}

			log.Printf("GRPC server for '%s' crashed with error: %v", m.resourceName, err)

			// restart if it has not been too often
			// i.e. if server has crashed more than 5 times and it didn't last more than one hour each time
			if restartCount > 5 {
				// quit
				log.Fatal("GRPC server for '%s' has repeatedly crashed recently. Quitting", m.resourceName)
			}
			timeSinceLastCrash := time.Since(lastCrashTime).Seconds()
			lastCrashTime = time.Now()
			if timeSinceLastCrash > 3600 {
				// it has been one hour since the last crash.. reset the count
				// to reflect on the frequency
				restartCount = 1
			} else {
				restartCount += 1
			}
		}
	}()

	// Wait for server to start by launching a blocking connexion
	conn, err := m.dial(m.socket, 5*time.Second)
	if err != nil {
		return err
	}
	conn.Close()

	return nil
}

// Register registers the device plugin for the given resourceName with Kubelet.
func (m *XPUDevicePlugin) Register() error {
	conn, err := m.dial(pluginapi.KubeletSocket, 5*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pluginapi.NewRegistrationClient(conn)
	reqt := &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		Endpoint:     path.Base(m.socket),
		ResourceName: m.resourceName,
	}

	_, err = client.Register(context.Background(), reqt)
	if err != nil {
		return err
	}
	return nil
}

func (m *XPUDevicePlugin) GetDevicePluginOptions(context.Context, *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	return &pluginapi.DevicePluginOptions{}, nil
}

// ListAndWatch lists devices and update that list according to the health status
func (m *XPUDevicePlugin) ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
	s.Send(&pluginapi.ListAndWatchResponse{Devices: m.apiDevices()})

	for {
		select {
		case <-m.stop:
			return nil
		case d := <-m.health:
			// FIXME: there is no way to recover from the Unhealthy state.
			if d.Health == pluginapi.Unhealthy {
				log.Printf("'%s' device marked unhealthy: %s", m.resourceName, d.ID)
			} else if d.Health == pluginapi.Healthy {
				log.Printf("'%s' device marked healthy: %s", m.resourceName, d.ID)
			}
			s.Send(&pluginapi.ListAndWatchResponse{Devices: m.apiDevices()})
		}
	}
}

// Allocate which return list of devices.
func (m *XPUDevicePlugin) Allocate(ctx context.Context, reqs *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	responses := pluginapi.AllocateResponse{}
	for _, req := range reqs.ContainerRequests {
		for _, id := range req.DevicesIDs {
			if !m.deviceExists(id) {
				return nil, fmt.Errorf("invalid allocation request for '%s': unknown device: %s", m.resourceName, id)
			}
		}

		response := pluginapi.ContainerAllocateResponse{
			/*
				Envs: map[string]string{
					m.allocateEnvvar: strings.Join(req.DevicesIDs, ","),
				},
			*/
		}
		if *passDeviceSpecs {
			response.Devices = m.apiDeviceSpecs(req.DevicesIDs)
		}

		responses.ContainerResponses = append(responses.ContainerResponses, &response)
	}

	return &responses, nil
}

func (m *XPUDevicePlugin) PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return &pluginapi.PreStartContainerResponse{}, nil
}

// dial establishes the gRPC communication with the registered device plugin.
func (m *XPUDevicePlugin) dial(unixSocketPath string, timeout time.Duration) (*grpc.ClientConn, error) {
	c, err := grpc.Dial(unixSocketPath, grpc.WithInsecure(), grpc.WithBlock(),
		grpc.WithTimeout(timeout),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}),
	)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (m *XPUDevicePlugin) deviceExists(id string) bool {
	for _, d := range m.cachedDevices {
		if d.ID == id {
			return true
		}
	}
	return false
}

func (m *XPUDevicePlugin) apiDevices() []*pluginapi.Device {
	var pdevs []*pluginapi.Device
	for _, d := range m.cachedDevices {
		pdevs = append(pdevs, &d.Device)
	}
	return pdevs
}

func (m *XPUDevicePlugin) apiDeviceSpecs(filter []string) []*pluginapi.DeviceSpec {
	var specs []*pluginapi.DeviceSpec

	test_path_prefix := "/dev"
	host_path_prefix := "/dev"
	container_path_prefix := "/dev"
	if *useFakeDev {
		log.Printf("Warning: useFakeDev")
		test_path_prefix = "/fake_dev"
		host_path_prefix = "/home/miaotianxiang/fake_dev"
	}

	paths := []string{
		"xpuctrl",
	}

	for _, p := range paths {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", test_path_prefix, p)); err == nil {
			spec := &pluginapi.DeviceSpec{
				ContainerPath: fmt.Sprintf("%s/%s", container_path_prefix, p),
				HostPath:      fmt.Sprintf("%s/%s", host_path_prefix, p),
				Permissions:   "rw",
			}
			specs = append(specs, spec)
		}
	}

	for _, d := range m.cachedDevices {
		for idx, id := range filter {
			if d.ID == id {
				containerPath := d.ContainerPath
				if *renameDev {
					// rename to /dev/xpu0, /dev/xpu1 ...
					// renumber from 0
					containerPath = fmt.Sprintf("/dev/xpu%d", idx)
				}
				spec := &pluginapi.DeviceSpec{
					ContainerPath: containerPath,
					HostPath:      d.HostPath,
					Permissions:   "rw",
				}
				specs = append(specs, spec)
			}
		}
	}

	return specs
}
