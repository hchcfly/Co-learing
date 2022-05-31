# Go并发编程

## Go 协程routine
> Golang中的并发是函数相互独立运行的能力,Goroutines是并发运行的函数.

+ 创建一个协程 (os:进程,线程,协程)

```go
go task()
```

## Go channel
>作为协程间通信的一种机制.

1.数据在通道上传递,在任何给定时间只有一个goroutine可以访问数据项,所以按照设计不会发生数据竞争.


2.channel分类:


    2.1有缓冲channel(异步通信)
    2.2无缓冲channel(同步通信)


## Golang WaitGroup

> 实现两个协程间互相等待(同步)


