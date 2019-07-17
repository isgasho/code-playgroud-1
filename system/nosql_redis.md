## 基本数据格式

### 字符串
编码：int，raw，embstr
- int：保存的对象为整数，且可以用 long 类型表示。
- raw：保存的对象为字符串，且这个字符串的长度大于 32 字节，使用一个简单动态字符串（SDS）保存，并设置编码为 raw
- embstr：保存的对象为字符串，且字符串的长度小于等于 32 字节，则使用 embstr 编码保存。相较于 raw ，需要一次内存分配，内存地址连续可更好利用缓存。

int 和 embstr 编码的字符串在条件满足的情况下，会转为 raw 编码:
1. int 编码的字符串，如果执行命令使得这个对象保存的不再是整数，则转为 raw 编码
2. embstr 编码的对象为只读对象。如果对 embstr 对象编码的字符串进行修改操作，则转为 raw 编码。

### 列表
编码：ziplist 和 linkedlist
- ziplist：底层使用压缩列表实现
- linkedlist：底层使用双端链表实现，嵌套多个字符串对象。

当列表对象同时满足以下两个条件，使用 ziplist 编码，否则使用 linkedlist
1. 列表对象保存的所有字符串元素的长度都小于 64 字节
2. 列表对象保存的元素数量少于 512 个

### 哈希
编码：ziplist 和 hashtable
- hashtable：哈希对象使用字典作为底层实现，每个键值对使用一个字典键值对来保存

当哈希对象同时满足以下两个条件时，哈希使用 ziplist 编码
1. 哈希对象保存的所有键值对的键和值的字符串长度都小于 64 字节
2. 哈希对象保存的键值对数量小于 512 个

### 集合
编码：intset 和 hashtable
- intset：使用整数集合作为底层实现，集合对象包含的所有元素都保存在整数集合里

当集合对象可以同时满足以下两个条件时，对象使用 intset 编码
1. 集合对象保存的所有元素是整数
2. 集合对象保存的元素数量不超过 512 个

### rehash 策略
负载因子 = 哈希表已保存节点数量 / 哈希表大小 (lf = ht[0].used / ht[0].size)

##### rehash 步骤
1. 为字典的 ht[1] 哈希表分配空间。
   - 扩展：ht[1].size 为 第一个 2^n 大于 ht[0].used * 2 的值。没有执行 BGSAVE 或 BGREWRITEAOF 时，需要负载因子大于等于 1。在执行时， BGSAVE 或 BGREWRITEAOF，需要负载因子大于等于 5。
   - 收缩：ht[1].size 为第一个 2^n 大于 ht[0].used 的值。在负载因子小于 0.1 时执行。
2. 将 ht[0] 的键值对 rehash 到 ht[1] 中，rehash 是指重新计算哈希值和索引值，然后放到 ht[1] 的指定位置
3. 当 ht[0] 的值都迁移到 ht[1] 之后，释放 ht[0]，将 ht[1] 设置为 ht[0]，并在 ht[1] 创建一个空白哈希表。

##### 渐进式 rehash
为避免数据量大的时候，rehash 造成性能影响。rehash 不是一次性的，而是分多次，慢慢将 ht[0] 迁移到 ht[1]。

1. 分配 ht[1] 空间。字典同时拥有 ht[0] 和 ht[1] 两个哈希表。
2. 在字典中维持一个索引计数器 rehashidx，设置值为 0，表示 rehash 开始。
3. 在 rehash 期间，每次对字典的操作都会顺带 rehash 到 ht[1]，当此次 rehash 完成时将 rehashidx 值 +1。
4. 随着字典操作的进行，在某个时间点，ht[0] 全部 rehash 到 ht[1] ，这是将 rehashidx 置为 -1 代表 rehash 全部完成。

删除查找更新都会在两个哈希表上操作，新增的键值对会保存到 ht[1]，保证了 ht[0] 不会再增加。


### 有序集合
编码：ziplist 和 skiplist
- ziplist：使用两个挨在一起的压缩列表的节点保存，第一个保存成员，第二个保存分值。
- skiplist：使用 zset 结构作为底层实现，一个 zset 结构同时包含一个字典和一个跳表。

