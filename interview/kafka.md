## 什么是 kafka 
[kafka 安装](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-18-04)  

[伪分布式集群搭建](https://kafka.apache.org/08/documentation/#quickstart)搭建一个单节点的 kafka 伪分布式集群

#### 问题
搭建多节点过程发生异常`OpenJDK 64-Bit Server VM warning: INFO: os::commit_memory(0x00000000c0000000, 1073741824, 0) failed; error='Cannot allocate memory' (errno=12)`，原因是我的虚拟机内存是 1G 的，需要修改`kafka-server-start.sh`调小脚本的`KAFKA_HEAP_OPTS`。

## kafka 原理
![](https://user-gold-cdn.xitu.io/2019/6/30/16ba68eca692b99a?w=1128&h=506&f=png&s=226758)

## 概念
todo 消费规则？如何获取消息？同一分组不能消费同分区的消息
#### Broker
已发布的消息保存在一组服务器中，称之为 Kafka 集群。集群中的每一个服务器都是一个代理(Broker). 消费者可以订阅一个或多个主题（topic），并从 Broker 拉数据，从而消费这些已发布的消息。  
#### Topic
Kafka 将消息种子(Feed)分门别类，每一类的消息称之为一个主题(Topic).  

![](https://user-gold-cdn.xitu.io/2019/7/1/16bac34a3c573225?w=416&h=267&f=png&s=19550) 

每一个分区都是一个顺序的、不可变的消息队列， 并且可以持续的添加。分区中的消息都被分了一个序列号，称之为偏移量(offset)，在每个分区中此偏移量都是唯一的。

Kafka 集群保持所有的消息，直到它们过期， 无论消息是否被消费了。 实际上消费者所持有的仅有的元数据就是这个偏移量，也就是消费者在这个 log 中的位置。

Partition（分区）: 
Log 的分区被分布到集群中的多个服务器上。每个服务器处理它分到的分区。 根据配置每个分区还可以复制到其它服务器作为备份容错。 每个分区有一个 leader，零或多个 follower。Leader 处理此分区的所有的读写请求，而 follower 被动的复制数据。如果 leader 宕机，其它的一个 follower 会被推举为新的 leader。 一台服务器可能同时是一个分区的 leader，另一个分区的 follower。 这样可以平衡负载，避免所有的请求都只让一台或者某几台服务器处理。

Kafka 中采用分区的设计有几个目的。
- 可以处理更多的消息，不受单台服务器的限制。Topic 拥有多个分区意味着它可以不受限的处理更多的数据。
- 分区可以作为并行处理的单元。

#### Producer
生产者：往 Topic 生产消息，可以选分区，开发者可以选择选分区的算法，最简单的是轮流选择。

#### Consumer
消费消息 消费者 
如何保证消息的顺序消费？顺序处理 Topic 的所有消息，那就只提供一个分区，这样在分区内可以保证顺序。
消费的两种模式：
- 点对点模式（队列）：消费者都在同一组中
- 发布/订阅模式：消费者都在不同的组中

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
Producer API 允许应用程序发送数据流到 kafka 集群中的 topic。
Consumer API 允许应用程序从 kafka 集群的 topic 中读取数据流。
Streams API 允许从输入 topic 转换数据流到输出 topic。
Connect API 通过实现连接器（connector），不断地从一些源系统或应用程序中拉取数据到 kafka，或从 kafka 提交数据到宿系统（sink system）或应用程序。

## 生产过程 数据如何写入 topic
写入方式: 生产者采用 push 发布到 broker，消息被追加到 partition 中，消息在分区中是有序的，顺序存储保证了吞吐率。  
分区: 发送消息被发送到 topic，消息分区内有序，todo   为什么分区？分区原则?  
副本:  
写入流程: 如何确保生产者不丢数据？ack 机制 ack=all。
![](https://user-gold-cdn.xitu.io/2019/6/30/16ba6c3f58a7b6a7?w=1053&h=512&f=png&s=164281)

## 存储
存储方式:  
存储策略:  
zk 存储结构:  

## 消费过程
todo 如何避免重复消费的？  
高级 api:  
低级 api:  


## 疑问
kafka 节点之间如何复制备份的？  

kafka 消息是否会丢失？为什么？  

kafka 最合理的配置是什么？  

kafka 的 leader 选举机制是什么？ 

kafka 对硬件的配置有什么要求？ 

kafka 的消息保证有几种方式？  

- kafka 为什么会丢消息？  
leader 选举的时候可能会有数据丢失，但是 committed 的消息保证不会丢失。

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