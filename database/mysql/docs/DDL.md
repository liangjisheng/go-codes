# ddl

Online DDL语法上，其实并没有特殊的之处的。我们知道Online DDL这个概念在MySQL5.6首次出现，虽然在语法上没有特殊的之处，
但官方为我们在原来的DDL语法上添加了一些控制性能与并发的属性

```sql
alter table t_base_user modify telephone varchar(50),lock=none ;
```

其中 lock=none 是我们比较陌生的，这个就是MySQL5.6中用来控制性能的属性，需要注意的是，在MySQL5.6之前的版本中，
这样的语法是不支持的。执行时会直接报语法不支持的错误。那 lock=none 代表什么意思呢

官方给我们提供了几个可选项

LOCK=EXCLUSIVE : 表示独占锁，DDL语句执行期间会阻塞该表的所有请求。
LOCK=SHARED：共享锁，DDL语句执行期间会阻塞除查询外的所有DML操作，如: insert，update等。
LOCK=NONE: 允许所有查询以及DML操作。
LOCK=DEFAULT 默认级别，MySQL尽可能允许最大的并发操作。
当我们不显示指定时，默认就为LOCK=DEFAULT类型

在执行有些DDL语句时(修改字段类型)，其实是有风险的，那怎么避免风险呢？这个时候就要通过执rows affected类分析了

更改列的默认值 (超快, 完全不影响表数据)
Query OK, 0 rows affected (0.07 sec)

添加一个索引（需要时间，但是0 rows affected显示表格不被复制）
Query OK, 0 rows affected (21.42 sec)

更改列的数据类型（需要大量的时间，并且需要重建表的所有行）
Query OK, 1671168 rows affected (1 min 35.54 sec)

我们可以通过上述的执行结果，查看该语句是否复制记录，重建整个表格(如果重建，成本会非常高，而且还会影响线上DML)等。
显然，直接在生产上运行该语句后，再看结果来分析。已经没有意义了，我们可以通过在测试环境中来看。具体步骤如下:

1. 复制需要ddl的生产表结构到测试环境中。
2. 添加部分数据到该表中。
3. 执行ddl操作。
4. 查看rows affected 值是否为0，非0时，意味着操作需要重建整个表，这时候就需要重新指定方案，如：在停机时进行操作，或业务低估时进行操作。