package core

type Host struct {
	Name      string // 缓存服务器的 Ip 地址 + 端口, 如：127.0.0.1:8000
	LoadBound int64  // 缓存服务器负载
}
