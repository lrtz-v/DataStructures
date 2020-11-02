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
    - 如果数量大于 1，则需要删除多余的 Pod
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

## Kubernetes“声明式 API”

- 概念

  - 所谓“声明式”，指的就是我只需要提交一个定义好的 API 对象来“声明”，我所期望的状态是什么样子
  - “声明式 API”允许有多个 API 写端，以 PATCH 的方式对 API 对象进行修改，而无需关心本地原始 YAML 文件的内容
  - Kubernetes 项目才可以基于对 API 对象的增、删、改、查，在完全无需外界干预的情况下，完成对“实际状态”和“期望状态”的调谐（Reconcile）过程

- 一个 API 对象在 Etcd 里的完整资源路径，是由：Group（API 组）、Version（API 版本）和 Resource（API 资源类型）三个部分组成的

  - 例如

  ```yaml
  apiVersion: batch/v2alpha1
  kind: CronJob
  ```

  - “batch”就是它的组（Group），“v2alpha1” 就是它的版本（Version），“CronJob”就是这个 API 对象的资源类型（Resource）

- Kubernetes 是如何对 Resource、Group 和 Version 进行解析，找到对象的定义呢

  - 首先，Kubernetes 会匹配 API 对象的组
    - 对于 Kubernetes 里的核心 API 对象，比如：Pod、Node 等，是不需要 Group 的（即：它们的 Group 是“”），Kubernetes 会直接在 /api 这个层级进行下一步的匹配过程
    - 而对于 CronJob 等非核心 API 对象来说，Kubernetes 就必须在 /apis 这个层级里查找它对应的 Group
  - 然后，Kubernetes 会进一步匹配到 API 对象的版本号
  - 最后，Kubernetes 会匹配 API 对象的资源类型

- APIServer 创建对象的过程

  - 当我们发起了创建（例如：CronJob）的 POST 请求之后，我们编写的 YAML 的信息就被提交给了 APIServer；而 APIServer 的第一个功能，就是过滤这个请求，并完成一些前置性的工作，比如授权、超时处理、审计等
  - 然后，请求会进入 MUX 和 Routes 流程；APIServer 的 Handler 要做的事情，就是按照上面介绍的匹配过程，找到对应的 CronJob 类型定义
  - 根据这个（例如：CronJob） 类型定义，使用用户提交的 YAML 文件里的字段，创建一个 CronJob 对象
    - APIServer 会进行一个 Convert 工作：把用户提交的 YAML 文件，转换成一个叫作 Super Version 的对象，它正是该 API 资源类型所有版本的字段全集。这样用户提交的不同版本的 YAML 文件，就都可以用这个 Super Version 对象来进行处理了
  - APIServer 会先后进行 Admission() 和 Validation() 操作
    - Admission：例如 Admission Controller 和 Initializer
    - Validation：负责验证这个对象里的各个字段是否合法；这个被验证过的 API 对象，都保存在了 APIServer 里一个叫作 Registry 的数据结构中
  - APIServer 会把验证过的 API 对象转换成用户最初提交的版本，进行序列化操作，并调用 Etcd 的 API 把它保存起来

- 通过 CRD 自定义资源，并自定义控制器
  - //TODO

## RBAC

- 在 Kubernetes 项目中，负责完成授权（Authorization）工作的机制
- 概念
  - Role：角色，它其实是一组规则，定义了一组对 Kubernetes API 对象的操作权限
  - Subject：被作用者，既可以是“人”，也可以是“机器”，也可以是你在 Kubernetes 里定义的“用户”
  - RoleBinding：定义了“被作用者”和“角色”的绑定关系

## Etcd Operator

- //TODO

## 容器持久化存储

- 概念
  - PVC 描述的，是 Pod 想要使用的持久化存储的属性，比如存储的大小、读写权限等
  - PV 描述的，则是一个具体的 Volume 的属性，比如 Volume 的类型、挂载目录、远程存储服务器地址等
  - StorageClass 的作用，一是充当 PV 的模板，并且只有同属于一个 StorageClass 的 PV 和 PVC，才可以绑定在一起；二是指定 PV 的 Provisioner（存储插件）
- //TODO
