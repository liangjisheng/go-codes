# stream

[stream](https://blog.csdn.net/m0_63463510/article/details/138763773)

Redis Stream 是在Redis 5.0版本中引入的一种新的数据结构，它主要用于实时数据处理场景，
如消息队列、日志记录和实时数据分析等。Stream的设计灵感部分来源于消息队列系统，如Kafka，
但它提供了更直接集成到Redis生态系统中的能力。

Redis Stream 类型本身设计时就充分考虑了生产者消费者模型的需求。它不仅包含了这一模型，
还对其进行了优化和扩展，以便更好地适应现代分布式系统中的消息传递场景。

生产者：任何可以向Redis Stream写入消息的客户端都可以视为生产者。生产者使用XADD命令向
Stream中添加消息，每个消息都附有一个全局唯一的ID，确保消息的顺序性和可追踪性。

消费者：消费者使用XREAD或XREADGROUP命令从Stream中读取消息。特别是XREADGROUP命令，
它支持消费者组（Consumer Group）的概念，这是生产者消费者模型中的一个重要组成部分。
消费者组让多个消费者可以协作处理Stream中的消息，同时保证了消息不会被重复处理。

消费者组：Stream支持消费者组，组内的消费者共享消息，但每个消息只由组内的一个消费者处理，
从而实现了消息的有序和公平分配。消费者组还允许消费者在处理失败时重新分配消息，以及通过XACK
命令确认消息已成功处理，实现消息的确认机制。

消息持久化与顺序性：Stream中的消息是持久化的，确保了即使在Redis服务器重启后，消息也不会丢失。
同时，Stream保持消息的严格顺序，这对于某些依赖消息顺序的应用场景至关重要。

阻塞读取与自定义读取策略：消费者可以选择阻塞读取模式，这意味着当没有新消息时，消费者会等待
直至新消息到达。此外，还可以通过参数定制读取的起始位置、消息数量或时间范围，提供了高度的灵活性。

Stream类型主要特点：

有序性：Stream中的消息按照ID排序，每个消息都有一个全局唯一的ID，确保了消息的顺序。

持久化：Stream中的数据是持久化的，即使Redis服务器重启，消息也不会丢失。

多播与分组消费：支持多个消费者同时消费同一流中的消息，而且可以将消费者组织成消费组，实现消息的分组消费，
每个消息可以被一个或多个组消费，但组内每个消息只会被其中一个消费者消费（类似于Kafka的分区消费者模型）。

灵活的数据结构：每条消息可以包含多个字段（field-value对），提供了高度的灵活性来携带复杂的数据。

消费者进度跟踪：消费者可以在读取消息时自动追踪自己的消费进度，Redis使用Last Seen Displacement(LSD)
来跟踪消费者的读取位置。

读取控制：支持多种读取模式，包括从特定消息ID读取、读取最近的N个消息、读取某个时间范围内的消息等。

阻塞读取：可以使用XREAD和XREADGROUP命令以阻塞的方式等待新消息，直到有新消息到达或超时。

XADD key [NOMKSTREAM] [MAXLEN|MINID [=|~] threshold [LIMIT count]] *|id field value [field value ...]
向Stream（key）中添加一条消息，ID可以是自动生成或指定的唯一标识符，后面跟着一个或多个字段值对。

这个命令有三个注意点:

消息id要比上个id大
默认用*表示自动生成 id
*:用于XADD命令,表示让系统自动生成id(类似于MySQL的自增主键)
生成的消息ID，有两部分组成，毫秒时间戳-该毫秒内产生的第1条消息

```redis
127.0.0.1:6379> xadd k1 * name zhangsan age 18
1722134003480-0 # 系统自动生成的id
127.0.0.1:6379> xadd k1 * name lisi age 19
1722134011535-0
127.0.0.1:6379> xadd k1 * name wangwu age 20
1722134018763-0
127.0.0.1:6379> xadd k1 1722134018763-0 name wangwu age 20 # 重复的id会出错
ERR The ID specified in XADD is equal or smaller than the target stream top item
127.0.0.1:6379> xadd k1 1722134018763-1 name zhaoliu age 21 # 注意这里是 -1
1722134018763-1
127.0.0.1:6379>
```

XRANGE key start end [COUNT count]
获取Stream中指定范围内的消息，start和end定义了消息ID的范围，COUNT限制返回结果数量。

```redis
127.0.0.1:6379> xrange k1 - + count 2 
1722134003480-0
name
zhangsan
age
18
1722134011535-0
name
lisi
age
19
127.0.0.1:6379>
```

-:表示Stream中的最小ID
+:表示Stream中的最大ID

XREVRANGE key end start [COUNT count]:类似于XRANGE，但消息按逆序返回

```redis
127.0.0.1:6379> xrevrange k1 + - count 2
1722134018763-1
name
zhaoliu
age
21
1722134018763-0
name
wangwu
age
20
127.0.0.1:6379>
```

XLEN key:返回Stream中消息的数量

```redis
127.0.0.1:6379> xlen k1
4
127.0.0.1:6379>
```

## 删除消息命令(XDEL)

XDEL key ID [ID ...]:删除Stream中指定ID的消息

```redis
127.0.0.1:6379> xlen k1
4
127.0.0.1:6379> xdel k1 1722134003480-0
1
127.0.0.1:6379> xlen k1
3
127.0.0.1:6379>
```

## 截取消息命令(XTRIM)

XTRIM key MAXLEN|MINID [=|~] threshold [LIMIT count]:截取Stream

MAXLEN:表示允许的最大长度,保留大的
MIDID:允许的最小id,这个id之前的消息会被截取掉

```redis
127.0.0.1:6379> xrange k1 - + 
1722134011535-0
name
lisi
age
19
1722134018763-0
name
wangwu
age
20
1722134018763-1
name
zhaoliu
age
21
127.0.0.1:6379> xtrim k1 maxlen 2 
1
127.0.0.1:6379> xrange k1 - + 
1722134018763-0
name
wangwu
age
20
1722134018763-1
name
zhaoliu
age
21
127.0.0.1:6379> xtrim k1 minid 1722134018763-1
1
127.0.0.1:6379> xrange k1 - +
1722134018763-1
name
zhaoliu
age
21
127.0.0.1:6379>
```

## 消费消息命令(XREAD)

XREAD [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] id [id ...]
非阻塞或阻塞式读取一个或多个Stream中的消息

COUNT 指定最多读取的消息数量
BLOCK milliseconds指定阻塞等待新消息的最长时间（毫秒）。默认不阻塞,如果milliseconds设置为0,则永远阻塞

```redis
127.0.0.1:6379> xread count 2 streams k1 0-0
k1
1722134018763-1
name
zhaoliu
age
21
127.0.0.1:6379>
```

0-0代表从最小的ID开始获取Stream中的消息，当不指定count，将会返回Stream中的所有消息，
注意也可以使用 0(00/000也都是可以的)

阻塞读取示例:
客户端1:

```redis
127.0.0.1:6379> xread count 1 block 0 streams k1 $
k1
1722135859526-0
name
ddd
127.0.0.1:6379>
```

客户端2:

```redis
127.0.0.1:6379> xadd k1 * name ddd
"1722135859526-0"
127.0.0.1:6379>
```

$:表示只消费新的消息，比当前id还要大的id, 只有当客户端2添加数据之后,客户端1才会进行消费

## 消费者组管理命令

消息创建好之后,就需要消费者来进行消费.而创建消费者要分组进行创建.

创建消费者组(XGROUP)
XGROUP create key groupname id|$ [MKSTREAM] [ENTRIESREAD entries_read]

```redis
127.0.0.1:6379> XGROUP create k1 group1 $
OK
127.0.0.1:6379> XGROUP create k1 group2 0
OK
127.0.0.1:6379>
```

$表示消费新来的数据
0表示从Stream头部开始消费

在消费者组中读取消息(XREADGROUP)
XREADGROUP GROUP group consumer [COUNT count] [BLOCK milliseconds] [NOACK] STREAMS key [key ...] id [id ...]
在消费者组中读取消息，支持消息确认机制（通过ACK或NOACK）

```redis
127.0.0.1:6379> xreadgroup group group2 consumer1 streams k1 >
k1
1722134018763-1
name
zhaoliu
age
21
1722135859526-0
name
ddd
127.0.0.1:6379> xreadgroup group group2 consumer2 streams k1 >

127.0.0.1:6379>
```

> 用于XREADGROUP命令，表示迄今还没有发送给组中使用者的信息，会更新消费者组的最后ID

消费者不存在时,会自动创建，group2这个组中的consumer1读完了k1中的所有消息,
那么当group2中consumer2再来读时,就读不到任何消息了

但当我新建一个group3组,在group3组中读数据时,又能够读到消息

```redis
127.0.0.1:6379> XGROUP create k1 group3 0
OK
127.0.0.1:6379> xreadgroup group group3 consumer3 streams k1 > 
k1
1722134018763-1
name
zhaoliu
age
21
1722135859526-0
name
ddd
127.0.0.1:6379>
```

为了防止一个消费者读完所有消息,我们可以使用count参数来限制消费者读取几条消息,以此实现负载均衡

```redis
127.0.0.1:6379> xadd k1 * name zhangsan
1722136817088-0
127.0.0.1:6379> xadd k1 * name lisi
1722136826357-0
127.0.0.1:6379> XGROUP create k1 group4 0
OK
# 限制只能读2条消息
127.0.0.1:6379> xreadgroup group group4 consumer1 count 2 streams k1 >
k1
1722134018763-1
name
zhaoliu
age
21
1722135859526-0
name
ddd
127.0.0.1:6379> xreadgroup group group4 consumer2 count 2 streams k1 >
k1
1722136817088-0
name
zhangsan
1722136826357-0
name
lisi
127.0.0.1:6379>
```

查询消费者组的信息(XPENDING)
XPENDING key group [[IDLE min-idle-time] start end count [consumer]]
查询消费者组的信息，包括待处理消息的数量等。

```redis
127.0.0.1:6379> xpending k1 group2 # group2组中消费读取情况
2
1722134018763-1
1722135859526-0
consumer1
2
127.0.0.1:6379> xpending k1 group4 # group4组中消费读取情况
4
1722134018763-1
1722136826357-0
consumer1
2
consumer2
2
127.0.0.1:6379>
```

确认消费者成功处理了消息(XACK)
XACK key group id [id ...]:用于确认（acknowledge）消费者已经成功处理了消息。

```redis
127.0.0.1:6379> xpending k1 group4 - + 5 consumer1
1722134018763-1
consumer1
174359
1
1722135859526-0
consumer1
174359
1
127.0.0.1:6379> xack k1 group4 1722134018763-1
1
127.0.0.1:6379> xpending k1 group4 - + 5 consumer1
1722135859526-0
consumer1
210301
1
127.0.0.1:6379>
```

当客户端从Stream中读取消息后，会使用XACK命令告知Redis，它所指定的消息ID对应的消息
已经得到了妥善处理。Redis接收到XACK命令后，会将这些消息ID从消费者组的待确认（pending）
消息列表中移除，这样，这些消息就不会再次被同一个消费者组内的消费者获取到，实现了消息的确认与去重处理。
