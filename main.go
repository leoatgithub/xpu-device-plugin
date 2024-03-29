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
	"log"
	"os"
	"syscall"
	"fmt"

	"github.com/fsnotify/fsnotify"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"github.com/kunlunxin/go-xpuml/pkg/xpuml"
)

func getAllPlugins() []*XPUDevicePlugin {
	return []*XPUDevicePlugin{
		NewXPUDevicePlugin(
			"baidu.com/xpu",
			NewXPUDeviceManager(),
			"", //"XPU_VISIBLE_DEVICES"
			pluginapi.DevicePluginPath+"xpu.sock"),
	}
}

func main() {
	flag.Parse()

	log.Println("Loading XPUML")
	ret := xpuml.Init()
	if ret != xpuml.SUCCESS {
		log.Fatalf("Failed to initialize XPUML: %v", xpuml.ErrorString(ret))
	}
	defer func() { log.Println("Shutdown of XPUML returned:", xpuml.Shutdown()) }()

	log.Println("Starting FS watcher.")
	watcher, err := newFSWatcher(pluginapi.DevicePluginPath)
	if err != nil {
		log.Println("Failed to created FS watcher.")
		os.Exit(1)
	}
	defer watcher.Close()

	log.Println("Starting OS watcher.")
	sigs := newOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println("Retreiving plugins.")
	plugins := getAllPlugins()

restart:
	// Loop through all plugins, idempotently stopping them, and then starting
	// them if they have any devices to serve. If even one plugin fails to
	// start properly, try starting them all again.
	started := 0
	pluginStartError := make(chan struct{})
	for _, p := range plugins {
		p.Stop()

		// Just continue if there are no devices to serve for plugin p.
		devs := p.Devices()
		if len(devs) == 0 {
			continue
		}

		// Update ResourceName by model
		if *useDetailResName {
			model := devs[0].Model
			for _, dev := range devs {
				// hybrid model(K200 or R200) or type(PF or VF) between xpus on one host is not support now, goto restart
				if model != dev.Model {
					log.Fatalf("All devs on node must be the same model or VF spec, but get: %s and %s", model, dev.Model)
					close(pluginStartError)
					goto events
				}
			}
			p.resourceName = fmt.Sprintf("baidu.com/%s", model)
		}

		// Start the gRPC server for plugin p and connect it with the kubelet.
		if err := p.Start(); err != nil {
			log.Println("Could not contact Kubelet, retrying. Did you enable the device plugin feature gate?")
			log.Printf("You can check the prerequisites at: https://github.com/NVIDIA/k8s-device-plugin#prerequisites")
			log.Printf("You can learn how to set the runtime at: https://github.com/NVIDIA/k8s-device-plugin#quick-start")
			close(pluginStartError)
			goto events
		}
		started++
	}

	if started == 0 {
		log.Println("No devices found. Waiting indefinitely.")
	}

events:
	// Start an infinite loop, waiting for several indicators to either log
	// some messages, trigger a restart of the plugins, or exit the program.
	for {
		select {
		// If there was an error starting any plugins, restart them all.
		case <-pluginStartError:
			goto restart

		// Detect a kubelet restart by watching for a newly created
		// 'pluginapi.KubeletSocket' file. When this occurs, restart this loop,
		// restarting all of the plugins in the process.
		case event := <-watcher.Events:
			if event.Name == pluginapi.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
				log.Printf("inotify: %s created, restarting.", pluginapi.KubeletSocket)
				goto restart
			}

		// Watch for any other fs errors and log them.
		case err := <-watcher.Errors:
			log.Printf("inotify: %s", err)

		// Watch for any signals from the OS. On SIGHUP, restart this loop,
		// restarting all of the plugins in the process. On all other
		// signals, exit the loop and exit the program.
		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				log.Println("Received SIGHUP, restarting.")
				goto restart
			default:
				log.Printf("Received signal \"%v\", shutting down.", s)
				for _, p := range plugins {
					p.Stop()
				}
				break events
			}
		}
	}
}
