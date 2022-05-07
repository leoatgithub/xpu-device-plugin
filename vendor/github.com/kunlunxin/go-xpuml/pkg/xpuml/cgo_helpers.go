// Copyright (c) 2022, KUNLUNXIN CORPORATION. All rights reserved.
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

// WARNING: THIS FILE WAS AUTOMATICALLY GENERATED.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package xpuml

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#cgo CFLAGS: -DXPUML_NO_UNVERSIONED_FUNC_DEFS=1
#include "xpuml.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *PciInfo) Ref() *C.xpumlPciInfo_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlPciInfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *PciInfo) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewPciInfoRef converts the C object reference into a raw struct reference without wrapping.
func NewPciInfoRef(ref unsafe.Pointer) *PciInfo {
	return (*PciInfo)(ref)
}

// NewPciInfo allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewPciInfo() *PciInfo {
	return (*PciInfo)(allocPciInfoMemory(1))
}

// allocPciInfoMemory allocates memory for type C.xpumlPciInfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPciInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPciInfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfPciInfoValue = unsafe.Sizeof([1]C.xpumlPciInfo_t{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *PciInfo) PassRef() *C.xpumlPciInfo_t {
	if x == nil {
		x = (*PciInfo)(allocPciInfoMemory(1))
	}
	return (*C.xpumlPciInfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Utilization) Ref() *C.xpumlUtilization_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlUtilization_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Utilization) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewUtilizationRef converts the C object reference into a raw struct reference without wrapping.
func NewUtilizationRef(ref unsafe.Pointer) *Utilization {
	return (*Utilization)(ref)
}

// NewUtilization allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewUtilization() *Utilization {
	return (*Utilization)(allocUtilizationMemory(1))
}

// allocUtilizationMemory allocates memory for type C.xpumlUtilization_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUtilizationMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUtilizationValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfUtilizationValue = unsafe.Sizeof([1]C.xpumlUtilization_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Utilization) PassRef() *C.xpumlUtilization_t {
	if x == nil {
		x = (*Utilization)(allocUtilizationMemory(1))
	}
	return (*C.xpumlUtilization_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Memory) Ref() *C.xpumlMemory_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlMemory_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Memory) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewMemoryRef converts the C object reference into a raw struct reference without wrapping.
func NewMemoryRef(ref unsafe.Pointer) *Memory {
	return (*Memory)(ref)
}

// NewMemory allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewMemory() *Memory {
	return (*Memory)(allocMemoryMemory(1))
}

// allocMemoryMemory allocates memory for type C.xpumlMemory_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocMemoryMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfMemoryValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfMemoryValue = unsafe.Sizeof([1]C.xpumlMemory_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Memory) PassRef() *C.xpumlMemory_t {
	if x == nil {
		x = (*Memory)(allocMemoryMemory(1))
	}
	return (*C.xpumlMemory_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *BAR4Memory) Ref() *C.xpumlBAR4Memory_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlBAR4Memory_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *BAR4Memory) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewBAR4MemoryRef converts the C object reference into a raw struct reference without wrapping.
func NewBAR4MemoryRef(ref unsafe.Pointer) *BAR4Memory {
	return (*BAR4Memory)(ref)
}

// NewBAR4Memory allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewBAR4Memory() *BAR4Memory {
	return (*BAR4Memory)(allocBAR4MemoryMemory(1))
}

// allocBAR4MemoryMemory allocates memory for type C.xpumlBAR4Memory_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBAR4MemoryMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBAR4MemoryValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfBAR4MemoryValue = unsafe.Sizeof([1]C.xpumlBAR4Memory_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *BAR4Memory) PassRef() *C.xpumlBAR4Memory_t {
	if x == nil {
		x = (*BAR4Memory)(allocBAR4MemoryMemory(1))
	}
	return (*C.xpumlBAR4Memory_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *ProcessInfo) Ref() *C.xpumlProcessInfo_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlProcessInfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ProcessInfo) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewProcessInfoRef converts the C object reference into a raw struct reference without wrapping.
func NewProcessInfoRef(ref unsafe.Pointer) *ProcessInfo {
	return (*ProcessInfo)(ref)
}

// NewProcessInfo allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewProcessInfo() *ProcessInfo {
	return (*ProcessInfo)(allocProcessInfoMemory(1))
}

// allocProcessInfoMemory allocates memory for type C.xpumlProcessInfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocProcessInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfProcessInfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfProcessInfoValue = unsafe.Sizeof([1]C.xpumlProcessInfo_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ProcessInfo) PassRef() *C.xpumlProcessInfo_t {
	if x == nil {
		x = (*ProcessInfo)(allocProcessInfoMemory(1))
	}
	return (*C.xpumlProcessInfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *DeviceAttributes) Ref() *C.xpumlDeviceAttributes_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlDeviceAttributes_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *DeviceAttributes) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDeviceAttributesRef converts the C object reference into a raw struct reference without wrapping.
