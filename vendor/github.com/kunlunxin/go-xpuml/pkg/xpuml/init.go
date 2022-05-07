// Copyright (c) 2022, KUNLUNXIN CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xpuml

import (
	"fmt"

	"github.com/kunlunxin/go-xpuml/pkg/dl"
)

import "C"

const (
	xpumlLibraryName      = "libxpuml.so"
	xpumlLibraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

var xpuml *dl.DynamicLibrary

// xpuml.Init()
func Init() Return {
	lib := dl.New(xpumlLibraryName, xpumlLibraryLoadFlags)
	if lib == nil {
		panic(fmt.Sprintf("error instantiating DynamicLibrary for %s", xpumlLibraryName))
	}

	err := lib.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening %s: %v", xpumlLibraryName, err))
	}

	xpuml = lib
	updateVersionedSymbols()

	return xpumlInit()
}

// xpuml.Shutdown()
func Shutdown() Return {
	ret := xpumlShutdown()
	if ret != SUCCESS {
		return ret
	}

	err := xpuml.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing %s: %v", xpumlLibraryName, err))
	}

	return ret
}

// Default all versioned APIs to v1 (to infer the types)
var xpumlInit = xpumlInit_v1
var xpumlDeviceGetCount = xpumlDeviceGetCount_v1
var xpumlDeviceGetAttributes = xpumlDeviceGetAttributes_v1
var xpumlDeviceGetHandleByIndex = xpumlDeviceGetHandleByIndex_v1
var xpumlDeviceGetComputeRunningProcesses = xpumlDeviceGetComputeRunningProcesses_v1
var xpumlDeviceGetState = xpumlDeviceGetState_v1
var xpumlDeviceGetHostVxpuMode = xpumlDeviceGetHostVxpuMode_v1
var xpumlDeviceSetSriovVfNum = xpumlDeviceSetSriovVfNum_v1

// updateVersionedSymbols()
func updateVersionedSymbols() {
	// err := xpuml.Lookup("xpumlInit_v2")
	// if err == nil {
	// 	xpumlInit = xpumlInit_v2
	// }
}
