# TiKV

## TiKV整体架构

![Alt](./image/tikv架构.jpg#pic_center)

TiDB中的TiDB server会接收用户请求,然后生成执行计划,该执行计划会发到TiKV中执行,并将对应的数据持久化.

## 1.数据持久化
> 数据持久化是由TiKV Node中的RocksDB来实现的.数据以region为单位,存储在各个分布式节点中,其中我们用raft一致性算法来实现分布式下节点数据的强一致性.


### 1.1RocksDB是什么？


### 1.2RocksDB的数据写入?
![Alt](./image/tikvwrite.png#pic_center)

+ 内存向磁盘中写入数据
```c++
---------------写入数据---------------
1.PUT(1,Tom)

2.WAL,Write Ahead Log.先将该操作记录在磁盘日志中,防止内存中的数据MemTable因掉电而丢失.如果内存中的数据因掉电而丢失,可以在故障恢复的时候读取WAL,重新恢复数据.
2.1.当我们写WAL时,由于写文件,操作系统层面也有缓存,这就可能会出现,操作系统的缓存还没有被批量刷新到磁盘中,系统掉电,数据丢失,所以这里提供一个参数sync_log,将sync_log设置为true时,我们可以直接将写入压入磁盘中,避免经过操作系统的缓存.

3.将数据追加写入MemTable中,如果MemTable的数据超过了一定的数据量(write_buffer_size),则将MemTable转存到immutable MemTable(依然在内存中)(可以防止写阻塞),RocksDB会重新开辟一个MemTable.
3.1.immutable MemTable是固定的数据,可以防止MemTable写阻塞,后台会开启一个线程,将immutable MemTable刷新到磁盘中.
3.2.immutable MemTable达到5个则会使得RicksDB触发write stall,进行限速,即客户端写入速度变慢.(生产者消费者速率不匹配问题)

4.当MemTable 的 flush disk成功后,WAL就可以被覆盖了.
```
![Alt](./image/write2.jpg#pic_center)

+ 磁盘中文件的组织
```c++
1.内存中的immutable MemTable会通过后台进程flush到磁盘中,亦即是图中的Level 0层的文件.
当Level 0层的文件数达到4个的时候,就会触发一个compaction操作,将上层的一定数量的文件压缩合并成一个文件,并移动到下一层,以此类推.其中的文件会通过key值进行排序,所以在文件中查找数据的时候,可以采用二分查找法.
```

+ **和b+树相比?**

从写入流程来看,RocksDB写入一个数据首先要写一个WAL,需要一次磁盘io,还需要写一次内存,即将数据写入MemTable中,剩余的操作只需要后台进程来完成就可以了.

如果采用b+树,则首先需要多次磁盘io找到对应的页面,所以上面的性能更好一点.

**实际相当于随机写变为了顺序写**

+ **更新和删除操作**
不需要找到对应的数据进行删除,只需要将delete/update操作写入MemTable中就可以,一定的时间段内,最新的数据一定会再Level最高层,也会最先被找到.


+ **查询操作**

![Alt](./image/find.jpg#pic_center)



```c++
1.首先查询Block Cache,Block Cache中保存着最近最常使读的数据,所以大概率会命中该缓存.如果缓存未命中,则下一步.
2.查询MemTable,看MemTable中是否有数据,没有则继续查找immutable MemTable.否则进入下一步.
3.从磁盘中的,Level 0开始一直往下,查找对应的数据.最新的数据永远在高层.
```
> 如果文件的数据量很大,如果快速判断文件中是否有数据,可以给每个文件加一个布隆过滤器,布隆过滤器对于判断数据存在有一定的误判率,但是可以百分百确认数据不在某个文件中.由此可以加速查询.


+ **Column Families列簇**
>  列簇可以理解为数据分片技术.(分内存和STT文件)
>　列簇之间共享一份写前日志WAL.
>  rocksdb中有默认的列簇.

![Alt](./image/column_Families.jpg#pic_center)


举例来理解列簇:原本在一个RocksDB中的两张表,id-(name,age) ,id-(addr,tele).如果采用列簇则可以分到两个Column Families.

```c++
则我们写入数据时,需要指定对应的列簇.
write({cf1,id(name,age)})
write({cf2,id(addr,tele)})
```

> 一次磁盘io一次内存io
>MemTable采用LSM或者跳表.
>write stall
>压缩compaction


## 2.分布式事务

![Alt](./image/tran1.jpg#pic_center)

```c++
三个列簇,一个Column Families保存一类键值对.
Default Column Families保存真正的数据.
Lock Column Families保存锁数据.
Write Column Families保存事务的提交信息.
```

```
1. preWrite:start timestamp=100,开启一个时间戳,标志者事务的开始.


```



+ 两阶段提交









Coprocessor协同处理器(算子下推)