当有序集合同时满足以下两个条件，使用 ziplist 编码：
1. 有序集合元素数量小于 128 个
2. 有序集合保存的所有元素成员的长度都小于 64 字节

##### 跳表是什么？

##### 有序集合为什么要用跳跃表加字典来实现？
为了保证有序集合的查找和范围操作。跳跃表保证有序性，比如使用 ZRANK，ZRANGE等操作，字典可以是查询的时间复杂度为O(1)。

## 单机部分
### 数据库

##### 过期键删除策略
1. 定时删除：创建定时器。性能要求高，redis 没有考虑这种方案。
2. 惰性删除：在取出时对键进行过期检查，对内存不友好。
3. 定期删除：定时调用 activeExpireCycle 函数，分多次遍历各数据库，从数据库的 expires 字典中随机检查一部分键的过期时间，并删除过期键。

- 执行 SAVE 或者 BGSAVE 命令所产生的新 RDB 文件不包含过期键。
- 执行 BGREWRITEAOF 命令所产生的重写 AOF 文件不包含过期的键。

##### 内存淘汰策略
Redis的内存淘汰机制
- noeviction: 当达到内存限制并且客户端尝试执行可能导致使用更多内存的命令时返回错误（大多数写命令，但DEL和一些例外）。
- allkeys-lru：通过删除最近最少使用的（LRU）key，以便为添加的新数据腾出空间（常用）
- allkeys-random：随机删除 key。
- volatile-lru：通过删除最近最少使用的（LRU）key，但仅在过期集的 key 之间。
- volatile-random：随机删除 key，但只删除在过期集的 key。
- volatile-ttl：在过期集删除 key，并尝试首先使用较早的生存时间（TTL）的 key 进行删除。

### 如何保证主从数据库数据一致性？
从服务器发现过期键不会主动删除，而是等待主节点发出 DEL 命令，中心化？如何保证高可用呢？

### RDB 持久化
RDB 是一个压缩二进制文件，通过保存数据库中的键值对来记录数据库状态，会落盘。所以即使机器崩溃也能从文件中恢复数据。执行 SAVE 或者 BGSAVE 命令用于生成 RDB 文件。
- SAVE 命令会阻塞 redis 进程，导致不能对外服务
- BGSAVE 是派生一个新的子进程取创建 RDB 文件，而主进程继续对外服务，可以配置自动间隔保存机制，满足条件即调用一次 BGSAVE 命令。

### AOF 持久化
Redis 默认不开启 AOF 持久化，AOF（Append only file）通过保存服务器执行的写命令来记录数据库状态。 分为三个步骤：
- 命令追加
- 文件写入
- 文件同步

fsync 写入策略：
- always: 每个请求
- everysec：每隔一秒写入
- no：redis 服务器不负责写入

BGREWRITEAOF：命令用于异步执行一个 AOF（AppendOnly File） 文件重写操作。重写会创建一个当前 AOF 文件的体积优化版本。

### 持久化选择
- RDB：RDB 文件恢复数据快，且文件非常紧凑，适合数据备份。
使用子进程(bgsave命令)生成对数据备份，性能影响较小。缺点：通常每五分钟或更长时间创建一个RDB快照，会丢失该段时间数据。如果数据量大，fork 子进程可能会非常耗时，并且如果数据集非常大且CPU性能不佳，可能会导致Redis停止服务客户端几毫秒甚至一秒钟
- AOF：可以使用不同的fsync策略：使用fsync的默认策略，每秒写入的性能仍然很好（使用后台线程执行fsync，并且当没有fsync正在进行时，主线程将努力执行写入。）最多也只丢失一秒的写入。AOF 日志是叠加日志，因此如果停电，也没有损坏问题。即使由于某种原因（磁盘已满或其他原因）日志以半写命令结束，redis-check-aof 工具也能够轻松修复它。缺点：AOF 文件体积更大，即使重写后，文件依然要比 RDB 大得多。AOF 恢复是命令重放，所以速度也较慢。

