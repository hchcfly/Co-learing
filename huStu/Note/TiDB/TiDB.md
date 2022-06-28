# TiKV

## TiKV整体架构

![Alt](./image/tikv架构.jpg#pic_center)

TiDB中的TiDB server会接收用户请求,然后生成执行计划,该执行计划会发到TiKV中执行,并将对应的数据持久化.

## 1.数据持久化
> 数据持久化是由TiKV Node中的RocksDB来实现的.


### 1.1RocksDB是什么？


### 1.2RocksDB的数据写入?
![Alt](./image/tikvwrite.png#pic_center)

```c++
---------------写入数据---------------
1.PUT(1,Tom)

2.WAL,Write Ahead Log.先将该操作记录在磁盘日志中,防止内存中的数据MemTable因掉电而丢失.如果内存中的数据因掉电而丢失,可以在故障恢复的时候读取WAL,重新恢复数据.
2.1.当我们写WAL时,由于写文件,操作系统层面也有缓存,这就可能会出现,操作系统的缓存还没有被批量刷新到磁盘中,系统掉电,数据丢失,所以这里提供一个参数sync_log,将sync_log设置为true时,我们可以直接将写入压入磁盘中,避免经过操作系统的缓存.

3.将数据追加写入MemTable中,如果MemTable的数据超过了一定的数据量(write_buffer_size),则将MemTable转存到immutable MemTable(依然在内存中)(可以防止写阻塞),RocksDB会重新开辟一个MemTable.
3.1.immutable MemTable是固定的数据,可以防止MemTable写阻塞,后台会开启一个线程,将immutable MemTable刷新到磁盘中.
3.2.immutable MemTable达到5个则会使得RicksDB触发write stall,进行限速,即客户端写入速度变慢.

4.当MemTable 的 flush disk成功后,WAL就可以被覆盖了.
```




> 一次磁盘io一次内存io
>MemTable采用LSM或者跳表.
>write stall
>压缩compaction







region

## CF列簇 Column Families 数据分片技术 ##
有默认列簇
数据分片但是日志没有分


Coprocessor协同处理器(算子下推)

