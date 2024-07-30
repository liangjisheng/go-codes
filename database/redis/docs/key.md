# key

SET key value：设置key对应的值。
GET key：获取key对应的值。
SETNX key value：只有当key不存在时，才设置key的值

```redis
127.0.0.1:6379> set k1 hello
OK
127.0.0.1:6379> get k1
"hello"
127.0.0.1:6379> setnx k1 1 # k1存在,不修改
(integer) 0
127.0.0.1:6379> get k1
"hello"
127.0.0.1:6379> setnx k2 2 # k2不存在,生成key并赋值
(integer) 1
127.0.0.1:6379> get k2
"2"
127.0.0.1:6379>
```

EXPIRE key seconds：为key设置过期时间（单位秒）。
SETEX key seconds value：设置key的值，并为其设置过期时间（单位秒）。
TTL key：获取key的剩余生存时间（单位秒），如果key没有设置过期时间则返回-1，如果key不存在则返回-2。
PTTL key：同TTL，但返回毫秒精度的剩余生存时间。
PEXPIRE key milliseconds：为key设置过期时间（单位毫秒）。
PERSIST key：移除key的过期时间，使其永不过期。

```redis
127.0.0.1:6379> set k1 1 # 没设置过期时间
OK
127.0.0.1:6379> ttl k1 
(integer) -1  
127.0.0.1:6379> expire k1 5 # 设置过期时间 5s
(integer) 1
127.0.0.1:6379> ttl k1
(integer) 3 
127.0.0.1:6379> ttl k1
(integer) 1
127.0.0.1:6379> ttl k1
(integer) -2 # 已过期
     
127.0.0.1:6379> setex k2 5 2 # 设置key并设置过期时间
OK
127.0.0.1:6379> ttl k2
(integer) 4
127.0.0.1:6379> ttl k2
(integer) 1
127.0.0.1:6379> ttl k2
(integer) -2
127.0.0.1:6379> 
     
127.0.0.1:6379> setex k3 100 3 
OK
127.0.0.1:6379> ttl k3
(integer) 96
127.0.0.1:6379> persist k3 # 移除k3的过期时间
(integer) 1
127.0.0.1:6379> ttl k3
(integer) -1
127.0.0.1:6379>
```

DEL key [key ...]：删除一个或多个key

```redis
127.0.0.1:6379> set k1 1
OK
127.0.0.1:6379> set k2 2
OK
127.0.0.1:6379> set k3 3
OK
127.0.0.1:6379> del k1
(integer) 1
127.0.0.1:6379> del k2 k3
(integer) 2
127.0.0.1:6379>
```

EXISTS key：检查key是否存在，存在返回1，否则返回0。
TYPE key：返回key所储存的值的类型

```redis
127.0.0.1:6379> set k1 1
OK
127.0.0.1:6379> exists k1
(integer) 1
127.0.0.1:6379> exists k2
(integer) 0
127.0.0.1:6379> type k1
string
127.0.0.1:6379>
```

DBSIZE：返回当前数据库中key的数量

```redis
127.0.0.1:6379> keys *
1) "k1"
127.0.0.1:6379> dbsize 
(integer) 1
127.0.0.1:6379>
```

KEYS pattern：查找所有符合给定模式的key，但此命令可能阻塞服务器，慎用

```redis
127.0.0.1:6379> set k1 1
OK
127.0.0.1:6379> set k2 2
OK
127.0.0.1:6379> keys * #key * 这个操作十分危险,慎用!
1) "k2"
2) "k1"
127.0.0.1:6379>
```

MOVE key dbindex：将key移动到另一个数据库dbindex中。
select dbindex: 切换数据库

```redis
# Redis中默认是有16个数据库,数据库下标从0开始
# 数据库数量 可在Redis的配置文件中修改
127.0.0.1:6379> set k1 1
OK
127.0.0.1:6379> set k2 2
OK
127.0.0.1:6379> move k1 6 #将k1从0号数据库移到6号数据库
(integer) 1
127.0.0.1:6379> keys *
1) "k2" 
127.0.0.1:6379> select 6 # 进入6号数据库
OK
127.0.0.1:6379[6]> keys *
1) "k1" 
127.0.0.1:6379[6]> select 15 # 第16个数据库
OK
127.0.0.1:6379[15]> select 16 
(error) ERR DB index is out of range 
127.0.0.1:6379[15]>
```

RANDOMKEY：随机返回数据库中的一个key。
flushdb: 清空当前数据库
flushall: 清空全部数据库

```redis
127.0.0.1:6379> keys *
1) "k4"
2) "k2"
3) "k3"
127.0.0.1:6379> randomkey
"k4"
127.0.0.1:6379> randomkey
"k4"
127.0.0.1:6379> randomkey
"k3"
127.0.0.1:6379> randomkey
"k2"
    
127.0.0.1:6379> flushdb # 清空当前数据库
OK
127.0.0.1:6379> keys *
(empty array)
127.0.0.1:6379> 127.0.0.1:6379> select 6 # 查看6号数据库
OK
127.0.0.1:6379[6]> keys *
1) "k1" # 6号库key没有被清除 

127.0.0.1:6379[6]> select 0 # 进入0号数据库
OK
127.0.0.1:6379> flushall # 清空全部数据库
OK
127.0.0.1:6379> select 6 
OK
127.0.0.1:6379[6]> keys * 
(empty array) # 6号库key已被清除
```
