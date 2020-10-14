# Golang

## 基础

### Go 协程

### Go Worker Pool

### GMP

### Atomic

- 原子操作：增减、比较并交换、载入、存储、交换
- CPU 需要操作一个内存块的时候，向总线发送一个 LOCK 信号，所有 CPU 收到这个信号后就不对这个内存块进行操作了。 等待操作的 CPU 执行完操作后，发送 UNLOCK 信号，才结束

### CAS

### sync.Mutex 和 sync.RWMutex

### Map, Slice, Channel

### IO
