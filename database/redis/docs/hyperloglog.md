# hyper loglog

Redis的HyperLogLog是一种高级数据结构，专门用于基数估算。它主要用于在极低的空间成本
下计算一个数据集中不重复元素的数量（即基数）。HyperLogLog是基于概率的数据结构，能够
以极高的效率和极小的内存占用（固定12KB）来近似计算2^64个不同元素的基数，特别适合处理
大规模数据集的统计计数，如网站独立访客（UV）计数等场景。

PFADD key element [element ...]：向HyperLogLog添加元素。
PFCOUNT key [key ...]：估算单个或多个HyperLogLog的基数。
PFMERGE destkey sourcekey [sourcekey ...]：将多个HyperLogLog合并到一个destkey中。

```redis
127.0.0.1:6379> pfadd k1 1 2 3 4 5 6
(integer) 1
127.0.0.1:6379> pfcount k1
(integer) 6
127.0.0.1:6379> pfadd k2 3 3 4 5 5 6 7
(integer) 1
127.0.0.1:6379> pfcount k2
(integer) 5
127.0.0.1:6379> pfmerge k3 k1 k2
OK
127.0.0.1:6379> pfcount k3
(integer) 7
127.0.0.1:6379>
```

但是，因为 HyperLogLog 只会根据输入元素来计算基数，而不会储存输入元素本身，所以
HyperLogLog 不能像集合那样，返回输入的各个元素。

```redis
127.0.0.1:6379> pfadd key1 1 1 2 2 3 3
(integer) 1
127.0.0.1:6379> type key1
string
127.0.0.1:6379> get key1
"HYLL\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x80]f\x80Mt\x80Q,\x8cC\xf3"
127.0.0.1:6379>
```
