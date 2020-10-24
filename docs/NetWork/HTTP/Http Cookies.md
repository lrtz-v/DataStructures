# Http Cookie

- 服务器发送到用户浏览器并保存在本地的一小块数据，会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上
- 作用：
  - 会话状态管理（如用户登录状态、购物车、游戏分数或其它需要记录的信息）
  - 个性化设置（如用户自定义设置、主题等）
  - 浏览器行为跟踪（如跟踪分析用户行为等）

## 创建 Cookie

- 服务器使用[`Set-Cookie`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Set-Cookie)响应头部向用户代理（一般是浏览器）发送 Cookie 信息。一个简单的 Cookie 可能像这样：Set-Cookie: <cookie 名>=<cookie 值>
- 会话期 Cookie
  - 浏览器关闭之后它会被自动删除
- 持久性 Cookie
  - 持久性 Cookie 可以指定一个特定的过期时间（`Expires`）或有效期（`Max-Age`）。当 Cookie 的过期时间被设定时，设定的日期和时间只与客户端相关，而不是服务端。
  - Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT;

### Cookie 的`Secure` 和`HttpOnly` 标记

- 标记为 `Secure` 的 Cookie 只应通过被 HTTPS 协议加密过的请求发送给服务端
- 为避免跨域脚本 ([XSS](https://developer.mozilla.org/en-US/docs/Glossary/XSS)) 攻击，通过 JavaScript 的 [`Document.cookie`](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/cookie) API 无法访问带有 `HttpOnly` 标记的 Cookie，它们只应该发送给服务端

- Set-Cookie: id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT; Secure; HttpOnly

### Cookie 的作用域

- `Domain` 和 `Path` 标识定义了 Cookie 的*作用域：*即 Cookie 应该发送给哪些 URL。

### `SameSite` Cookies

- `SameSite` Cookie 允许服务器要求某个 cookie 在跨站请求时不会被发送，从而可以阻止跨站请求伪造攻击（[CSRF](https://developer.mozilla.org/en-US/docs/Glossary/CSRF)）。
- Set-Cookie: key=value; SameSite=Strict
- SameSite
  - **None**： 浏览器会在同站请求、跨站请求下继续发送 cookies，不区分大小写。
  - **Strict**：浏览器将只发送相同站点请求的 cookie(即当前网页 URL 与请求目标 URL 完全一致)。如果请求来自与当前 location 的 URL 不同的 URL，则不包括标记为 Strict 属性的 cookie。
  - **Lax**： 在新版本的浏览器中，SameSite 的默认属性是 SameSite=Lax。换句话说，当 Cookie 没有设置 SameSite 属性时，将会视作 SameSite 属性被设置为 Lax——这意味着 Cookies 将会在当前用户使用时被自动发送。如果想要指定 Cookies 在同站、跨站请求都被发送，那么需要明确指定 SameSite 为 None
