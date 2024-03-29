# Copyright (c) 2020, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


.PHONY: all build builder test
.DEFAULT_GOAL := all

##### Global variables #####

DOCKER   ?= docker
REGISTRY ?= kunlunxpu
VERSION  ?= v1.0.1

##### Public rules #####

all: ubuntu16.04

push:
	$(DOCKER) push "$(REGISTRY)/xpu-device-plugin:$(VERSION)-ubuntu16.04"
	$(DOCKER) push "$(REGISTRY)/xpu-device-plugin:$(VERSION)-centos7"
	$(DOCKER) push "$(REGISTRY)/xpu-device-plugin:$(VERSION)-ubi8"

push-short:
	$(DOCKER) tag "$(REGISTRY)/xpu-device-plugin:$(VERSION)-ubuntu16.04" "$(REGISTRY)/xpu-device-plugin:$(VERSION)"
	$(DOCKER) push "$(REGISTRY)/xpu-device-plugin:$(VERSION)"

push-latest:
	$(DOCKER) tag "$(REGISTRY)/xpu-device-plugin:$(VERSION)-ubuntu16.04" "$(REGISTRY)/xpu-device-plugin:latest"
	$(DOCKER) push "$(REGISTRY)/xpu-device-plugin:latest"

ubuntu16.04:
	$(DOCKER) build --pull \
		--tag $(REGISTRY)/xpu-device-plugin:$(VERSION)-ubuntu16.04 \
		--file docker/amd64/Dockerfile.ubuntu16.04 .

ubi8:
	$(DOCKER) build --pull \
		--tag $(REGISTRY)/xpu-device-plugin:$(VERSION)-ubi8 \
		--file docker/amd64/Dockerfile.ubi8 .

centos7:
	$(DOCKER) build --pull \
		--tag $(REGISTRY)/xpu-device-plugin:$(VERSION)-centos7 \
		--file docker/amd64/Dockerfile.centos7 .

