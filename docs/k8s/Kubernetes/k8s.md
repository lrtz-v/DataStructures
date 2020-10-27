# Kubernetes

## Kubernetes 解决的问题

- 容器编排
- 提供了一套基于容器构建分布式系统的基础依赖

## Kubernetes 架构

- Master 负责控制节点
  - 负责 API 服务的 kube-apiserver
  - 负责调度的 kube-scheduler
  - 负责容器编排的 kube-controller-manager
- Node 负责计算节点
  - kubelet 的组件
    - 主要负责同容器运行时打交道。所依赖 CRI
    - 通过 gRPC 协议同一个叫作 Device Plugin 的插件进行交互，管理宿主机物理设备
    - 调用网络插件（CNI）和存储插件（CSI）为容器配置网络和持久化存储
