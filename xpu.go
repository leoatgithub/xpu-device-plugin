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
	"os"
	"strconv"
	"strings"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"github.com/kunlunxin/go-xpuml/pkg/xpuml"
)

var useFakeDev = flag.Bool("use-fake-dev", false, "use /fake_dev ?")
var vfNum = flag.Int("sriov-num-vfs", 99, "SR-IOV virtual function number per xpu")
var useDetailResName = flag.Bool("use-detail-resName", false, "register 'ResourceName' base on xpu-model or default: baidu.com/xpu")

const (
	envDisableHealthChecks = "DP_DISABLE_HEALTHCHECKS"
	envNodeName            = "NODE_NAME"
	allHealthChecks        = "xids"
	labelSriovVfNum        = "baidu.com/sriov-num-vfs"
)

const (
    XPUML_DEVICE_MODEL_UNKNOWN                        = 9999

    XPUML_DEVICE_MODEL_KL1_K200                       = 0
    XPUML_DEVICE_MODEL_KL1_K100                       = 1

    XPUML_DEVICE_MODEL_KL2_R200                       = 2
    XPUML_DEVICE_MODEL_KL2_R300                       = 3
    XPUML_DEVICE_MODEL_KL2_R200_8F                    = 4

    XPUML_DEVICE_MODEL_KL2_R200_SRIOV_PF              = 200
    XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_ONE   = 201
    XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_TWO   = 202
    XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_THREE = 203
)

type Device struct {
	pluginapi.Device
	Model         string
	HostPath      string
	ContainerPath string
}

type ResourceManager interface {
	Devices() []*Device
	CheckHealth(stop <-chan interface{}, devices []*Device, unhealthy chan<- *Device)
}

type XPUDeviceManager struct{}

func UpdateVFNumFromLabel() {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		fmt.Errorf("error building kubernetes clientcmd config: %s", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Errorf("error building kubernetes clientset from config: %s", err)
	}

	listWatch := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"nodes",
		v1.NamespaceAll,
		fields.OneTermEqualSelector("metadata.name", os.Getenv(envNodeName)),
	)

	_, controller := cache.NewInformer(
		listWatch, &v1.Node{}, 0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				label := obj.(*v1.Node).Labels[labelSriovVfNum]
				num, err := strconv.ParseInt(label, 10, 32)
				if err != nil {
					log.Printf("Find no availabble config labels, use config in file")
				} else {
					*vfNum = int(num)
					log.Printf("Got vfNum[%d] from label:%v", *vfNum, labelSriovVfNum)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
			},
		},
	)

	stop := make(chan struct{})
	go controller.Run(stop)

	// wait for event handler to add func
	time.Sleep(time.Second)
}

func check(err error) {
	if err != nil {
		log.Panicln("Fatal:", err)
	}
}

func NewXPUDeviceManager() *XPUDeviceManager {
	return &XPUDeviceManager{}
}

func SetVfNum() {

	UpdateVFNumFromLabel()

	count, ret := xpuml.DeviceGetCount()
	if ret != xpuml.SUCCESS {
		log.Fatalf("Failed to get device count: %v", xpuml.ErrorString(ret))
	}
	for i := int(0); i < count; i++ {
		device, ret := xpuml.DeviceGetHandleByIndex(i)
		if ret != xpuml.SUCCESS {
			log.Printf("Failed to get device handle: %v", xpuml.ErrorString(ret))
			continue
		}

		mode, ret1 := xpuml.DeviceGetHostVxpuMode(device)
		if ret1 != xpuml.SUCCESS || mode == xpuml.HOST_VXPU_MODE_NON_SRIOV {
			continue
		}

		// default value is 99, means got no params from label or file, skip VF configuration
		if *vfNum == 99 {
			continue
		}

		err := xpuml.DeviceSetSriovVfNum(device, int32(*vfNum))
		if err == xpuml.SUCCESS {
			time.Sleep(time.Second)
		}
	}
}

