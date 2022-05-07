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

// xpuml.DeviceGetCount()
func DeviceGetCount() (int, Return) {
	var DeviceCount uint32
	ret := xpumlDeviceGetCount(&DeviceCount)
	return int(DeviceCount), ret
}

// xpuml.DeviceGetHandleByIndex()
func DeviceGetHandleByIndex(Index int) (Device, Return) {
	var Device Device
	ret := xpumlDeviceGetHandleByIndex(uint32(Index), &Device)
	return Device, ret
}

// xpuml.DeviceGetAttributes()
func DeviceGetAttributes(Device Device) (DeviceAttributes, Return) {
	var Attributes DeviceAttributes
	ret := xpumlDeviceGetAttributes(Device, &Attributes)
	return Attributes, ret
}

func (Device Device) GetAttributes() (DeviceAttributes, Return) {
	return DeviceGetAttributes(Device)
}

// xpuml.DeviceGetMemoryInfo()
func DeviceGetMemoryInfo(Device Device) (Memory, Return) {
	var Memory Memory
	ret := xpumlDeviceGetMemoryInfo(Device, &Memory)
	return Memory, ret
}

func (Device Device) GetMemoryInfo() (Memory, Return) {
	return DeviceGetMemoryInfo(Device)
}

// xpuml.DeviceGetBoardId()
func DeviceGetBoardId(Device Device) (uint32, Return) {
	var BoardId uint32
	ret := xpumlDeviceGetBoardId(Device, &BoardId)
	return BoardId, ret
}

func (Device Device) GetBoardId() (uint32, Return) {
	return DeviceGetBoardId(Device)
}

// xpuml.DeviceGetUtilizationRates()
func DeviceGetUtilizationRates(Device Device) (Utilization, Return) {
	var Utilization Utilization
	ret := xpumlDeviceGetUtilizationRates(Device, &Utilization)
	return Utilization, ret
}

func (Device Device) GetUtilizationRates() (Utilization, Return) {
	return DeviceGetUtilizationRates(Device)
}

// xpuml.DeviceGetComputeRunningProcesses()
func DeviceGetComputeRunningProcesses(Device Device) ([]ProcessInfo, Return) {
	var Infos []ProcessInfo  // This is the v2 version of process info data structure
	var InfoCount uint32 = 1 // Will be reduced upon returning
	var ret Return           // Will be changed upon returning
	for {
		Infos = make([]ProcessInfo, InfoCount)
		ret = xpumlDeviceGetComputeRunningProcesses(Device, &InfoCount, &Infos[0]) // Call v2 version directly
		if ret == SUCCESS {
			break
		}
		if ret != ERROR_INSUFFICIENT_SIZE {
			return nil, ret
		}
		InfoCount *= 2
	}

	if InfoCount == 0 {
		return []ProcessInfo{}, SUCCESS
	}
	return Infos[:InfoCount], SUCCESS
}

func (Device Device) GetComputeRunningProcesses() ([]ProcessInfo, Return) {
	return DeviceGetComputeRunningProcesses(Device)
}

// xpuml.DeviceGetState()
func DeviceGetState(Device Device) (DeviceState, Return) {
	var State DeviceState
	ret := xpumlDeviceGetState(Device, &State)
	return State, ret
}

// xpuml.DeviceGetHostVxpuMode()
func DeviceGetHostVxpuMode(Device Device) (HostVxpuMode, Return) {
	var Mode HostVxpuMode
	ret := xpumlDeviceGetHostVxpuMode(Device, &Mode)
	return Mode, ret
}

// xpuml.DeviceSetSriovVfNum()
func DeviceSetSriovVfNum(Device Device, VfNum int32) Return {
	ret := xpumlDeviceSetSriovVfNum(Device, VfNum)
	return ret
}
