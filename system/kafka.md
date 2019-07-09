## 什么是 kafka ？
[kafka 安装](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-18-04)  

[伪分布式集群搭建](https://kafka.apache.org/08/documentation/#quickstart)搭建一个单节点的 kafka 伪分布式集群

#### 问题
搭建多节点过程发生异常`OpenJDK 64-Bit Server VM warning: INFO: os::commit_memory(0x00000000c0000000, 1073741824, 0) failed; error='Cannot allocate memory' (errno=12)`，原因是我的虚拟机内存是1G的，需要修改`kafka-server-start.sh`调小脚本的`KAFKA_HEAP_OPTS`。

## kafka 原理
![](https://user-gold-cdn.xitu.io/2019/6/30/16ba68eca692b99a?w=1128&h=506&f=png&s=226758)

## 概念
todo 消费规则？如何获取消息？同一分组不能消费同分区的消息
#### Broker
已发布的消息保存在一组服务器中，称之为Kafka集群。集群中的每一个服务器都是一个代理(Broker). 消费者可以订阅一个或多个主题（topic），并从Broker拉数据，从而消费这些已发布的消息。  
#### Topic
Kafka将消息种子(Feed)分门别类，每一类的消息称之为一个主题(Topic).  

![](https://user-gold-cdn.xitu.io/2019/7/1/16bac34a3c573225?w=416&h=267&f=png&s=19550) 

每一个分区都是一个顺序的、不可变的消息队列， 并且可以持续的添加。分区中的消息都被分了一个序列号，称之为偏移量(offset)，在每个分区中此偏移量都是唯一的。

Kafka集群保持所有的消息，直到它们过期， 无论消息是否被消费了。 实际上消费者所持有的仅有的元数据就是这个偏移量，也就是消费者在这个log中的位置。

Partition（分区）: 
Log的分区被分布到集群中的多个服务器上。每个服务器处理它分到的分区。 根据配置每个分区还可以复制到其它服务器作为备份容错。 每个分区有一个leader，零或多个follower。Leader处理此分区的所有的读写请求，而follower被动的复制数据。如果leader宕机，其它的一个follower会被推举为新的leader。 一台服务器可能同时是一个分区的leader，另一个分区的follower。 这样可以平衡负载，避免所有的请求都只让一台或者某几台服务器处理。

Kafka中采用分区的设计有几个目的。
- 可以处理更多的消息，不受单台服务器的限制。Topic拥有多个分区意味着它可以不受限的处理更多的数据。
- 分区可以作为并行处理的单元。

#### Producer
生产者：往 Topic 生产消息的时候，存在多个分区怎么选择呢？这里如果消息没有指定 partition 的话，就需要依赖分区器了，分区器根据key这个字段来计算 partition 的值。分区器的作用就是为消息分配分区。
生产者分区器：
- DefaultPartitioner：默认方式，轮询
- 也可以实现 Partitioner 接口自定义分区器

#### Consumer
消费消息 消费者 
如何保证消息的顺序消费？顺序处理 Topic 的所有消息，那就只提供一个分区，这样在分区内可以保证顺序。
消费的两种模式：
- 点对点模式（队列）：消费者都在同一组中
- 发布/订阅模式：消费者都在不同的组中

消费者分区分配:
Kafka的默认规则，每一个分区只能被同一个消费组中的一个消费者消费。消费者的分区分配是指为消费组中的消费者分配所订阅主题中的分区。

![](https://user-gold-cdn.xitu.io/2019/7/2/16bae51e958a3d24?w=1154&h=656&f=png&s=132261)

同一个消费组的消费者不能同时消费订阅主题的同一个分区。那么这里就存在一个消费者的分区分配规则，也可以通过 ParitionAssignor 接口来自定义分区分配策略：
- RangeAssignor：默认规则
- RoundRobinAssignor：
- StickyAssignor：


消费关系图：
![](https://user-gold-cdn.xitu.io/2019/7/1/16bac39b1512d406?w=2041&h=1243&f=png&s=139259)

## 数据读取
offset: 老版本存储在 zk，新版本存储在本地。为什么呢？

## 集群管理依赖 zookeeper 及 架构图
![](https://user-gold-cdn.xitu.io/2019/6/30/16ba6a7706c0d1fc?w=1848&h=930&f=png&s=679133)
选举机制？  
leader/follower:  
partition:  
replication:  

## 核心接口
Producer API 允许应用程序发送数据流到kafka集群中的topic。
Consumer API 允许应用程序从kafka集群的topic中读取数据流。
Streams API 允许从输入topic转换数据流到输出topic。
Connect API 通过实现连接器（connector），不断地从一些源系统或应用程序中拉取数据到kafka，或从kafka提交数据到宿系统（sink system）或应用程序。

## 生产过程 数据如何写入topic？
写入方式: 生产者采用 push 发布到 broker，消息被追加到 partition 中，消息在分区中是有序的，顺序存储保证了吞吐率。  
分区: 发送消息被发送到 topic，消息分区内有序，todo   为什么分区？分区原则?  
副本:  
写入流程: 如何确保生产者不丢数据？ack 机制 ack=all。
![](https://user-gold-cdn.xitu.io/2019/6/30/16ba6c3f58a7b6a7?w=1053&h=512&f=png&s=164281)

## 存储
存储方式:  
存储策略:  
zk存储结构:  

## 消费过程
todo 如何避免重复消费的？  
高级api:  
低级api:  


## 训练题
- kafka节点之间如何复制备份的？  

- kafka消息是否会丢失？为什么？  

- kafka最合理的配置是什么？  

- kafka的leader选举机制是什么？ 

- kafka对硬件的配置有什么要求？ 

- kafka的消息保证有几种方式？  

- kafka为什么会丢消息？  
leader 选举的时候可能会有数据丢失，但是committed的消息保证不会丢失。

- Kafka的用途有哪些？使用场景如何？
- Kafka中的ISR、AR又代表什么？ISR的伸缩又指什么
- Kafka中的HW、LEO、LSO、LW等分别代表什么？
- Kafka中是怎么体现消息顺序性的？
- Kafka中的分区器、序列化器、拦截器是否了解？它们之间的处理顺序是什么？
- Kafka生产者客户端的整体结构是什么样子的？
- Kafka生产者客户端中使用了几个线程来处理？分别是什么？
- Kafka的旧版Scala的消费者客户端的设计有什么缺陷？
- “消费组中的消费者个数如果超过topic的分区，那么就会有消费者消费不到数据”这句话是否正确？如果正确- ，那么有没有什么hack的手段？
- 消费者提交消费位移时提交的是当前消费到的最新消息的offset还是offset+1?
- 有哪些情形会造成重复消费？
- 那些情景下会造成消息漏消费？
- KafkaConsumer是非线程安全的，那么怎么样实现多线程消费？
- 简述消费者与消费组之间的关系
- 当你使用kafka-topics.sh创建（删除）了一个topic之后，Kafka背后会执行什么逻辑？
- topic的分区数可不可以增加？如果可以怎么增加？如果不可以，那又是为什么？
- topic的分区数可不可以减少？如果可以怎么减少？如果不可以，那又是为什么？
- 创建topic时如何选择合适的分区数？
- Kafka目前有那些内部topic，它们都有什么特征？各自的作用又是什么？
- 优先副本是什么？它有什么特殊的作用？
- Kafka有哪几处地方有分区分配的概念？简述大致的过程及原理
- 简述Kafka的日志目录结构
- Kafka中有那些索引文件？
- 如果我指定了一个offset，Kafka怎么查找到对应的消息？
- 如果我指定了一个timestamp，Kafka怎么查找到对应的消息？
- 聊一聊你对Kafka的Log Retention的理解
- 聊一聊你对Kafka的Log Compaction的理解
- 聊一聊你对Kafka底层存储的理解（页缓存、内核层、块层、设备层）
- 聊一聊Kafka的延时操作的原理
- 聊一聊Kafka控制器的作用
- 消费再均衡的原理是什么？（提示：消费者协调器和消费组协调器）
- Kafka中的幂等是怎么实现的
- Kafka中的事务是怎么实现的（这题我去面试6加被问4次，照着答案念也要念十几分钟，面试官简直凑不要脸- ）
- Kafka中有那些地方需要选举？这些地方的选举策略又有哪些？
- 失效副本是指什么？有那些应对措施？
- 多副本下，各个副本中的HW和LEO的演变过程
- 为什么Kafka不支持读写分离？
- Kafka在可靠性方面做了哪些改进？（HW, LeaderEpoch）
- Kafka中怎么实现死信队列和重试队列？
- Kafka中的延迟队列怎么实现（这题被问的比事务那题还要多！！！听说你会Kafka，那你说说延迟队列怎么实- 现？）
- Kafka中怎么做消息审计？
- Kafka中怎么做消息轨迹？
- Kafka中有那些配置参数比较有意思？聊一聊你的看法
- Kafka中有那些命名比较有意思？聊一聊你的看法
- Kafka有哪些指标需要着重关注？
- 怎么计算Lag？(注意read_uncommitted和read_committed状态下的不同)
- Kafka的那些设计让它有如此高的性能？
- Kafka有什么优缺点？
- 还用过什么同质类的其它产品，与Kafka相比有什么优缺点？
- 为什么选择Kafka?
- 在使用Kafka的过程中遇到过什么困难？怎么解决的？
- 怎么样才能确保Kafka极大程度上的可靠性？
- 聊一聊你对Kafka生态的理解

## Demo
producer 代码
```go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.99.102"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "my-replicated-topic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

```

consumer 代码
```go
package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "192.168.99.102",
        "group.id":          "te-1",
        "auto.offset.reset": "earliest",
    })

    if err != nil {
        panic(err)
    }

    c.SubscribeTopics([]string{"my-replicated-topic", "^aRegex.*[Tt]opic"}, nil)

    for {
        msg, err := c.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
        } else {
            // The client will automatically try to recover from all errors.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }

    c.Close()
}

```

## 参考资料
[orchome kafka 中文教程](https://www.orchome.com/kafka/index)  
[bilibili 尚硅谷 kafka 视频教程](https://www.bilibili.com/video/av35354301?from=search&seid=8195427520893763824)  
[掘金朱小厮专栏](https://juejin.im/user/5baf7ec26fb9a05cff32266e/posts)