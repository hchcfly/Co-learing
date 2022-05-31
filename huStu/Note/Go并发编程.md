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

```go
channel在写少读多的情况下会出现deadlock.

channel关闭,读不到数据就默认为0

记得关闭通道.
```


## Golang WaitGroup

> 相当于muduo中的 CountDownLatch
> 实现两个协程间互相等待(同步)


## Golang runtime

```go

runtime.Gosched() //让出CPU时间片,重新等待安排任务

runtime.Goexit()  //退出当前协程

runtime.GOMAXPROCS() //设置最大核心数

```


## Golang select

> select是Go中的一个控制结构,类似于switch语句,用于处理异步IO操作.select会监听case语句中channel的读写操作。当case中的channel可读或可写时,将会触发相应的动作.

> select中的case语句必须是一个channel操作
> select中的default子句总是可运行的

```
1.如果有多个case都可以运行,Select会随机公平的选出一个执行,其他不会执行.
2.如果没有可运行的select语句,且有default语句,那么就会执行default的动作.
3.如果没有可运行的case语句,且没有default语句,select将阻塞,直到某个case通信可以运行.

```