在过去， Redis 用户通常会因为 RDB 持久化和 AOF 持久化之间不同的优缺点而陷入两难的选择当中：

RDB 持久化能够快速地储存和回复数据，但是在服务器停机时却会丢失大量数据；AOF 持久化能够有效地提高数据的安全性，但是在储存和恢复数据方面却要耗费大量的时间。为了让用户能够同时拥有上述两种持久化的优点， Redis 4.0 推出了一个 RDB-AOF 的混合持久化方案： 
-  这种持久化能够通过 AOF 重写操作创建出一个同时包含 RDB 数据和 AOF 数据的 AOF文件， 其中 RDB 数据位于 AOF 文件的开头，它们储存了服务器开始执行重写操作时的数据库状态：至于那些在重写操作执行之后执行的 Redis 命令， 则会继续以 AOF 格式追加到 AOF 文件的末尾， 也即是 RDB 数据之后。

## 分布式

### 主从复制

### 集群


## 参考资料
##### 全栈群每周知识点|本周：Redis
##### 入门
- Redis数据类型的介绍
    - https://redis.io/topics/data-types-intro
    - https://redis.io/topics/data-types
- FAQ: https://redis.io/topics/faq
- 命令：https://redis.io/commands
-redis入门教程：https://www.runoob.com/redis/redis-tutorial.html
##### 进阶
1. 锁
- 谈谈Redis的SETNX：https://huoding.com/2015/09/14/463
- 基于Redis的分布式锁到底安全吗？
    - http://zhangtielei.com/posts/blog-redlock-reasoning.html
    - http://zhangtielei.com/posts/blog-redlock-reasoning-part2.html
- redis分布式锁：
    - 英文版：https://redis.io/topics/distlock
    - 中文版：https://cloudfeng.github.io/2018/07/22/arts/review/R-Distributed-locks-with-redis/  
- Martin Kleppmann对redis 分布式锁的分布：http://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html
- Redis作者Antirez的回应：http://antirez.com/news/101
2. 高可用与高性能
- 复制：https://redis.io/topics/replication
- Redis 哨兵模式：https://redis.io/topics/replication
- Redis 集群：
    - https://redis.io/topics/replication
    - https://redis.io/topics/cluster-spec
    - Redis 集群管理常见操作一览
：http://blog.huangz.me/2018/redis-cluster-manage-cheatsheet.html
- redis高可用原理：https://www.codedump.info/post/20190409-redis-sentinel/
3. 持久化
- RDB & AOF：https://redis.io/topics/persistence
- RDB-AOF 混合持久化：http://blog.huangz.me/2017/redis-rdb-aof-mixed-persistence.html
-过期删除策略
  -https://redis.io/topics/lru-cache
4.管道
-https://redis.io/topics/pipelining
-http://mattcamilli.com/glory-of-redis-pipelines.html
##### 客户端
- https://redis.io/clients
##### 源码
- 如何阅读Redis源码
    - http://zhangtielei.com/posts/blog-redis-how-to-start.html
    - http://blog.huangz.me/diary/2014/how-to-read-redis-source-code.html
- Redis 设计与实现：http://origin.redisbook.com/
- https://github.com/huangz1990/annotated_redis_source
- https://github.com/menwengit/redis_source_annotation
- 张铁蕾老师的Redis系列：https://mp.weixin.qq.com/s/3TU9qxHJyxHJgVDaYXoluA
##### 开发规范
- 阿里云Redis开发规范： https://yq.aliyun.com/articles/531067
##### 书籍
- Redis In Action： http://redisinaction.com/
- Redis 设计与实现：http://redisbook.com/
- Redis开发运维实践指南： https://legacy.gitbook.com/book/gnuhpc/redis-all-about/details
- Redis开发与运维：https://cachecloud.github.io/2016/10/24/Redis3%E5%BC%80%E5%8F%91%E8%BF%90%E7%BB%B4%E6%9C%80%E4%BD%B3%E5%AE%9E%E8%B7%B5-%E7%9B%AE%E5%BD%95/
##### 社区
- https://redis.io/community
编辑：云枫
