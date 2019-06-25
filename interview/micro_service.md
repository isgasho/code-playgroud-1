# go-micro 读书笔记
概览
Go Micro提供分布式系统开发的核心库，包含RPC与事件驱动的通信机制。

micro的设计哲学是可插拔的架构理念，她提供可快速构建系统的组件，并且可以根据自身的需求剥离默认实现并自行定制。


## 特性
Go Micro 抽象了分布式系统的细节。以下是主要功能。

#### 服务发现（Service Discovery） - 自动服务注册与名称解析。服务发现是微服务开发中的核心。当服务A要与服务B协作时，它得知道B在哪里。默认的服务发现系统是Consul，而multicast DNS (mdns，组播)机制作为本地解决方案，或者零依赖的P2P网络中的SWIM协议（gossip）。
服务发现常用的框架 
- zookeeper
- eureka
- etcd
- consul go-micro 默认



- 负载均衡（Load Balancing） - 在服务发现之上构建了负载均衡机制。当我们得到一个服务的任意多个的实例节点时，我们要一个机制去决定要路由到哪一个节点。我们使用随机处理过的哈希负载均衡机制来保证对服务请求颁发的均匀分布，并且在发生问题时进行重试。
- 消息编码（Message Encoding） - 支持基于内容类型（content-type）动态编码消息。客户端和服务端会一起使用content-type的格式来对Go进行无缝编/解码。各种各样的消息被编码会发送到不同的客户端，客户端服服务端默认会处理这些消息。content-type默认包含proto-rpc和json-rpc。
- Request/Response - RPC通信基于支持双向流的请求/响应方式，我们提供有抽象的同步通信机制。请求发送到服务时，会自动解析、负载均衡、拨号、转成字节流，默认的传输协议是http/1.1，而tls下使用http2协议。
- 异步消息（Async Messaging） - 发布订阅（PubSub）头等功能内置在异步通信与事件驱动架构中。事件通知在微服务开发中处于核心位置。默认的消息传送使用点到点http/1.1，激活tls时则使用http2。
- 可插拔接口（Pluggable Interfaces） - Go Micro为每个分布式系统抽象出接口。因此，Go Micro的接口都是可插拔的，允许其在运行时不可知的情况下仍可支持。所以只要实现接口，可以在内部使用任何的技术。更多插件请参考：github.com/micro/go-plugins。

