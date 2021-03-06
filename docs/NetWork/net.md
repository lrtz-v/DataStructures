# net

## OSI

### 物理层

- 通过网络设备传输信号，并完成电信号与数字信号的转换

### 数据链路层

- 发送数据时，进行帧编码，给数据加上一些控制信息（如信息频率、位同步、源地址、目标地址），最终将数据以及控制信息以规定的格式发送到物理层
- 接收数据时，接收来自物理层的数据，检查传输错误，并将数据封装成数据帧，去掉控制信息

### 网络层

- 通过 ARP/NDP 支持网络路由
- 网络层把传输层产生的报文段或用户数据报封装成分组和包进行传送
- 对数据包的地址做校验

### 传输层

- 负责向不同机器进程之间的通信提供通用的数据传输服务
- 协议
  - TCP
  - UDP

### 会话层

- 负责管理网络连接会话

### 展示层

- 负责数据协议的转换

### 应用层

- 通过应用进程间的交互来完成特定网络应用，应用层协议定义的是应用进间的通信和交互的规则，如：HTTP 协议、邮件协议 SMTP、域名系统 DNS

## 三次握手

### 握手过程

- 客户端–发送带有 SYN 标志、seq=x 的数据包–一次握手–服务端，客户端进入 SYN-SENT 状态
- 服务端–发送带有 ACK/SYN 标志、ack=x+1、seq=y 的数据包–二次握手–客户端，服务端从 LISTEN 进入 SYN-RCVD 状态
- 客户端–发送带有带有 ACK 标志、ack=y+1，seq=x+1 的数据包–三次握手–服务端，客户端、服务端进入 ESTABLISHED 状态

### 为什么要三次握手

- 双方确认自己与对方的发送与接收是正常的
  - 第一次握手：Client 什么都不能确认；Server 确认了对方发送正常，自己接收正常
  - 第二次握手：Client 确认了：自己发送、接收正常，对方发送、接收正常；Server 确认了：对方发送正常，自己接收正常
  - 第三次握手：Client 确认了：自己发送、接收正常，对方发送、接收正常；Server 确认了：自己发送、接收正常，对方发送、接收正常
- 防止已经失效的连接请求报文突然又传送到了服务器，从而产生错误
  - 客户端发送了第一个请求连接并且没有丢失，只是因为在网络结点中滞留的时间太长了，客户端重新建立连接完成传输之后，滞留的那一次请求连接，到达了服务器，两次握手的机制将会让客户端和服务器再次建立连接

### 第 2 次握手传回了 ACK，为什么还要传回 SYN

- 接收端传回发送端所发送的 ACK 是为了告诉客户端，我接收到的信息确实就是你所发送的信号了，这表明从客户端到服务端的通信是正常的
- 回传 SYN 则是为了建立并确认从服务端到客户端的通信

## 四次挥手

### 挥手过程

- 客户端-发送一个 FIN、seq=u，用来关闭客户端到服务器的数据传送，客户端进入 FIN-WAIT-1 状态
- 服务器-收到这个 FIN，它返回一个 ACK=1，ack=u+1，带上自己的序列号 seq=v。服务端进入 CLOSE-WAIT 状态
- 客户端收到服务器的确认请求后，此时，客户端就进入 FIN-WAIT-2（终止等待 2）状态
- 服务器-关闭与客户端的连接，发送 FIN=1，ack=u+1 给客户端，服务端进入 LAST-ACK 状态
- 客户端收到服务器的连接释放报文后，必须发出确认，ACK=1，ack=w+1，而自己的序列号是 seq=u+1，客户端进入 TIME-WAIT 状态；经过 2\*MSL（最长报文段寿命）的时间后，才进入 CLOSED 状态
- 服务器只要收到了客户端发出的确认，立即进入 CLOSED 状态

### 为什么客户端最后还要等待 2MSL

- 保证客户端发送的最后一个 ACK 报文能够到达服务器，因为这个 ACK 报文可能丢失
- 防止已经失效的连接请求报文段
  - 客户端发送完最后一个确认报文后，在这个 2MSL 时间中，就可以使本连接持续的时间内所产生的所有报文段都从网络中消失。这样新的连接中不会出现旧连接的请求报文

### 为什么建立连接是三次握手，关闭连接确是四次挥手呢

