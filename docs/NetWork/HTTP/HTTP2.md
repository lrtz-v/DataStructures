# HTTP/2

- HTTP/2 在 HTTP/1.1 有几处基本的不同:
  - HTTP/2 是二进制协议而不是文本协议。不再可读，也不可无障碍的手动创建，改善的优化技术现在可被实施。
  - 这是一个复用协议。并行的请求能在同一个链接中处理，移除了 HTTP/1.x 中顺序和阻塞的约束。
  - 压缩了 headers。因为 headers 在一系列请求中常常是相似的，其移除了重复和传输重复数据的成本。
  - 其允许服务器在客户端缓存中填充数据，通过一个叫服务器推送的机制来提前请求。

## 帧

- HTTP/1.x 报文有一些性能上的缺点：

  - Header 不像 body，它不会被压缩。
  - 两个报文之间的 header 通常非常相似，但它们仍然在连接中重复传输。
  - 无法复用。当在同一个服务器打开几个连接时：TCP 热连接比冷连接更加有效。

- HTTP/2 引入了一个额外的步骤

  - 它将 HTTP/1.x 消息分成帧并嵌入到流 (stream) 中。数据帧和报头帧分离，这将允许报头压缩。将多个流组合，这是一个被称为 _多路复用 (multiplexing)_ 的过程，它允许更有效的底层 TCP 连接。
  - ![HTTP/2 modify the HTTP message to divide them in frames (/Users/lvtao/Documents/Note/HTTP/Binary_framing2.png), allowing for more optimization.](https://mdn.mozillademos.org/files/13819/Binary_framing2.png)

- 在 HTTP/2 中，这是一个在 HTTP/1.1 和底层传输协议之间附加的步骤。Web 开发人员不需要在其使用的 API 中做任何更改来利用 HTTP 帧
