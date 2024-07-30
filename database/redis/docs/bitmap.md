# bitmap

Redis的Bitmap类型并不是一个独立的数据类型，而是对String类型的一种巧妙使用，
允许开发者以位级别操作字符串中的数据。Bitmap可以视为一个巨大的位数组，每个位
（bit）可以存储0或1的值，这使得Bitmap非常适合用于存储和操作大量的布尔值或者
进行高效的统计计数。

SETBIT key offset value：设置位图中指定偏移量的值。  
GETBIT key offset：获取位图中指定偏移量的值。  
offset从0开始

```redis
127.0.0.1:6379> setbit k1 1 1
(integer) 0 # 返回值是设置前的值
127.0.0.1:6379> setbit k1 1 0
(integer) 1
127.0.0.1:6379> setbit k1 2 1 
(integer) 0
127.0.0.1:6379> setbit k1 3 3 # 值只能是0和1
(error) ERR bit is not an integer or out of range
127.0.0.1:6379> getbit k1 0
(integer) 0
127.0.0.1:6379> getbit k1 1
(integer) 0
127.0.0.1:6379> getbit k1 2
(integer) 1
127.0.0.1:6379>
```

strlen key:获取该Bitmap所占用的字节数，而不是比特位中1的个数。 

```redis
127.0.0.1:6379> setbit k1 7 1
(integer) 1
127.0.0.1:6379> strlen k1
(integer) 1
127.0.0.1:6379> setbit k1 8 0
(integer) 0
127.0.0.1:6379> strlen k1
(integer) 2
127.0.0.1:6379>
```

BITCOUNT key [start end]：计算位图中位值为1的个数，可选地限制在指定范围内。  

```redis
127.0.0.1:6379> setbit k1 0 1
(integer) 0
127.0.0.1:6379> setbit k1 1 1
(integer) 0
127.0.0.1:6379> setbit k1 7 1
(integer) 0
127.0.0.1:6379> bitcount k1 
(integer) 3
127.0.0.1:6379> bitcount k1 0 3
(integer) 3
127.0.0.1:6379> bitcount k1 0 3 byte # 以byte为单位
(integer) 3
127.0.0.1:6379> bitcount k1 0 3 bit # 以bit为单位
(integer) 2
127.0.0.1:6379>
```

BITOP operation destkey key [key ...]：对多个位图执行AND、OR、NOT、XOR操作，并将结果保存到destkey。
假设我们有两个Bitmap键use1和user2，分别代表了两天内用户的在线状态，其中1表示在线，0表示离线

他们三天中的在线状态如下

```redis
127.0.0.1:6379> setbit user1 0 1
(integer) 0
127.0.0.1:6379> setbit user1 1 1
(integer) 0
127.0.0.1:6379> setbit user2 0 1
(integer) 0
127.0.0.1:6379> setbit user2 2 1
(integer) 0
127.0.0.1:6379>
```

如果想要找出这两天都在线的用户，可以使用AND操作。

```redis
127.0.0.1:6379> bitop and k1 user1 user2
(integer) 1
127.0.0.1:6379> getbit k1 0
(integer) 1
127.0.0.1:6379> getbit k1 1
(integer) 0
127.0.0.1:6379>
```

如果我们想找出至少有一天在线的用户，可以使用OR操作

```redis
127.0.0.1:6379> bitop or k2 user1 user2
(integer) 1
127.0.0.1:6379> bitcount k2
(integer) 3
127.0.0.1:6379>
```

如果我们要找出只在某一天在线，而不在另一天在线的用户，可以使用XOR操作

```redis
127.0.0.1:6379> getbit k3 0
(integer) 0
127.0.0.1:6379> getbit k3 1
(integer) 1
127.0.0.1:6379> getbit k3 2
(integer) 1
127.0.0.1:6379>
```

另外，BITOP命令不直接支持NOT操作，因为NOT操作需要一个源位图和目标位图。但是，可以通过创建一个全1的Bitmap
（假设长度与原Bitmap相同），然后使用XOR操作达到NOT的效果。

BITFIELD key [GET type offset] [SET type offset value] [...]：更复杂的位操作，可以一次执行多个位操作。
