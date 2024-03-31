# lock

在 MySQL 中，存储引擎就支持不同类型锁。如： MyISAM 只支持表锁。InnoDB 支持：行锁，表锁，Gap 锁等等
在 InnoDB 中，行锁分为共享锁(Share) 和 排他 (Exclusive) 锁

例1: 事务1持有S 锁, 事务2 请求持有S 锁

事务1

```sql
begin;

mysql> select * from t_base_info where oid = 1 lock in share mode;
+-----+----------+---------------------+---------------------+
| oid | name     | create_time         | updated_time        |
+-----+----------+---------------------+---------------------+
|   1 | andyqian | 2020-03-21 14:34:08 | 2020-03-21 14:34:08 |
+-----+----------+---------------------+---------------------+
1 row in set (0.00 sec)
```

事务2

```sql
begin;

mysql> select * from t_base_info where oid = 1 lock in share mode;
+-----+----------+---------------------+---------------------+
| oid | name     | create_time         | updated_time        |
+-----+----------+---------------------+---------------------+
|   1 | andyqian | 2020-03-21 14:34:08 | 2020-03-21 14:34:08 |
+-----+----------+---------------------+---------------------+
1 row in set (0.00 sec)
```

例子: 事务1持有S 锁, 事务2 请求持有X 锁

事务1

```sql
begin;

mysql> select * from t_base_info where oid = 1 lock in share mode;
+-----+----------+---------------------+---------------------+
| oid | name     | create_time         | updated_time        |
+-----+----------+---------------------+---------------------+
|   1 | andyqian | 2020-03-21 14:34:08 | 2020-03-21 14:34:08 |
+-----+----------+---------------------+---------------------+
1 row in set (0.00 sec)
```

事务2

```sql
mysql> begin;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from t_base_info where oid = 1 for update;
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```

## 常见锁

在 InnoDB 存储引擎， REPEATABLE READ (可重复读) 隔离级别下，为我们提供了多种锁

表锁，获取锁效率高，有效避免死锁 (破坏了死锁的竞争条件)。但严重影响性能，并发性低，在生产系统上，表锁简直属于灾难

事务1

```sql
begin;
lock table t_base_info read;
```

事务2

```sql
mysql> begin;
Query OK, 0 rows affected (0.00 sec)

mysql> insert into t_base_info(name,create_time,updated_time)values("name",now(),now());
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```

表锁的超时时间不受：innodb_lock_wait_timeout 限制，而是受：lock_wait_timeout 限制

```sql
mysql> show variables like "lock_wait_timeout";
+-------------------+----------+
| Variable_name     | Value    |
+-------------------+----------+
| lock_wait_timeout | 31536000 |
+-------------------+----------+
1 row in set (0.00 sec)
```

单位为秒，默认为31536000 秒，365 天, 可通过下述命令进行修改

```sql
set global lock_wait_timeout = 10;
```

Record 锁
锁定范围：索引记录，如果表中没有索引记录，则会自动创建一个隐藏的聚簇索引
只锁定单条索引记录，获取锁效率稍低，可能会产生死锁， 但并发高，性能好

Gap 锁 顾名思义，锁定的是一个范围
Gap 锁: 锁定一个范围，锁定后该范围不支持新增，其他事物在该范围中insert，delete 需要lock wait 直至释放
适用隔离级别： REPEATABLE READ (可重复读)
在 READ UNCOMMITTED (未提交读)，READ COMMITTED (提交读) 隔离级别下无效

事务1

```sql
begin;
select * from t_base_info where oid < 8 for update;
```

事务2

```sql
begin;
mysql> insert into t_base_info(oid,name,create_time,updated_time)
values(4,"name4",now(),now());
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```

Next-Key 锁,顾名思义，锁定的是一个范围
Next-Lock：Record 锁 + Gap 锁的集合。其中锁定的范围是：当前记录和范围
适用隔离级别：REPEATABLE READ (可重复读)
该锁在 READ UNCOMMITTED (未提交读)，READ COMMITTED (提交读) 隔离级别下无效

事务1

```sql
begin;
select * from t_base_info where oid < 8 for update;
```

事务2

```sql
begin;
mysql> update t_base_info set name = "andyqian008" where oid = 8;
ERROR 1205 (HY000): Lock wait timeout exceeded; try restarting transaction
```

锁超时

```sql
-- 查看锁超时时间
show variables like "innodb_lock_wait_timeout";
-- 修改全局锁超时时间
set global innodb_lock_wait_timeout = 30;
-- 修改当前会话锁超时时间
set session innodb_lock_wait_timeout = 20;
```
