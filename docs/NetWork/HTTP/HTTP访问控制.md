# CORS

- 使用额外的 HTTP 头来告诉浏览器让运行在一个 origin (domain) 上的 Web 应用被准许访问来自不同源服务器上的资源
- 当一个资源从与该资源本身所在的服务器**不同的域、协议或端口**请求一个资源时，资源会发起一个**跨域 HTTP 请求**

## 功能概述

- 允许服务器声明哪些源站通过浏览器有权限访问哪些资源
- 对那些可能对服务器数据产生副作用的 HTTP 请求方法（特别是 [`GET`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/GET) 以外的 HTTP 请求，或者搭配某些 MIME 类型的 [`POST`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/POST) 请求），浏览器必须首先使用 [`OPTIONS`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/OPTIONS) 方法发起一个预检请求（preflight request），从而获知服务端是否允许该跨域请求

## 若干访问控制场景

### 简单请求（不会触发 [CORS 预检请求](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS#Preflighted_requests)）

- 使用下列方法之一：

  - [`GET`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/GET)
  - [`HEAD`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/HEAD)
  - [`POST`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/POST)

- Fetch 规范定义了

  对 CORS 安全的首部字段集合

  ，不得人为设置该集合之外的其他首部字段。该集合为：

  - [`Accept`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept)
  - [`Accept-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept-Language)
  - [`Content-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Language)
  - [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) （需要注意额外的限制）
  - `DPR`
  - `Downlink`
  - `Save-Data`
  - `Viewport-Width`
  - `Width`

- `Content-Type`的值仅限于下列三者之一：

  - `text/plain`
  - `multipart/form-data`
  - `application/x-www-form-urlencoded`

- 请求中的任意`XMLHttpRequestUpload` 对象均没有注册任何事件监听器；`XMLHttpRequestUpload` 对象可以使用 [`XMLHttpRequest.upload`](https://developer.mozilla.org/zh-CN/docs/Web/API/XMLHttpRequest/upload) 属性访问。

- 请求中没有使用 [`ReadableStream`](https://developer.mozilla.org/zh-CN/docs/Web/API/ReadableStream) 对象。

- ![img](./simple_req.png)

### 预检请求

- 使用了下面任一 HTTP 方法：

  - [`PUT`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/PUT)
  - [`DELETE`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/DELETE)
  - [`CONNECT`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/CONNECT)
  - [`OPTIONS`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/OPTIONS)
  - [`TRACE`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/TRACE)
  - [`PATCH`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods/PATCH)

- 人为设置了

  对 CORS 安全的首部字段集合

  之外的其他首部字段。该集合为：

  - [`Accept`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept)
  - [`Accept-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Accept-Language)
  - [`Content-Language`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Language)
  - [`Content-Type`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Content-Type) (需要注意额外的限制)
  - `DPR`
  - `Downlink`
  - `Save-Data`
  - `Viewport-Width`
  - `Width`

- `Content-Type`

  的值不属于下列之一:

  - `application/x-www-form-urlencoded`
  - `multipart/form-data`
  - `text/plain`

- 请求中的`XMLHttpRequestUpload` 对象注册了任意多个事件监听器。

- 请求中使用了[`ReadableStream`](https://developer.mozilla.org/zh-CN/docs/Web/API/ReadableStream)对象。

- ![img](./preflight_correct.png)

### 附带身份凭证的请求

- [Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) 与 CORS 的一个有趣的特性是，可以基于 [HTTP cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) 和 HTTP 认证信息发送身份凭证。一般而言，对于跨域 [`XMLHttpRequest`](https://developer.mozilla.org/zh-CN/docs/Web/API/XMLHttpRequest) 或 [Fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) 请求，浏览器**不会**发送身份凭证信息。如果要发送凭证信息，需要设置 `XMLHttpRequest `的某个特殊标志位。
- ![img](./cred-req.png)

## HTTP 响应首部字段

### Access-Control-Allow-Origin

- Access-Control-Allow-Origin: http://mozilla.com
- Access-Control-Allow-Origin: \*

### Access-Control-Expose-Headers

- 让服务器把允许浏览器访问的头放入白名单
- Access-Control-Expose-Headers: X-My-Custom-Header, X-Another-Custom-Header

### Access-Control-Max-Age

- 指定了 preflight 请求的结果能够被缓存多久
- Access-Control-Max-Age: <delta-seconds>

### Access-Control-Allow-Credentials

- 当浏览器的`credentials`设置为 true 时是否允许浏览器读取 response 的内容
- Access-Control-Allow-Credentials: true

### Access-Control-Allow-Methods

- 用于预检请求的响应。其指明了实际请求所允许使用的 HTTP 方法。
- Access-Control-Allow-Methods: <method>[, <method>]\*

### Access-Control-Allow-Headers

- 用于预检请求的响应。其指明了实际请求中允许携带的首部字段。
- Access-Control-Allow-Headers: <field-name>[, <field-name>]\*

## HTTP 请求首部字段

### Origin

- 表明预检请求或实际请求的源站。

### Access-Control-Request-Method

- 用于预检请求。其作用是，将实际请求所使用的 HTTP 方法告诉服务器。
- Access-Control-Request-Method: <method>

### Access-Control-Request-Headers

- 用于预检请求。其作用是，将实际请求所携带的首部字段告诉服务器。
- Access-Control-Request-Headers: <field-name>[, <field-name>]\*