func getModelNameByModelId(modelId int32) string {
	model := ""
	switch modelId {
	case XPUML_DEVICE_MODEL_UNKNOWN:
		model = "unknow"
	case XPUML_DEVICE_MODEL_KL1_K100:
		model = "K100"
	case XPUML_DEVICE_MODEL_KL1_K200:
		model = "K200"
	case XPUML_DEVICE_MODEL_KL2_R200:
		model = "R200"
	case XPUML_DEVICE_MODEL_KL2_R300:
		model = "R300"
	case XPUML_DEVICE_MODEL_KL2_R200_8F:
		model = "R200-8F"
	case XPUML_DEVICE_MODEL_KL2_R200_SRIOV_PF:
		model = "R200-PF"
	case XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_ONE:
		model = "R200-VF-16G"
	case XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_TWO:
		model = "R200-VF-8G"
	case XPUML_DEVICE_MODEL_KL2_R200_SRIOV_VF_ONE_OF_THREE:
		model = "R200-VF-5G"
	default:
		model = "unknow"
	}
	return model
}

func (g *XPUDeviceManager) Devices() []*Device {
	test_path_prefix := "/dev"
	host_path_prefix := "/dev"
	container_path_prefix := "/dev"
	if *useFakeDev {
		log.Printf("Warning: useFakeDev")
		test_path_prefix = "/fake_dev"
		host_path_prefix = "/home/miaotianxiang/fake_dev"
	}

	var devs []*Device

	SetVfNum()

	count, ret := xpuml.DeviceGetCount()
	if ret != xpuml.SUCCESS {
		log.Fatalf("Failed to get device count: %v", xpuml.ErrorString(ret))
	}

	for i := int(0); i < count; i++ {
		device, ret := xpuml.DeviceGetHandleByIndex(i)
		if ret != xpuml.SUCCESS {
			log.Printf("Failed to get device handle: %v", xpuml.ErrorString(ret))
			continue
		}
		test_path := fmt.Sprintf("%s/xpu%d", test_path_prefix, i)
		mode, _ := xpuml.DeviceGetHostVxpuMode(device)
		if mode == xpuml.HOST_VXPU_MODE_SRIOV_ON {
			log.Printf("Info: Path %s is PF node", test_path)
		} else {
			attr, _ := xpuml.DeviceGetAttributes(device)
			model := getModelNameByModelId(attr.ModelId)
			if _, err := os.Stat(test_path); err == nil {
				log.Printf("Info: Path %s is OK", test_path)
				host_path := fmt.Sprintf("%s/xpu%d", host_path_prefix, i)
				container_path := fmt.Sprintf("%s/xpu%d", container_path_prefix, i)
				devs = append(devs, buildDevice(uint(i), model, host_path, container_path))
			}
		}
	}

	return devs
}

func (g *XPUDeviceManager) CheckHealth(stop <-chan interface{}, devices []*Device, unhealthy chan<- *Device) {
	checkHealth(stop, devices, unhealthy)
}

func buildDevice(i uint, model string, hp string, cp string) *Device {
	dev := Device{}
	dev.ID = fmt.Sprintf("%d", i)
	dev.Health = pluginapi.Healthy
	dev.Model = model
	dev.HostPath = hp
	dev.ContainerPath = cp
	return &dev
}

func checkHealth(stop <-chan interface{}, devices []*Device, unhealthy chan<- *Device) {
	disableHealthChecks := strings.ToLower(os.Getenv(envDisableHealthChecks))
	if disableHealthChecks == "all" {
		disableHealthChecks = allHealthChecks
	}
	if strings.Contains(disableHealthChecks, "xids") {
		return
	}

	const (
		xre_tools_path_opt1 = "/usr/local/xpu/tools"
		xre_tools_path_opt2 = "/home/work/xpu/xpu_runtime/tools"
	)

	for {
		select {
		case <-stop:
			return
		default:
		}

		for _, d := range devices {
			id, _ := strconv.Atoi(d.ID)
			device, err := xpuml.DeviceGetHandleByIndex(id)
			if err != xpuml.SUCCESS {
				log.Printf("Failed to get device handle: %v, dev path: %v",
					xpuml.ErrorString(err), d.HostPath)
				d.Health = pluginapi.Unhealthy
			} else {
				state, err := xpuml.DeviceGetState(device)
				if err != xpuml.SUCCESS {
					log.Printf("Failed to get device state: %v, dev path: %v",
						xpuml.ErrorString(err), d.HostPath)
					d.Health = pluginapi.Unhealthy
				} else {
					if state == xpuml.DEVICE_STATE_ERROR {
						log.Printf("Device state error, dev path: %v", d.HostPath)
						d.Health = pluginapi.Unhealthy
					} else {
						d.Health = pluginapi.Healthy
					}
				}
			}

			unhealthy <- d
		}

		time.Sleep(5000 * time.Millisecond)
	}
}
