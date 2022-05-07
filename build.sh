#!/bin/bash
set -ex

sudo docker build \
  -t kunlunxpu/xpu-device-plugin:v1.0.0 \
  -f docker/amd64/Dockerfile.ubuntu16.04 \
  .
