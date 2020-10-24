# Protocal

## Protocol Buffer

- Protocol Buffer 利用 varint 原理压缩数据以后，二进制数据非常紧凑，option 也算是压缩体积的一个举措。所以 pb 体积更小，如果选用它作为网络数据传输，势必相同数据，消耗的网络流量更少。但是并没有压缩到极限，float、double 浮点型都没有压缩
- Protocol Buffer 比 JSON 和 XML 少了 {、}、: 这些符号，体积也减少一些。再加上 varint 压缩，gzip 压缩以后体积更小
- Protocol Buffer 是 Tag - Value (Tag - Length - Value)的编码方式的实现，减少了分隔符的使用，数据存储更加紧凑
- Protocol Buffer 另外一个核心价值在于提供了一套工具，一个编译工具，自动化生成 get/set 代码。简化了多语言交互的复杂度，使得编码解码工作有了生产力
- Protocol Buffer 不是自我描述的，离开了数据描述 .proto 文件，就无法理解二进制数据流。这点即是优点，使数据具有一定的“加密性”，也是缺点，数据可读性极差。所以 Protocol Buffer 非常适合内部服务之间 - RPC 调用和传递数据
- Protocol Buffer 具有向后兼容的特性，更新数据结构以后，老版本依旧可以兼容，这也是 Protocol Buffer 诞生之初被寄予解决的问题。因为编译器对不识别的新增字段会跳过不处理
