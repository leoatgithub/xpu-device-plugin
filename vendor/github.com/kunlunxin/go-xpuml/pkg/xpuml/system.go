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

// xpuml.SystemGetDriverVersion()
func SystemGetDriverVersion() (string, Return) {
	Version := make([]byte, SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	ret := xpumlSystemGetDriverVersion(&Version[0], SYSTEM_DRIVER_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}

// xpuml.SystemGetXPUMLVersion()
func SystemGetXPUMLVersion() (string, Return) {
	Version := make([]byte, SYSTEM_XPUML_VERSION_BUFFER_SIZE)
	ret := xpumlSystemGetXPUMLVersion(&Version[0], SYSTEM_XPUML_VERSION_BUFFER_SIZE)
	return string(Version[:clen(Version)]), ret
}
