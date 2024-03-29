## 分布式 ID 方案
- 全局唯一性：不能出现重复的ID号，既然是唯一标识，这是最基本的要求。
- 趋势递增：在MySQL InnoDB引擎中使用的是聚集索引，由于多数RDBMS使用B-tree的数据结构来存储索引数据，在主键的选择上面我们应该尽量使用有序的主键保证写入性能。
- 单调递增：保证下一个ID一定大于上一个ID，例如事务版本号、IM增量消息、排序等特殊需求。
- 信息安全：如果ID是连续的，恶意用户的扒取工作就非常容易做了，直接按照顺序下载指定URL即可；如果是订单号就更危险了，竞争对手可以直接知道我们一天的单量。所以在一些应用场景下，会需要ID无规则、不规则。

上述123对应三类不同的场景，3和4需求还是互斥的，无法使用同一个方案满足。

同时除了对ID号码自身的要求，业务还对ID号生成系统的可用性要求极高，想象一下，如果ID生成系统瘫痪，整个美团点评支付、优惠券发券、骑手派单等关键动作都无法执行，这就会带来一场灾难。

由此总结下一个ID生成系统应该做到如下几点：

- 平均延迟和TP999延迟都要尽可能低；
- 可用性5个9；
- 高QPS。

### 雪花算法
除了最高位 bit 标记为不可用以外，其余三组 bit 占位均可浮动，看具体的业务需求而定。默认情况按下图分配，第一个 bit 不用，是因为二进制第一位为 1 代表负数。41bit 的时间戳到毫秒级别，2^41 = 2199023255552，一年有 365*24*3600*1000 毫秒，那么计算出来是 69.7 年。第三部分是 10bit 为 1024，也就是说有 1023 个号码可以用，最后是序列号 12bit，4096 个号码。那么在单机的情况下一秒可以确保生成 1000 * 4095 个 ID。

![](https://user-gold-cdn.xitu.io/2019/7/20/16c0e1ed6a461619?w=1021&h=346&f=png&s=54352)

#### 实现伪代码

### 注意问题
1. 当服务器同步时间时，可能会出现时间回拨，造成生成的 id 重复
2. 单机递增，但是分布式时间不一致，可能不会全局递增

### MongoDB 的 ObjectId
[MongoDB官方文档 ObjectID](https://docs.mongodb.com/manual/reference/method/ObjectId/#description)可以算作是和snowflake类似方法，通过“时间+机器码+pid+inc”共12个字节，通过4+3+2+3的方式最终标识成一个24长度的十六进制字符。

### 数据库自增
各数据库节点采用不同的步长来解决单点性能问题，但是扩容存在难题。  
优点
- 非常简单，利用现有数据库系统的功能实现，成本小，有DBA专业维护。
- ID号单调自增，可以实现一些对ID有特殊要求的业务。

缺点
- 强依赖DB，当DB异常时整个系统不可用，属于致命问题。配置主从复制可以尽可能的增加可用性，但是数据一致性在特殊情况下难以保证。主从切换时的不一致可能会导致重复发号。
- ID发号性能瓶颈限制在单台MySQL的读写性能。

## 参考资料
- [Leaf——美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)
- [www.lanindex.com](https://www.lanindex.com/twitter-snowflake%EF%BC%8C64%E4%BD%8D%E8%87%AA%E5%A2%9Eid%E7%AE%97%E6%B3%95%E8%AF%A6%E8%A7%A3/)
- [分布式 ID 方案](http://www.cnblogs.com/haoxinyue/p/5208136.html)