func NewDeviceAttributesRef(ref unsafe.Pointer) *DeviceAttributes {
	return (*DeviceAttributes)(ref)
}

// NewDeviceAttributes allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewDeviceAttributes() *DeviceAttributes {
	return (*DeviceAttributes)(allocDeviceAttributesMemory(1))
}

// allocDeviceAttributesMemory allocates memory for type C.xpumlDeviceAttributes_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDeviceAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDeviceAttributesValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfDeviceAttributesValue = unsafe.Sizeof([1]C.xpumlDeviceAttributes_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *DeviceAttributes) PassRef() *C.xpumlDeviceAttributes_t {
	if x == nil {
		x = (*DeviceAttributes)(allocDeviceAttributesMemory(1))
	}
	return (*C.xpumlDeviceAttributes_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Sample) Ref() *C.xpumlSample_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlSample_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Sample) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewSampleRef converts the C object reference into a raw struct reference without wrapping.
func NewSampleRef(ref unsafe.Pointer) *Sample {
	return (*Sample)(ref)
}

// NewSample allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewSample() *Sample {
	return (*Sample)(allocSampleMemory(1))
}

// allocSampleMemory allocates memory for type C.xpumlSample_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSampleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSampleValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfSampleValue = unsafe.Sizeof([1]C.xpumlSample_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Sample) PassRef() *C.xpumlSample_t {
	if x == nil {
		x = (*Sample)(allocSampleMemory(1))
	}
	return (*C.xpumlSample_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *ViolationTime) Ref() *C.xpumlViolationTime_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlViolationTime_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ViolationTime) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewViolationTimeRef converts the C object reference into a raw struct reference without wrapping.
func NewViolationTimeRef(ref unsafe.Pointer) *ViolationTime {
	return (*ViolationTime)(ref)
}

// NewViolationTime allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewViolationTime() *ViolationTime {
	return (*ViolationTime)(allocViolationTimeMemory(1))
}

// allocViolationTimeMemory allocates memory for type C.xpumlViolationTime_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocViolationTimeMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfViolationTimeValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfViolationTimeValue = unsafe.Sizeof([1]C.xpumlViolationTime_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ViolationTime) PassRef() *C.xpumlViolationTime_t {
	if x == nil {
		x = (*ViolationTime)(allocViolationTimeMemory(1))
	}
	return (*C.xpumlViolationTime_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *ProcessUtilizationSample) Ref() *C.xpumlProcessUtilizationSample_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlProcessUtilizationSample_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ProcessUtilizationSample) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewProcessUtilizationSampleRef converts the C object reference into a raw struct reference without wrapping.
func NewProcessUtilizationSampleRef(ref unsafe.Pointer) *ProcessUtilizationSample {
	return (*ProcessUtilizationSample)(ref)
}

// NewProcessUtilizationSample allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewProcessUtilizationSample() *ProcessUtilizationSample {
	return (*ProcessUtilizationSample)(allocProcessUtilizationSampleMemory(1))
}

// allocProcessUtilizationSampleMemory allocates memory for type C.xpumlProcessUtilizationSample_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocProcessUtilizationSampleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfProcessUtilizationSampleValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfProcessUtilizationSampleValue = unsafe.Sizeof([1]C.xpumlProcessUtilizationSample_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ProcessUtilizationSample) PassRef() *C.xpumlProcessUtilizationSample_t {
	if x == nil {
		x = (*ProcessUtilizationSample)(allocProcessUtilizationSampleMemory(1))
	}
	return (*C.xpumlProcessUtilizationSample_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FieldValue) Ref() *C.xpumlFieldValue_t {
	if x == nil {
		return nil
	}
	return (*C.xpumlFieldValue_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FieldValue) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFieldValueRef converts the C object reference into a raw struct reference without wrapping.
func NewFieldValueRef(ref unsafe.Pointer) *FieldValue {
	return (*FieldValue)(ref)
}

// NewFieldValue allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFieldValue() *FieldValue {
	return (*FieldValue)(allocFieldValueMemory(1))
}

// allocFieldValueMemory allocates memory for type C.xpumlFieldValue_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFieldValueMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFieldValueValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFieldValueValue = unsafe.Sizeof([1]C.xpumlFieldValue_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FieldValue) PassRef() *C.xpumlFieldValue_t {
	if x == nil {
		x = (*FieldValue)(allocFieldValueMemory(1))
	}
	return (*C.xpumlFieldValue_t)(unsafe.Pointer(x))
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}
