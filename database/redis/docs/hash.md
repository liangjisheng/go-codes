# hash

Redis中的Hash类型是一种高效的数据结构，用于存储键值对的集合。这种类型特别适用于表示对象，
因为它允许你将对象的多个属性（fields）作为单独的条目存储在同一个键（key）之下

HSET key field value：设置field对应的值。如果field已存在则更新，否则创建。
HGET key field：获取指定field的值。
HGETALL key：获取该key下所有field及其对应的值。
HLEN key：返回该key下field的数量。
HKEYS key：获取该key下所有field。
HVALS key：获取该key下所有field的值。
HEXISTS key field：检查field是否存在。
HDEL key field [field ...]：删除一个或多个field。
HINCRBY key field increment：将哈希表中field字段的值增加指定的整数值increment
HINCRBYFLOAT key field increment: 与HINCRBY类似，但是increment可以是浮点数，
即支持对哈希表中field字段的值进行浮点数的原子性递增。
HSETNX key field value：只有当field不存在时才设置其值。

HSET key field value：设置field对应的值。如果field已存在则更新，否则创建。
HGET key field：获取指定field的值

```redis
127.0.0.1:6379> hset user1 name zhangsan age 18
(integer) 2
127.0.0.1:6379> hget user1 name
"zhangsan"
127.0.0.1:6379> hget user1 age
"18"
127.0.0.1:6379>
```

HGETALL key：获取该key下所有field及其对应的值

```redis
127.0.0.1:6379> hgetall user1
1) "name"
2) "zhangsan"
3) "age"
4) "18"
127.0.0.1:6379>
```

HLEN key：返回该key下field的数量

```redis
127.0.0.1:6379> hlen user1
(integer) 2
127.0.0.1:6379>
```

HKEYS key：获取该key下所有field。
HVALS key：获取该key下所有field的值

```redis
127.0.0.1:6379> hkeys user1
1) "name"
2) "age"
127.0.0.1:6379> hvals user1
1) "zhangsan"
2) "18"
127.0.0.1:6379>
```

HEXISTS key field：检查field是否存在。存在返回1,不存在返回0

```redis
127.0.0.1:6379> hkeys user1
1) "name"
2) "age"
127.0.0.1:6379> hexists user1 name
(integer) 1
127.0.0.1:6379> hexists user1 email
(integer) 0
```

HDEL key field [field ...]：删除一个或多个field

```redis
127.0.0.1:6379> hkeys user1
1) "name"
2) "age"
127.0.0.1:6379> hdel user1 name
(integer) 1
127.0.0.1:6379> hkeys user1
1) "age"
127.0.0.1:6379>
```

HINCRBY key field increment
HINCRBYFLOAT key field increment

```redis
127.0.0.1:6379> hgetall user1
1) "age"
2) "18"
3) "name"
4) "lisi"
5) "score"
6) "55.5"
127.0.0.1:6379> hincrby user1 age 1
(integer) 19
127.0.0.1:6379> hincrby user1 age -2 # 可以是负数
(integer) 17
127.0.0.1:6379> hincrbyfloat user1 score 0.5
"56"
127.0.0.1:6379> hincrbyfloat user1 score -1.5 # 可以是负数
"54.5"
127.0.0.1:6379>
```

HSETNX key field value：只有当field不存在时才设置其值

```redis
127.0.0.1:6379> hgetall user1
1) "age"
2) "17"
3) "name"
4) "lisi"
5) "score"
6) "54.5"
127.0.0.1:6379> hsetnx user1 age 22
(integer) 0
127.0.0.1:6379> hsetnx user1 phone 123123
(integer) 1
127.0.0.1:6379>
```