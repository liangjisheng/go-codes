# zset

Redis的Zset（Sorted Set，有序集合）是一种特殊的数据结构，它结合了集合（Set）和排序功能。
在Zset中，每个成员（member）都关联一个唯一的分数（score），这个分数用于对集合中的成员进行排序。
因此，Zset中的元素不仅像Set那样不允许重复，还能够根据score的值进行排序。

特性:

有序性：Zset中的元素可以根据score的值进行升序或降序排列。  
唯一性：每个成员在Zset中都是唯一的，就像Set一样。  
分数（score）：可以是任意浮点数，用于排序。相同的score值的成员会按照成员自身的字典顺序排序。  
时间复杂度：对于添加、删除和查找操作，平均时间复杂度通常为O(1)，具体取决于跳跃列表（skiplist）的实现细节。  
跳跃列表（Skiplist）：Redis使用跳跃列表作为Zset的底层实现，这是一种可以在对数时间内完成查找、插入和删除操作的数据结构。  
范围操作：支持快速地执行范围查询，如获取某个分数区间内的成员。  

ZADD key score member [score member ...]：向Zset中添加一个或多个成员及其分数。  
ZRANGE key start stop [WITHSCORES]：返回Zset中指定范围内的成员，可选地包括它们的分数。  
ZREVRANGE key start stop [WITHSCORES]：同ZRANGE，但返回的是按score降序排列的结果。

```redis
127.0.0.1:6379> zadd zset1 10 v1 22.2 v2 30 v3 # 可以是浮点数
(integer) 3
127.0.0.1:6379> zrange zset1 0 -1
1) "v1"
2) "v2"
3) "v3"# 从小到大
127.0.0.1:6379> zrange zset1 0 -1 withscores # 从小到大
1) "v1"
2) "10"
3) "v2"
4) "22.199999999999999"
5) "v3"
6) "30"
127.0.0.1:6379> zrevrange zset1 0 -1 withscores # 从大到小
1) "v3"
2) "30"
3) "v2"
4) "22.199999999999999"
5) "v1"
6) "10"
127.0.0.1:6379>
```

ZSCORE key member:用于获取有序集合中某个成员的分数（score）值
ZREM key member [member ...]：移除Zset中的一个或多个成员。  
ZCARD key：返回Zset中成员的数量。  

```redis
127.0.0.1:6379> zcard zset1
(integer) 3
127.0.0.1:6379> zscore zset1 v1 
"10"
127.0.0.1:6379> zrem zset1 v1
(integer) 1
127.0.0.1:6379> zcard zset1
(integer) 2
127.0.0.1:6379> zrem zset1 v2 v3
(integer) 2
127.0.0.1:6379> zcard zset1
(integer) 0
127.0.0.1:6379>
```

ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]：根据分数范围返回成员。  
ZCOUNT key min max：返回Zset中score值在给定范围内的成员数量。

```redis
127.0.0.1:6379> zadd zset1 10 v1 20 v2 30 v3 40 v4 50 v5
(integer) 5
127.0.0.1:6379> zrangebyscore zset1 20 40 # 不带分数
1) "v2"
2) "v3"
3) "v4"
127.0.0.1:6379> zrangebyscore zset1 20 40 withscores # 带分数
1) "v2"
2) "20"
3) "v3"
4) "30"
5) "v4"
6) "40"
127.0.0.1:6379> zrangebyscore zset1 (20 40 withscores # ( 的意思是不包含,可以理解为开区间,默认是闭区间
1) "v3"
2) "30"
3) "v4"
4) "40"
127.0.0.1:6379> zrangebyscore zset1 20 (40 withscores
1) "v2"
2) "20"
3) "v3"
4) "30"
127.0.0.1:6379> zrangebyscore zset1 20 40 withscores limit 0 2 # 从第几个结果返回几个数据
1) "v2"
2) "20"
3) "v3"
4) "30"
127.0.0.1:6379> zcount zset1 10 50 
(integer) 5
127.0.0.1:6379> zcount zset1 (10 50 
(integer) 4
127.0.0.1:6379>
```

ZRANK key member:获取下标值  
ZREVRANK key member:逆序获取下标值  

```redis
127.0.0.1:6379> zadd zset1 10 v1 20 v2 30 v3 40 v4 50 v5
(integer) 5
127.0.0.1:6379> zrank zset1 v3
(integer) 2
127.0.0.1:6379> zrank zset1 v1
(integer) 0
127.0.0.1:6379> zrevrank zset1 v1
(integer) 4
127.0.0.1:6379>
```

ZMPOP numkeys key [key ...] MIN|MAX [COUNT count]:从键名列表中的第一个非空排序集中弹出
一个或多个元素，它们是成员分数对

```redis
127.0.0.1:6379> ZMPOP 1 zset1 min count 1 # 从1个zset列表中弹出一个最小的
1) "zset1"
2) 1) 1) "v1"
      2) "10"
127.0.0.1:6379> ZMPOP 1 zset1 max count 2 # 从1个zset列表中弹出两个最大的
1) "zset1"
2) 1) 1) "v5"
      2) "50"
   2) 1) "v4"
      2) "40"
127.0.0.1:6379>
```
