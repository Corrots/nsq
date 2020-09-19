简介
nsq最初是由bitly公司开源出来的一款简单易用的消息中间件，它可用于大规模系统中的实时消息服务，并且每天能够处理数亿级别的消息。它有以下特性：
1. 分布式 — 它提供了分布式的、去中心化且没有单点故障的拓扑结构，稳定的消息传输发布保障，能够具有高容错和高可用特性
2. 易于扩展 — 它支持水平扩展，没有中心化的消息代理(Broker)，内置的发现服务让集群中增加节点非常容易。
3. 运维方便 — 它非常容易配置和部署，灵活性高。
4. 高度集成 — 现在已经有官方的Golang、Python和JavaScript客户端，社区也有了其他各个语言的客户端库方便接入，自定义客户端也非常容易。

## Setup
```bash
# MacOS
brew install nsq
# To have launchd start nsq now and restart at login:
brew services start nsq
# Or, if you don't want/need a background service you can just run:
nsqd -data-path=/usr/local/var/nsq
```

## 组件
### nsqlookupd
负责管理拓扑信息并提供最终一致性的发现服务的守护进程(daemon)
主要负责服务发现 负责nsqd的心跳、状态监测，给客户端、nsqadmin提供nsqd地址与状态

终端启动nsqlookupd：
```bash
> nsqlookupd
[nsqlookupd] 2020/09/18 16:48:33.275374 INFO: nsqlookupd v1.2.0 (built w/go1.13.5)
[nsqlookupd] 2020/09/18 16:48:33.276029 INFO: TCP: listening on [::]:4160
[nsqlookupd] 2020/09/18 16:48:33.276055 INFO: HTTP: listening on [::]:4161
```
**默认HTTP接口监听4161，TCP接口监听4160**

### nsqd
nsqd是一个负责接收、排队、投递消息给客户端的守护进程。客户端通过查询 nsqlookupd 来发现指定topic的nsqd生产者，nsqd节点会广播topic和channel信息。数据流模型如下：
[image:88470C8E-5272-461F-811B-4784355223CC-32942-0000CC0C881DAC3D/16c296b4c9505012.gif]
单个nsqd可以有多个topic，每个topic可以有多个channel。channel接收这个topic所有消息的副本，从而实现多播分发，而channel上的每个消息被分发给它的订阅者，从而实现负载均衡。

终端启动nsqd：
```bash
nsqd --lookupd-tcp-address=127.0.0.1:4160
[nsqd] 2020/09/18 16:59:22.691042 INFO: nsqd v1.2.0 (built w/go1.13.5)
[nsqd] 2020/09/18 16:59:22.699337 INFO: ID: 1023
[nsqd] 2020/09/18 16:59:22.705373 INFO: NSQ: persisting topic/channel metadata to nsqd.dat
[nsqd] 2020/09/18 16:59:22.713219 INFO: TCP: listening on [::]:4150
[nsqd] 2020/09/18 16:59:22.713290 INFO: HTTP: listening on [::]:4151
[nsqd] 2020/09/18 16:59:22.713426 INFO: LOOKUP(127.0.0.1:4160): adding peer
[nsqd] 2020/09/18 16:59:22.713448 INFO: LOOKUP connecting to 127.0.0.1:4160
```
**nsqd通过tcp端口连接到了nsqlookupd，它自己在4151接受HTTP请求，在4150接受TCP请求**

### nsqadmin
nsqadmin 是一套WEB管理UI，用来汇集集群的实时统计，并执行不同的管理任务。

终端启动nsqadmin：
```bash
nsqadmin --lookupd-http-address=127.0.0.1:4161
[nsqadmin] 2020/09/18 17:02:35.689163 INFO: nsqadmin v1.2.0 (built w/go1.13.5)
[nsqadmin] 2020/09/18 17:02:35.689952 INFO: HTTP: listening on [::]:4171
```

### 其他组件
* nsq_stat
* nsq_tail
* nsq_to_file

## cmd
```bash
# 发布消息到nsqd，基于REST API，没有topic会先创建之
curl -d 'message 1' 'http://127.0.0.1:4151/pub?topic=test'
# 
nsq_tail --lookupd-http-address=127.0.0.1:4161 --topic=test
#

```

test.DWMAC-C02VX372HV2L.2020-09-18_17.log

```bash
go get -u github.com/nsqio/go-nsq
```

