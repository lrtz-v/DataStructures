# 编译

##编译Go程序的时候需要注意以下几点
* 传递参数-ldflags “-s”，忽略debug的打印信息
* 传递-gcflags “-N -l” 参数，这样可以忽略Go内部做的一些优化，聚合变量和函数等优化，这样对于GDB调试来说非常困难，所以在编译的时候加入这两个参数避免这些优化。
* 编译示例代码go build -gcflags "-N -l" -ldflags "-s" main.go，目录下生成main可执行文件
