# XPU Device Plugin for Kubernetes

# 背景

XPU device plugin是昆仑芯提供的K8S昆仑设备(XPU)管理插件,该插件以Daemonset的形式运行在K8S的节点上，并提供如下几个功能:

- 以容器为粒度自动调度节点上昆仑设备资源。
- 监控节点上的昆仑设备健康状态。
- 配置昆仑设备的虚拟化功能。

类似：[k8s-device-plugin](https://github.com/NVIDIA/k8s-device-plugin)。

# 环境依赖

XPU device plugin依赖如下几个软硬件环境:
- 昆仑设备
- 昆仑XRE软件包 >= 4.0.11
    - kunlun.ko >= 4.0.11
    - xpuml.so  >= 4.0.11
- Kubernetes >= 1.10

# 部署方法

## 机器准备

部署XPU device plugin前，需确保node上的昆仑设备及昆仑已安装完毕，确保昆仑设备即'dev/xpu*'可被正常访问。

## 部署Daemonset

运行如下命令部署XPU devcie plugin daemonset:

```
$ git clone xxx/xpu-device-plugin
$ cd xpu-device-plugin
$ kubectl create -f ./deployment/xpu-device-plugin.yml
```
其中xpu-device-plugin.yml文件中默认使用的是昆仑芯官方发布的镜像:
```
 - image: kunlunxpu/xpu-device-plugin:v1.0.0
```

Damonset部署成功后，可通过`$ kubectl describe node`看到每个node上的昆仑设备数量(以baidu.com/xpu标识):

```
...
Capacity:
  baidu.com/xpu:      4
  cpu:                12
  ephemeral-storage:  73364224Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             7746896Ki
  pods:               110
Allocatable:
  baidu.com/xpu:      4
  cpu:                12
  ephemeral-storage:  67612468727
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             7644496Ki
  pods:               110
...
```

## 创建XPU Pod

用户创建pod时可通过`baidu.com/xpu`资源类型来申请昆仑设备，模板如下:

```
    apiVersion: v1
    kind: Pod
    metadata:
      name: xpu-pod
    spec:
      containers:
        - name: xpu-container
          image: ubuntu:16.04
          command:
            - /bin/sleep
            - 10h
          resources:
            limits:
              baidu.com/xpu: 2 # requesting 2 xpus
```

此外，`script/`目录下已存在若干deployment、pod创建模板，可供实验使用。

# 本地编译

若需要修改XPU device plugin, 可运行根目录的下的`build.sh`,或使用如下命令重新编译xpu-device-plugin镜像:

```
$ sudo docker build \
    -t kunlunxpu/xpu-device-plugin:v1.0.0 \
    -f docker/amd64/Dockerfile.ubuntu16.04 \
    .
```

# 虚拟化功能配置说明

XPU device plugin提供了'sriov-num-vfs'配置项用于配置昆仑设备的虚拟化功能，有两种方式可以修改该配置项:

1. 通过yaml文件的args配置项"--sriov-num-vfs=x"修改，其中'x'为每张卡需要切分的VF个数。
2. 通过给node添加label修改，label格式为"baidu.com/sriov-num-vfs=x"。

若采用方式2，需执行以下命令为K8S的默认serviceaccount赋权，保证device-plugin具备获取node label的能力：
```
 kubectl create -f ./deployment/xpu-operator_rbac.authorization.k8s.io_v1_clusterrole.yaml
 kubectl create -f ./deployment/xpu-operator_rbac.authorization.k8s.io_v1_clusterrolebinding.yaml
```

插件会优先使用label中的配置，若label不存在，则采用yaml文件中的配置，默认的'sriov-num-vfs'配置值为0，即虚拟化功能关闭。

开启虚拟化后昆仑设备将被拆分成多个昆仑虚拟设备，每个虚拟设备具备原昆仑物理设备的部分或全部资源，各个虚拟设备间资源相互隔离，可独立工作。关于昆仑设备虚拟化功能的具体说明可参考昆仑XRE用户使用文档。



