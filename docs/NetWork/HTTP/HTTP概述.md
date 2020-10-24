# Http 协议

- client-server 协议
  - 请求通过一个实体被发出，实体也就是用户代理，发送到服务器的请求，会被服务器处理并返回一个消息，也就是*response*。请求与响应之间，还有许许多多的被称为 proxies 的实体，他们的作用与表现各不相同，比如有些是网关，还有些是[caches](https://developer.mozilla.org/en-US/docs/Glossary/Cache)等。

## 客户端：user-agent

- 任何能够为用户发起行为的工具，例如浏览器，浏览器来负责发送 HTTP 请求，并进一步解析 HTTP 返回的消息，以向用户提供明确的响应。

## Web 服务端

- 由 Web Server 来*服务*并提供客户端所请求的资源

## 代理（Proxies）

- 在浏览器和服务器之间，有许多计算机和其他设备转发了 HTTP 消息，大多都出现在传输层、网络层和物理层上，对于 HTTP 应用层而言就是透明的
- 作用
  - 缓存（可以是公开的也可以是私有的，像浏览器的缓存）
  - 过滤（像反病毒扫描，家长控制...）
  - 负载均衡（让多个服务器服务不同的请求）
  - 认证（对不同资源进行权限管理）
  - 日志记录（允许存储历史信息）

## HTTP 和连接

- 一个连接是由传输层来控制的，这从根本上不属于 HTTP 的范围。HTTP 并不需要其底层的传输层协议是面向连接的，只需要它是可靠的，或不丢失消息的
- 在客户端（通常指浏览器）与服务器能够交互（客户端发起请求，服务器返回响应）之前，必须在这两者间建立一个 TCP 链接，打开一个 TCP 连接需要多次往返交换消息（因此耗时）。
- 以 UDP 为基础的传输协议[QUIC](https://en.wikipedia.org/wiki/QUIC)

## HTTP 能控制什么

- 缓存
  文档如何缓存能通过 HTTP 来控制。服务端能告诉代理和客户端哪些文档需要被缓存，缓存多久，而客户端也能够命令中间的缓存代理来忽略存储的文档。
- _开放同源限制_
  为了防止网络窥听和其它隐私泄漏，浏览器强制对 Web 网站做了分割限制。只有来自于**相同来源**的网页才能够获取网站的全部信息。这样的限制有时反而成了负担，HTTP 可以通过修改头部来开放这样的限制，因此 Web 文档可以是由不同域下的信息拼接成的（某些情况下，这样做还有安全因素考虑）。
- _认证_
  一些页面能够被保护起来，仅让特定的用户进行访问。基本的认证功能可以直接通过 HTTP 提供，使用`Authenticate`相似的头部即可，或用 HTTP Cookies 来设置指定的会话。
- _[代理和隧道](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Proxy_servers_and_tunneling)_
  通常情况下，服务器和/或客户端是处于内网的，对外网隐藏真实 IP 地址。因此 HTTP 请求就要通过代理越过这个网络屏障。但并非所有的代理都是 HTTP 代理。例如，SOCKS 协议的代理就运作在更底层，一些像 FTP 这样的协议也能够被它们处理。
- _会话_
  使用 HTTP Cookies 允许你用一个服务端的状态发起请求，这就创建了会话。虽然基本的 HTTP 是无状态协议。这很有用，不仅是因为这能应用到像购物车这样的电商业务上，更是因为这使得任何网站都能轻松为用户定制展示内容了。

## HTTP 流

- 当客户端想要和服务端进行信息交互时，过程表现为下面几步：
  - 打开一个 TCP 连接（或重用一个已经存在的连接）
  - 发送一个 HTTP 报文
  - 读取服务端返回的报文信息
  - 关闭连接或者为后续请求重用连接
- HTTP 流水线已被在有多请求下表现得更稳健的 HTTP/2 的帧所取代

## HTTP 报文

### 请求

- ![A basic HTTP request](./HTTP_Request.png)

- 一个 HTTP 的[method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)，经常是由一个动词像[`GET`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/GET), [`POST`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/POST) 或者一个名词像[`OPTIONS`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/OPTIONS)，[`HEAD`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/HEAD)来定义客户端的动作行为。

- 要获取的资源的路径，通常是上下文中就很明显的元素资源的 URL，它没有[protocol](https://developer.mozilla.org/en-US/docs/Glossary/protocol) （`http://`），[domain](https://developer.mozilla.org/en-US/docs/Glossary/domain)（`developer.mozilla.org`），或是 TCP 的[port](https://developer.mozilla.org/en-US/docs/Glossary/port)（HTTP 一般在 80 端口）。
- HTTP 协议版本号。
- 为服务端表达其他信息的可选头部[headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)。
- 对于一些像 POST 这样的方法，报文的 body 就包含了发送的资源，这与响应报文的 body 类似。

### 响应

- ![img](./HTTP_Response.png)

- HTTP 协议版本号。
- 一个状态码（[status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)），来告知对应请求执行成功或失败，以及失败的原因。
- 一个状态信息，这个信息是非权威的状态码描述信息，可以由服务端自行设定。
- HTTP [headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)，与请求头部类似。
- 可选项，比起请求报文，响应报文中更常见地包含获取的资源 body。
