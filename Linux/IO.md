# Linux IO模型

## 概念

### 用户空间与内核空间

- 操作系统将虚拟空间划分为两部分，一部分为内核空间，一部分为用户空间
- 操作系统的核心是内核，独立于普通的应用程序，可以访问受保护的内存空间，也有访问底层硬件设备的所有权限

### 进程切换

- 保存处理机上下文，包括程序计数器和其他寄存器。
- 更新PCB信息。
- 把进程的PCB移入相应的队列，如就绪、在某事件阻塞等队列。
- 选择另一个进程执行，并更新其PCB。
- 更新内存管理的数据结构。
- 恢复处理机上下文。

### 进程的阻塞

- 当进程进入阻塞状态，是不占用CPU资源的

### 文件描述符

- 是一个用于表述指向文件的引用的抽象化概念
- 它是一个索引值，指向内核为每一个进程所维护的该进程打开文件的记录表

### 缓存I/O

- 缓存 I/O 又被称作标准 I/O，大多数文件系统的默认 I/O 操作都是缓存 I/O
- 在 Linux 的缓存 I/O 机制中，操作系统会将 I/O 的数据缓存在文件系统的页缓存中
- 【*】数据会先被拷贝到操作系统内核的缓冲区中，然后才会从操作系统内核的缓冲区拷贝到应用程序的地址空间
- 缺点：数据在传输过程中需要在应用程序地址空间和内核进行多次数据拷贝操作，消耗较多CPU资源

## I/O模型

### 同步阻塞 I/O（blocking I/O）

- 在等待数据到数据复制的两个阶段，整个进程都被阻塞
- blocking IO的特点就是在IO执行的两个阶段都被block了

### 同步非阻塞 I/O（nonblocking I/O）

- 内核马上返回给进程，如果数据还没准备好，此时会返回一个error。进程在返回之后，可以干点别的事情，然后再发起recvform系统调用。重复上面的过程（polling）
- 直到数据准备好，再拷贝数据到进程，进行数据处理。但拷贝数据整个过程，进程仍然是属于阻塞的状态
- 优点：能够在等待任务完成的时间里干其他活了
- 缺点：任务完成的响应延迟增大了，轮询会消耗大量的CPU时间

### I/O 多路复用（ I/O multiplexing）

- 通过一种机制一个进程能同时等待多个I/O文件描述符
- 监视的方式分为select，poll，epoll
- select/epoll的好处就在于单个process就可以同时处理多个网络连接的IO
- 当用户进程调用了select，那么整个进程会被block，当任何一个socket中的数据准备好了，select就会返回
- select/epoll的优势并不是对于单个连接能处理得更快，而是在于能处理更多的连接
- I/O多路复用是阻塞在select，epoll这样的系统调用之上，而没有阻塞在真正的I/O系统调用如recvfrom之上
- I/O多路复用技术通过把多个I/O的阻塞复用到同一个select的阻塞上，从而使得系统在单线程的情况下可以同时处理多个客户端请求
- I/O多路复用的最大优势是系统开销小，系统不需要创建新的额外进程或者线程
- 同步阻塞

### 信号驱动式I/O（signal-driven I/O）

- 允许Socket进行信号驱动IO,并安装一个信号处理函数，进程继续运行并不阻塞
- 当数据准备好时，进程会收到一个SIGIO信号，可以在信号处理函数中调用I/O操作函数处理数据

### 异步非阻塞 I/O（asynchronous I/O）

- 异步IO不是顺序执行
- 用户进程进行aio_read系统调用之后，无论内核数据是否准备好，都会直接返回给用户进程，然后用户态进程可以去做别的事情
- kernel会等待数据准备完成，然后将数据拷贝到用户内存
- 然后kernel会给用户进程发送一个signal或执行一个基于线程的回调函数来完成这次 I/O 处理过程
- IO两个阶段，进程都是非阻塞的
- 异步IO库，例如libevent、libev、libuv

### 异步阻塞

- 使用阻塞的理由一：某些场景下，需要顺序执行
- 使用阻塞的理由二：降低响应延迟。异步任务处理完成时，进程可能在处理其他任务中，任务切换需要时间
- 为了在异步环境里模拟 “顺序执行” 的效果，就需要把同步代码转换成异步形式，这称为 CPS（Continuation Passing Style）变换
- 用户只需用同步方式编写代码，CPS变换器会把它转换成层层嵌套的异步回调形式

## 总结

- blocking I/O，non-blocking I/O，I/O multiplexing都属于synchronous I/O
- SIGIO，AIO 都属于asynchronous I/O
- select/poll/epoll本身是同步的，可以阻塞也可以不阻塞

### select

- int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout)
- 可以设置timeval决定该系统调用是否阻塞，设置为NULL即阻塞，为0不阻塞

### poll

- int poll(struct pollfd *fds, nfds_t nfds, int timeout)
- 可以通过指定timeout的值来决定是否阻塞（当timeout＜0时，会无限期阻塞；当timeout=0时，会立即返回）

### epoll

- epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout)
- 可以通过指定timeout来指定该调用是否阻塞（当timeout=-1时，会无限期阻塞；当timeout=0时，会立即返回；>-时，阻塞T毫秒）
