# string

Redis的String类型是其最基本的数据类型，它允许以键值对（key-value）的形式存储数据，其中key
始终是字符串类型，而value则可以是任何二进制安全的数据，这意味着它可以存储任何形式的字符串数据，
包括文本、图片、序列化对象等，最大可达512MB

NX– 仅键不存在的时候设置键值。
XX– 仅键存在的时候设置键值。

```redis
127.0.0.1:6379> set k1 1 nx 
OK
127.0.0.1:6379> set k1 2 nx
(nil)
127.0.0.1:6379> get k1
"1"
127.0.0.1:6379> 
127.0.0.1:6379> set k1 2 xx
OK
127.0.0.1:6379> set k2 2 xx
(nil)
127.0.0.1:6379>
```

GET– 返回存储在 key 中的旧字符串，如果 key 不存在，则返回 nil。如果存储在 key 的值不是字符串，
则返回并中止错误

```redis
127.0.0.1:6379> set k1 v1
OK
127.0.0.1:6379> set k1 v2 get # 
"v1"
127.0.0.1:6379> get k1
"v2"
127.0.0.1:6379>
```

EX seconds – 设置指定的到期时间，以秒为单位（正整数）。

PX 毫秒 （milliseconds） – 设置指定的过期时间，以毫秒为单位（正整数）。

EXAT timestamp-seconds – 设置密钥过期的指定 Unix 时间，以秒为单位（正整数）。

PXAT timestamp-milliseconds – 设置密钥过期的指定 Unix 时间，以毫秒为单位（正整数）。

```redis
# EX选项
127.0.0.1:6379> set k2 v2 ex 5 # 设置k2过期时间为5s
OK
127.0.0.1:6379> ttl k2
(integer) 3
127.0.0.1:6379> ttl k2
(integer) -2
# PX选项
127.0.0.1:6379> set k2 v2 px 5000 # 设置k2过期时间为5s
OK
127.0.0.1:6379> ttl k2
(integer) 3
127.0.0.1:6379> ttl k2
(integer) -2
127.0.0.1:6379>
```

KEEPTTL– 保留设置前指定键的生存时间

```redis
# set命令导致key的过期时间失效
127.0.0.1:6379> set k1 1 ex 30
OK
127.0.0.1:6379> ttl k1
(integer) 27
127.0.0.1:6379> set k1 v1 
OK
127.0.0.1:6379> ttl k1
(integer) -1
127.0.0.1:6379> 
# 使用keepttl会保留过期时间
127.0.0.1:6379> set k2 2 ex 30
OK
127.0.0.1:6379> ttl k2
(integer) 28
127.0.0.1:6379> set k2 v2 keepttl
OK
127.0.0.1:6379> ttl k2
(integer) 19
127.0.0.1:6379>
```

MSET命令是Redis中用于同时设置多个键值对的命令。它允许你一次执行多个SET操作

Redis还提供了一个MSETNX命令，它的工作方式与MSET类似，但只有在所有给定的键都不存在时才会执行设置操作，
即“Set Multiple if Not eXists”。

MGET命令用于一次性获取多个键对应的值

```redis
127.0.0.1:6379> mset k1 1 k2 2 k3 3
OK
127.0.0.1:6379> mget k1 k2 k3
1) "1"
2) "2"
3) "3"
127.0.0.1:6379> msetnx k1 1 k4 4
(integer) 0
127.0.0.1:6379> keys *
1) "k2"
2) "k1"
3) "k3"
127.0.0.1:6379>
```

GETRANGE命令用于获取存储在键中的字符串值的子字符串。

SETRANGE命令用于覆写或插入字符串键中存在的值的一部分。

```redis
127.0.0.1:6379> set k1 axbhcask
OK
127.0.0.1:6379> getrange k1 0 -1
"axbhcask"
127.0.0.1:6379> getrange k1 2 5
"bhca"
127.0.0.1:6379> setrange k1 2 ZZZZ
(integer) 8
127.0.0.1:6379> getrange k1 0 -1
"axZZZZsk"
127.0.0.1:6379> setrange k1 15  ZZZZ
(integer) 19
127.0.0.1:6379> getrange k1 0 -1
"axZZZZsk\x00\x00\x00\x00\x00\x00\x00ZZZZ"
127.0.0.1:6379>
```

INCR和INCRBY都是Redis中用于对字符串值进行原子性递增操作的命令，主要区别在于增量的指定方式：

INCR (increment):
增量值固定为1。当你执行INCR key时，Redis会将键key所储存的数字值增加1。如果key不存在，Redis
会先将其初始化为0，然后再执行递增操作。这个命令常用于实现计数器，如访问次数统计。

INCRBY (increment by):
允许你指定增量值。使用格式为INCRBY key increment，其中increment是你想要增加的具体数值，
可以是正数也可以是负数。如果key不存在，同样会被初始化为0后再执行递增或递减操作。这提供了更
灵活的数值调整方式，适用于需要非1步长递增或递减的场景，例如累加交易额、分数调整等。

```redis
127.0.0.1:6379> set k1 10
OK
127.0.0.1:6379> incr k1
(integer) 11
127.0.0.1:6379> incr k1
(integer) 12
127.0.0.1:6379> incrby k1 5
(integer) 17
127.0.0.1:6379> incrby k1 5
(integer) 22
127.0.0.1:6379>
```

DECR和DECRBY命令与INCR和INCRBY相似，但它们执行的是递减操作，具体区别如下：

DECR (decrement):
此命令将键key所储存的数值减少1。如果key不存在，Redis会先将其初始化为0，然后再执行递减操作。
类似于INCR，它适用于需要递减计数的场景，例如库存减少、票数递减等。

DECRBY (decrement by):
允许你指定一个减量值。命令格式为DECRBY key decrement，其中decrement是一个整数，表示要减少
的数量，可以是正数也可以是负数（虽然通常用于正值以减少计数）。如果key不存在，其值同样会被初始化
为0，然后执行减量操作。

```redis
127.0.0.1:6379> get k1
"22"
127.0.0.1:6379> decr k1 
(integer) 21
127.0.0.1:6379> decr k1 
(integer) 20
127.0.0.1:6379> decrby k1 5
(integer) 15
127.0.0.1:6379> decrby k1 5
(integer) 10
127.0.0.1:6379>
```

STRLEN和APPEND分别用于获取字符串的长度和在已有字符串后面追加内容

```redis
127.0.0.1:6379> strlen k1
(integer) 3
127.0.0.1:6379> append k1 qwe
(integer) 6
127.0.0.1:6379> get k1
"asdqwe"
127.0.0.1:6379>
```
