# bitfield

Bitfield(位域)命令可以将一个 Redis 字符串看作是一个由二进制位组成的数组， 并对这个数组中任意偏移进行访问

BITFIELD_RO: Read-only variant of the BITFIELD command. It is like the original BITFIELD
but only accepts GET subcommand and can safely be used in read-only replicas.

BITFIELD key [GET type offset]: 返回指定的位域

key: 要操作的Redis键。
GET: 表示要从字符串值中读取位。
type: 指定读取数据的类型，可以是u（无符号整数）、i（有符号整数）
offset: 位字段的起始偏移位置，从0开始计数

```redis
127.0.0.1:6379> set k1 abcd
OK
127.0.0.1:6379> get k1
"abcd"
127.0.0.1:6379> bitfield k1 get i8 0
1) (integer) 97
127.0.0.1:6379> bitfield k1 get i8 8
1) (integer) 98
127.0.0.1:6379>
```

BITFIELD key [SET type offset value]:设置指定位域的值并返回它的原值

key: 要操作的Redis键。
SET: 表示要设置字符串值中的位。
type: 类型标识，可以是u（无符号整数）、i（有符号整数）。同GET操作，目前Redis不直接支持浮点数的位操作。
offset: 要设置位字段的起始偏移位置。
value: 要设置的值，根据类型不同，这个值有不同的解释。对于整数，它是你希望写入的整数值

```redis
127.0.0.1:6379> bitfield k1 set i8 0 111
1) (integer) 97
127.0.0.1:6379> get k1
"obcd"
127.0.0.1:6379>
```

BITFIELD key [INCRBY type offset increment]

key: 要操作的Redis键。
INCRBY: 表示自增。
type: 类型标识，可以是u（无符号整数）、i（有符号整数）。同GET操作，目前Redis不直接支持浮点数的位操作。
offset: 要设置位字段的起始偏移位置。
increment: 自增的数值。

```redis
127.0.0.1:6379> get k1
"obcd"
127.0.0.1:6379> bitfield k1 incrby i8 0 1
1) (integer) 112
127.0.0.1:6379> get k1
"pbcd"
127.0.0.1:6379>
```

而Redis提供了三种移除控制方式:

WRAP:使用回绕(wrap around)方法处理有符号整数和无符号整数的溢出情况
SAT: 使用饱和计算(saturation arithmetic)方法处理溢出下溢计算的结果为最小的整数值，而上溢计算的结果为最大的整数值
FAIL: 命令将拒绝执行那些会导致上溢或者下溢情况出现的计算,并向用户返回空值表示计算未被执行

WRAP方式示例

```redis
127.0.0.1:6379> bitfield k1 set i8 0 127
1) (integer) 112
127.0.0.1:6379> bitfield k1 get i8 0
1) (integer) 127
127.0.0.1:6379> bitfield k1 incrby i8 0 1
1) (integer) -128
127.0.0.1:6379> bitfield k1 get i8 0
1) (integer) -128
127.0.0.1:6379>
```

SAT方式示例

```redis
127.0.0.1:6379> bitfield k1 overflow sat set i8 0 128
1) (integer) -128
127.0.0.1:6379> bitfield k1 get i8 0
1) (integer) 127
127.0.0.1:6379>
```

FAIL方式示例

```redis
127.0.0.1:6379> bitfield k1 overflow fail set i8 0 128
1) (nil)
127.0.0.1:6379> bitfield k1 get i8 0
1) (integer) 127
127.0.0.1:6379>
```
