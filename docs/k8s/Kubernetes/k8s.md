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

## POD

- 一组共享了某些资源的容器，共享的是同一个 Network Namespace，并且可以声明共享同一个 Volume

### Infra 容器

- 在一个 Pod 中，Infra 容器永远都是第一个被创建的容器，而其他用户定义的容器，则通过 Join Network Namespace 的方式，与 Infra 容器关联在一起

- 对于同一个 POD 中的容器
  - 它们可以直接使用 localhost 进行通信
  - 它们看到的网络设备跟 Infra 容器看到的完全一样
  - 一个 Pod 只有一个 IP 地址，也就是这个 Pod 的 Network Namespace 对应的 IP 地址
  - 当然，其他的所有网络资源，都是一个 Pod 一份，并且被该 Pod 中的所有容器共享；Pod 的生命周期只跟 Infra 容器一致，而与容器 A 和 B 无关。

### Init Container 容器

- 在一个 POD 中，所有 Init Container 定义的容器，都会比 spec.containers 定义的用户容器先启动（sidecar 模式）

### 特俗字段

- NodeSelector 供用户将 Pod 与 Node 进行绑定的字段

  - 这样的一个配置，意味着这个 Pod 永远只能运行在携带了特定标签（Label）的节点上；否则，它将调度失败

- NodeName

  - 一旦 Pod 的这个字段被赋值，Kubernetes 项目就会被认为这个 Pod 已经经过了调度，调度的结果就是赋值的节点名字

- HostAliases

  - 定义了 Pod 的 hosts 文件（比如 /etc/hosts）里的内容

- ImagePullPolicy 镜像拉取的策略

- Lifecycle
  - 是在容器状态发生变化时触发一系列“钩子”

## Volume

- Secret
  - 把 Pod 想要访问的加密数据，存放到 Etcd 中。然后，你就可以通过在 Pod 的容器里挂载 Volume 的方式，访问到这些 Secret 里保存的信息了。
- ConfigMap
  - 保存的是不需要加密的、应用所需的配置信息
- Downward API
  - 让 Pod 里的容器能够直接获取到这个 Pod API 对象本身的信息
- ServiceAccountToken
  - 任何运行在 Kubernetes 集群上的应用，都必须使用这个 ServiceAccountToken 里保存的授权信息，也就是 Token，才可以合法地访问 API Server
  - 每一个 Pod，都已经自动声明一个类型是 Secret、名为 default-token-xxxx 的 Volume

## Probe 健康检查“探针”

- Pod 恢复机制（restartPolicy）
  - Always: 在任何情况下，只要容器不在运行状态，就自动重启容器
  - OnFailure: 只在容器异常时才自动重启容器
  - Never: 从来不重启容器

## 控制器模式

- Deployment -> ReplicaSet -> Pods
  - 通过 ReplicaSet 的个数来描述应用的版本
  - 再通过 ReplicaSet 的属性（比如 replicas 的值），来保证 Pod 的副本数量
- ReplicaSet 负责通过“控制器模式”，保证系统中 Pod 的个数永远等于指定的个数
- 水平扩展 / 收缩
  - Deployment Controller 只需要修改它所控制的 ReplicaSet 的 Pod 副本个数就可以了
  - kubectl scale deployment nginx-deployment --replicas=4
- 滚动更新
  - 将一个集群中正在运行的多个 Pod 版本，交替地逐一升级的过程
  - 可以通过 RollingUpdateStrategy 配置一次“滚动”中创建/删除的 Pod 量
  - 通过 kubectl rollout undo 回滚操作
  - 修改过程
    - kubectl rollout pause 暂停 Deployment
    - kubectl edit / kubectl set image 修改 Deployment
    - kubectl rollout resume 恢复 Deployment
