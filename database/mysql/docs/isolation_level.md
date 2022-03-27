# isolation level

[mysql isolation level](https://juejin.cn/post/6844903805822173198)

mysql 4 种事务隔离级别, 在SQL中定义了四种隔离的级别，每一种隔离级别都规定了一个事务中的修改，
哪些是在事务内和事务间是可见的，哪些是不可见的。较低级别的隔离通常来说能承受更高的并发，系统的开销也会更小

READ UNCOMMITTED（未提交读）
在READ UNCOMMITTED级别，事务的修改，即使没有提交，对其他事务也都是可见的。事务可以读取未提交的数据，
这也被称为脏读（Dirty Read）。这个级别的隔离会导致很多问题，虽然在性能方面是最优的，但是缺乏其他级别的很多好处，
所以这种隔离的级别很少在实际中应用

设置隔离级别

```sql
set session transaction isolation level read uncommitted;
```

READ COMMITTED (读已提交)
大多数数据库系统默认的隔离级别都是READ COMMITTED(但MySQL不是)，"读已提交"简单的定义:一个事务只能看见已经提交的事务的修改结果
换句话说，一个事务从开启事务到提交事务之前，对其他事务都是不可见的，因此在同一个事务中的两次相同查询结果可能不一样
故这种隔离级别有时候也叫不可重复读（NONREPEATABLE READ）

REPEATABLE READ (可重复读)
"可重复读"是MySQL的默认事务隔离级别。REPEATABLE READ解决了脏读的问题，该级别保证了在同一次事务中多次查询相同的语句结果是一致的
但是"可重复读"隔离级别无法避免产生幻行（Phantom Row）的问题，MySQL的InnoDB引擎通过多版本并发控制
（MVCC，Multiversion Concurrency Controller）解决了幻读的问题

SERIALIZABLE是最高的隔离级别，它通常通过强制事务串行，避免了前面说的幻读问题。简单来说，
"可串行化"会在读取的每一行数据上都加锁，所以可能会导致大量的锁等待和超时问题，所以在实际的生产环境中也很少会用到这个隔离级别，
只有在非常需要确保数据的一致性切可以接受没有并发的情况下，才会考虑使用这个隔离级别