- 建立连接的时候， 服务器在 LISTEN 状态下，收到建立连接请求的 SYN 报文后，把 ACK 和 SYN 放在一个报文里发送给客户端
- 而关闭连接时，服务器收到对方的 FIN 报文时，仅仅表示对方不再发送数据了但是还能接收数据，而自己也未必全部数据都发送给对方了，所以己方可以立即关闭，也可以发送一些数据给对方后，再发送 FIN 报文给对方来表示同意现在关闭连接，因此，己方 ACK 和 FIN 一般都会分开发送，从而导致多了一次

### 如果已经建立了连接，但是客户端突然出现故障了怎么办

- TCP 还设有一个保活计时器，服务器每收到一次客户端的请求后都会重新复位这个计时器，时间通常是设置为 2 小时
- 若两小时还没有收到客户端的任何数据，服务器就会发送一个探测报文段，以后每隔 75 秒发送一次。若一连发送 10 个探测报文仍然没反应，服务器就认为客户端出了故障，接着就关闭连接

## TCP 协议如何保证可靠传输

- 应用数据被分割成 TCP 认为最适合发送的数据块
- TCP 给发送的每一个包进行编号，接收方对数据包进行排序，把有序数据传送给应用层
- 校验和
  - TCP 将保持它首部和数据的检验和，检测数据在传输过程中的任何变化
- TCP 的接收端会丢弃重复的数据
- 流量控制
  - TCP 连接的每一方都有固定大小的缓冲空间，TCP 的接收端只允许发送端发送接收端缓冲区能接纳的数据
  - 当接收方来不及处理发送方的数据，能提示发送方降低发送的速率，防止包丢失
  - TCP 使用的流量控制协议是可变大小的滑动窗口协议
- 拥塞控制
  - 当网络拥塞时，减少数据的发送
- ARQ 协议
  - 也是为了实现可靠传输的，它的基本原理就是每发完一个分组就停止发送，等待对方确认。在收到确认后再发下一个分组

### ARQ 协议（自动重传请求）

- 数据链路层和传输层的错误纠正协议之一，通过使用确认和超时这两个机制，在不可靠服务的基础上实现可靠的信息传输。如果发送方在发送后一段时间之内没有收到确认帧，它通常会重新发送
- 停止等待 ARQ 协议
  - 每发完一个分组就停止发送，等待对方确认（回复 ACK）。如果过了一段时间（超时时间后），还是没有收到 ACK 确认，说明没有发送成功，需要重新发送，直到收到确认后再发下一个分组
  - 若接收方收到重复分组，就丢弃该分组，但同时还要发送确认
  - 缺点： 信道利用率低，等待时间长
    - 超时重传
      - 每发送完一个分组需要设置一个超时计时器，只要超过一段时间仍然没有收到确认，就重传前面发送过的分组
    - 确认消息丢失
      - 确认消息在传输过程丢失，数据包重传
    - 确认消息迟到
      - 确认消息在传输过程中迟到，数据包重传
- 连续 ARQ 协议
  - 发送方维持一个发送窗口，凡位于发送窗口内的分组可以连续发送出去，而不需要等待对方确认
  - 接收方一般采用累计确认，对按序到达的最后一个分组发送确认，表明到这个分组为止的所有分组都已经正确收到了
  - 缺点
    - 不能向发送方反映出接收方已经正确收到的所有分组的信息
    - 比如：发送方发送了 5 条 消息，中间第三条丢失（3 号），这时接收方只能对前两个发送确认。发送方无法知道后三个分组的下落，而只好把后三个全部重传一次。这也叫 Go-Back-N（回退 N），表示需要退回来重传已经发送过的 N 个消息

### 滑动窗口流量控制

- 流量控制是为了控制发送方发送速率，保证接收方来得及接收
- 接收方发送的确认报文中的窗口字段可以用来控制发送方窗口大小，从而影响发送方的发送速率。将窗口字段设置为 0，则发送方不能发送数据

### 拥塞控制

- TCP 发送方要维持一个 拥塞窗口(cwnd) 的状态变量，拥塞控制窗口的大小取决于网络的拥塞程度，并且动态变化。发送方让自己的发送窗口取为拥塞窗口和接收方的接受窗口中较小的一个
- TCP 的拥塞控制采用了四种算法，即 慢开始、 拥塞避免 、快重传 和 快恢复
  - 慢开始：由小到大逐渐增大发送窗口
  - 拥塞避免：每经过一个往返时间 RTT 就把发送放的拥塞窗口加 1
  - 快重传与快恢复
    - 没有 FRR，如果数据包丢失了，TCP 将会使用定时器来要求传输暂停。在暂停的这段时间内，没有新的或复制的数据包被发送
    - 有了 FRR，如果接收机接收到一个不按顺序的数据段，它会立即给发送机发送一个重复确认。如果发送机接收到三个重复确认，它会假定确认件指出的数据段丢失了，并立即重传这些丢失的数据段，不会因为重传时要求的暂停被耽误

