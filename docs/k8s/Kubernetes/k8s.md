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

## Deployment Controller 控制器模式

### ReplicaSet

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
  - 设置 rollingUpdate 的 partition 控制 Pod 更新范围

### StatefulSet

- 应用状态

  - 拓扑状态
    - 服务之间存在启动顺序与依赖关系
    - 使用 Pod 模板创建 Pod 的时候，对它们进行编号，并且按照编号顺序逐一完成创建工作
    - 通过 Headless Service 的方式，StatefulSet 为每个 Pod 创建了一个固定并且稳定的 DNS 记录，来作为它的访问入口
  - 存储状态
    - 应用的多个实例分别绑定了不同的存储数据， Pod 的重启不影响数据一致
    - PVC
      - 使用方式
        - 定义一个 PVC，声明想要的 Volume 的属性
        - 在应用的 Pod 中，声明使用这个 PVC
    - PV
      - 根据 PVC 分配 Volume

- 应用管理
  - StatefulSet 的控制器直接管理的是 Pod
    - 每个 Pod 的 hostname、名字等都是不同的、携带了编号的。而 StatefulSet 区分这些实例的方式，就是通过在 Pod 的名字里加上事先约定好的编号
  - Kubernetes 通过 Headless Service，为这些有编号的 Pod，在 DNS 服务器中生成带有同样编号的 DNS 记录
  - StatefulSet 还为每一个 Pod 分配并创建一个同样编号的 PVC

## Service

- Service 是 Kubernetes 项目中用来将一组 Pod 暴露给外界访问的一种机制

- Service 又是如何被访问的
  - 以 Service 的 VIP（Virtual IP，即：虚拟 IP）方式
  - 以 Service 的 DNS 方式
    - Normal Service
      - 通过 DNS 解析获取，拿到 Service 的 VIP
    - Headless Service
      - 通过 DNS 解析获取，直接获取到一个 Pod 的 IP，不需要分配一个 VIP
      - 在 Yaml 文件中，clusterIP 字段的值是 None，这个 Service 被创建后并不会被分配一个 VIP，而是会以 DNS 记录的方式暴露出它所代理的 Pod；而它所代理的 Pod 是通过 Label Selector 机制选择出来的

## DaemonSet

- 运行 DaemonSet Pod
  - 这个 Pod 运行在 Kubernetes 集群里的每一个节点（Node）上
  - 每个节点上只有一个这样的 Pod 实例
  - 当有新的节点加入 Kubernetes 集群后，该 Pod 会自动地在新节点上被创建出来；而当旧节点被删除后，它上面的 Pod 也相应地会被回收掉
- 使用场景
  - 各种网络插件的 Agent 组件，都必须运行在每一个节点上，用来处理这个节点上的容器网络
  - 各种存储插件的 Agent 组件，也必须运行在每一个节点上，用来在这个节点上挂载远程存储目录，操作容器的 Volume 目录
  - 各种监控组件和日志组件，也必须运行在每一个节点上，负责这个节点上的监控信息和日志搜集
- 如何保证每个 Node 上有且只有一个被管理的 Pod 呢
  - Etcd 里获取所有的 Node 列表，然后遍历所有的 Node。检查当前这个 Node 上是不是有一个携带了 DaemonSet 指定标签的 Pod 在运行
    - 如果没有，则需要创建
    - 如果数量大于1，则需要删除多余的Pod
    - 只有一个，则为正常
- DaemonSet 只管理 Pod 对象，然后通过 nodeAffinity 和 Toleration 这两个调度器的小功能，保证了每个节点上有且只有一个 Pod

## Job Controller

- 并行作业的控制方法
  - spec.parallelism，它定义的是一个 Job 在任意时间最多可以启动多少个 Pod 同时运行
  - spec.completions，它定义的是 Job 至少要完成的 Pod 数目，即 Job 的最小完成数
- Job Controller 的工作原理
  - Job Controller 控制的对象，直接就是 Pod
  - 根据实际在 Running 状态 Pod 的数目、已经成功退出的 Pod 的数目，以及 parallelism、completions 参数的值共同计算出在这个周期里，应该创建或者删除的 Pod 数目，然后调用 Kubernetes API 来执行这个操作

- CronJob
  - 通过 schedule 配置 Cron 表达式
  - spec.concurrencyPolicy 配置 Job 执行时间过长，导致下一个创建 Pod 的周期到达的处理策略
