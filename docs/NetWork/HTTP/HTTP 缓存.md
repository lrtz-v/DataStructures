![What a cache provide, advantages/disadvantages of shared/private caches.](./HTTPCachtType.png)

## 缓存的种类

- (私有)浏览器缓存
  - 只能用于单独用户
- (共享)代理缓存
  - 可以被多个用户使用

## 缓存控制（Cache-control）

- 禁止进行缓存

  - 缓存中不得存储任何关于客户端请求和服务端响应的内容。每次由客户端发起的请求都会下载完整的响应内容。
  - Cache-Control: no-store

- 强制确认缓存

  - 每次有请求发出时，缓存会将此请求发到服务器（译者注：该请求应该会带有与本地缓存相关的验证字段），服务器端会验证请求中所描述的缓存是否过期，若未过期（注：实际就是返回 304），则缓存才使用本地缓存副本
  - Cache-Control: no-cache

- 私有缓存和公共缓存

  - "public" 指令表示该响应可以被任何中间人（译者注：比如中间代理、CDN 等）缓存
  - "private" 则表示该响应是专用于某单个用户的，中间人不能缓存此响应，该响应只能应用于浏览器私有缓存中。
  - Cache-Control: public
  - Cache-Control: private

- 缓存过期机制

  - 表示资源能够被缓存的最大时间。max-age 是距离请求发起的时间的秒数。针对应用中那些不会改变的文件，通常可以手动设置一定的时长以保证缓存有效，例如图片、css、js 等静态资源。
  - Cache-Control: max-age=31536000

- 缓存验证确认

  - 当使用了 "`must-revalidate`" 指令，那就意味着缓存在考虑使用一个陈旧的资源时，必须先验证它的状态，已过期的缓存将不被使用
  - Cache-Control: must-revalidate

## 新鲜度

- 由于缓存只有有限的空间用于存储资源副本，所以缓存会定期地将一些副本删除，这个过程叫做缓存驱逐。另一方面，当服务器上面的资源进行了更新，那么缓存中的对应资源也应该被更新
- 一个陈旧的资源（缓存副本）是不会直接被清除或忽略的，当客户端发起一个请求时，缓存检索到已有一个对应的陈旧资源（缓存副本），则缓存会先将此请求附加一个`If-None-Match`头，然后发给目标服务器，以此来检查该资源副本是否是依然还是算新鲜的，若服务器返回了 [`304`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304) (Not Modified)（该响应不会有带有实体信息），则表示此资源副本是新鲜的，
- ![Show how a proxy cache acts when a doc is not cache, in the cache and fresh, in the cache and stale.](./HTTPStaleness-20191218211606123.png)

- expirationTime = responseTime + freshnessLifetime - currentAge

### 加速资源

- 为了优化缓存，过期时间设置得尽量长是一种很好的策略。对于定期或者频繁更新的资源，这么做是比较稳妥的，但是对于那些长期不更新的资源会有点问题
- `revving` 的技术: 不频繁更新的文件会使用特定的命名方式：在 URL 后面（通常是文件名后面）会加上版本号。加上版本号后的资源就被视作一个完全新的独立的资源，同时拥有一年甚至更长的缓存过期时长

- ![img](./HTTPRevved.png)

## 缓存验证

- 当缓存的文档过期后，需要进行缓存验证或者重新获取资源。只有在服务器返回强校验器或者弱校验器时才会进行验证。
- ETag： 一种强校验器，ETag 响应头是一个对用户代理不透明的值。对于像浏览器这样的 HTTP UA，不知道 ETag 代表什么，不能预测它的值是多少。如果资源请求的响应头里含有 ETag, 客户端可以在后续的请求的头中带上 [`If-None-Match`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/If-None-Match) 头来验证缓存。

- Last-Modified 响应头可以作为一种弱校验器。说它弱是因为它只能精确到一秒。如果响应头里含有这个信息，客户端可以在后续的请求中带上 [`If-Modified-Since`](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/If-Modified-Since) 来验证缓存。

## 带 Vary 头的响应

- Vary HTTP 响应头决定了对于后续的请求头，如何判断是请求一个新的资源还是使用缓存的文件。当缓存服务器收到一个请求，只有当前的请求和原始（缓存）的请求头跟缓存的响应头里的 Vary 都匹配，才能使用缓存的响应。

- ![The Vary header leads cache to use more HTTP headers as key for the cache.](./HTTPVary.png)

- 使用 vary 头有利于内容服务的动态多样性。例如，使用 Vary: User-Agent 头，缓存服务器需要通过 UA 判断是否使用缓存的页面。如果需要区分移动端和桌面端的展示内容，利用这种方式就能避免在不同的终端展示错误的布局。

- eg：Vary: User-Agent。因为移动版和桌面的客户端的请求头中的 User-Agent 不同， 缓存服务器不会错误地把移动端的内容输出到桌面端到用户
