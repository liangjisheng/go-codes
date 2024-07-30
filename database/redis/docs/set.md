# set

Redis的Set类型是一个无序的不重复字符串集合。

SADD key member [member ...]：向集合添加一个或多个成员，如果成员已存在则不执行任何操作。  
SMEMBERS key：返回集合中的所有成员。  
SISMEMBER key member：判断成员是否在集合内。

```redis
127.0.0.1:6379> sadd set1 1 1 2 2 3 3 
(integer) 3 # 去重
127.0.0.1:6379> smembers set1
1) "1"
2) "2"
3) "3"
127.0.0.1:6379> sismember set1 3
(integer) 1
127.0.0.1:6379> sismember set1 4
(integer) 0
127.0.0.1:6379>
```

SREM key member [member ...]：从集合中移除一个或多个成员。  
SCARD key：返回集合中元素的数量。  

```redis
127.0.0.1:6379> srem set1 1 # 删除1
(integer) 1
127.0.0.1:6379> scard set1 
(integer) 2 # 还剩2个元素
127.0.0.1:6379> srem set1 2 3 # 删除 2 3
(integer) 2
127.0.0.1:6379> scard set1
(integer) 0  # 还剩0元素
127.0.0.1:6379>
```

SRANDMEMBER key [count]：随机返回集合中的一个或多个成员，但不移除。

```redis
127.0.0.1:6379> sadd set1 1 2 3 4 5
(integer) 5
127.0.0.1:6379> srandmember set1 1 # 随机返回一个元素
1) "2"
127.0.0.1:6379> srandmember set1 3 # 随机返回三个元素
1) "3"
2) "4"
3) "1"
127.0.0.1:6379> scard set1
(integer) 5  # set集合不变
127.0.0.1:6379>
```

SPOP key [count]：随机移除并返回集合中的一个或多个成员。  

```redis
127.0.0.1:6379> spop set1 1 # 随机返回并删除一个元素
1) "1"
127.0.0.1:6379> smembers set1 
1) "2"
2) "3"
3) "4"
4) "5"
127.0.0.1:6379> spop set1 3 # 随机返回并删除三个元素
1) "3"
2) "2"
3) "4"
127.0.0.1:6379> smembers set1
1) "5"
```

smove source destination member:将一个集合中的成员移动到另一个集合中

```redis
127.0.0.1:6379> sadd set1 1 2 3 4 5
(integer) 4
127.0.0.1:6379> sadd set2 a b c
(integer) 3
127.0.0.1:6379> smove set1 set2 1
(integer) 1
127.0.0.1:6379> smembers set2
1) "a"
2) "c"
3) "b"
4) "1"
127.0.0.1:6379>
```

SINTER key [key ...]：返回给定集合的交集。  

```redis
127.0.0.1:6379> sadd set1 1 2 3 a b c
(integer) 6
127.0.0.1:6379> sadd set2 2 3 4 c d e
(integer) 6
127.0.0.1:6379> sinter set1 set2
1) "3"
2) "c"
3) "2"
127.0.0.1:6379>
```

SUNION key [key ...]：返回给定集合的并集。  

```redis
127.0.0.1:6379> sadd set1 1 2 3 a b c
(integer) 6
127.0.0.1:6379> sadd set2 2 3 4 c d e
(integer) 6
127.0.0.1:6379> sunion set1 set2
1) "e"
2) "a"
3) "c"
4) "b"
5) "3"
6) "2"
7) "4"
8) "1"
9) "d"
127.0.0.1:6379>
```

SDIFF key [key ...]：返回给定集合的差集,即存在于第一个集合但不存在于其他集合的成员。  

```redis
127.0.0.1:6379> sadd set1 1 2 3 a b c
(integer) 6
127.0.0.1:6379> sadd set2 2 3 4 c d e
(integer) 6
127.0.0.1:6379> sdiff set1 set2
1) "a"
2) "b"
3) "1"
127.0.0.1:6379> sdiff set2 set1
1) "e"
2) "d"
3) "4"
127.0.0.1:6379>
```

Redis7新命令sintercard
sintercard numkeys key [key ...] [LIMIT limit]:它不返回结果集，而只返回结果的基数(个数)。
返回由所有给定集合的交集产生的集合的基数

```redis
127.0.0.1:6379> sinter set1 set2 # 返回结果集
1) "c"
2) "3"
3) "2"
127.0.0.1:6379> sintercard 2 set1 set2 # 返回结果集的个数
(integer) 3
127.0.0.1:6379> sintercard 2 set1 set2 limit 2 # 限制返回结果的个数
(integer) 2
127.0.0.1:6379>
```