## HTTP 版本

### HTTP 0.9

- 只支持 GET 请求方式
- 没有请求头概念：所以不能在请求中指定版本号，服务端也只具有返回 HTML 字符串的能力
- 服务端相响应之后，立即关闭 TCP 连接

### HTTP 1.0

- 请求方式新增了 POST，DELETE，PUT，HEADER 等方式
- 增添了请求头和响应头的概念，在通信中指定了 HTTP 协议版本号，以及其他的一些元信息 (比如: 状态码、权限、缓存、内容编码)
- 扩充了传输内容格式，图片、音视频资源、二进制等都可以进行传输
- 特性
  - 无状态：服务器不跟踪不记录请求过的状态
  - 无连接：浏览器每次请求都需要建立 tcp 连接

### HTTP 1.1

- 长连接
  - 新增 Connection 字段，可以设置 keep-alive 值保持连接不断开
  - 但域名分片等情况下仍然需要建立多个 connection
- 管道化
  - 基于长连接的基础，管道化可以不等第一个请求响应继续发送后面的请求，但响应的顺序还是按照请求的顺序返回
- 缓存处理
  - 当浏览器请求资源时，先看是否有缓存的资源，如果有缓存，直接取，不会再发请求，如果没有缓存，则发送请求。 通过设置字段 cache-control 来控制缓存。
- 断点传输
  - 资源过大，将其分割为多个部分，分别上传/下载

### HTTP 2

- 二进制分帧
  - HTTP 1.x 的解析是基于文本，HTTP 2 之后将所有传输的信息分割为更小的消息和帧，并对它们采用二进制格式的编码，提高传输效率
- 多路复用
  - 同时发送请求和响应，基于二进制分帧，在同一域名下所有访问都是从同一个 tcp 连接中走
  - 单个连接可以承载任意数量的双向数据流
  - 数据流以消息的形式发送，而消息又由一个或多个帧组成，多个帧之间可以乱序发送，因为根据帧首部的流标识可以重新组装
- Header 压缩
  - 在 HTTP/1 中，HTTP 请求和响应都是由「状态行、请求 / 响应头部、消息主体」三部分组成。一般而言，消息主体都会经过 gzip 压缩，或者本身传输的就是压缩过后的二进制文件（例如图片、音频），但状态行和头部却没有经过任何压缩，直接以纯文本传输
  - HTTP/2 在客户端和服务器端使用“首部表”来跟踪和存储之前发送的键－值对，对于相同的数据，不再通过每次请求和响应发送
    - 维护一份相同的静态字典（Static Table），包含常见的头部名称，以及特别常见的头部名称与值的组合
    - 维护一份相同的动态字典（Dynamic Table），可以动态地添加内容， 动态字典积累得越全，头部压缩效果也就越好
- 服务器推送：服务器可以额外的向客户端推送资源，而无需客户端明确的请求

### HTTP 3

- 使用基于 UDP 协议的 QUIC 协议
  - 0-RTT
    - 缓存当前会话的上下文，在下次恢复会话的时候，只需要将之前的缓存传递给服务端验证通过就可以进行传输了
  - 多路复用
    - 同一条 QUIC 连接上可以创建多个 stream，来发送多个 HTTP 请求，但是，QUIC 是基于 UDP 的，一个连接上的多个 stream 之间没有依赖
    - 丢了一个 UDP 包，不会影响后面跟着的 Stream，不存在 TCP 队头阻塞
  - 加密认证的报文
    - QUIC 的 packet 中，除了个别报文比如 PUBLIC_RESET 和 CHLO，所有报文头部都是经过认证的，报文 Body 都是经过加密的。
  - 向前纠错机制
    - 每个数据包除了它本身的内容之外，还包括了部分其他数据包的数据，因此少量的丢包可以通过其他包的冗余数据直接组装而无需重传
    - 牺牲了每个数据包可以发送数据的上限，但是减少了因为丢包导致的数据重传，因为数据重传将会消耗更多的时间
    - 只能使用在丢失一个包的情况下，如果出现丢失多个包就不能使用纠错机制了，只能使用重传的方式了
