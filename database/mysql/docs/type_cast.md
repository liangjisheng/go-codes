# mysql 类型转换

新建表

```sql
CREATE TABLE `t_base_user` (
    `oid` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(30) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'name',
    `email` varchar(30) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'email',
    `age` int DEFAULT NULL COMMENT 'age',
    `telephone` varchar(30) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'telephone',
    `status` tinyint DEFAULT NULL COMMENT '0 无效 1 有效',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`oid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

alter table t_base_user add index idx_name(name);
alter table t_base_user add index idx_telephone(telephone);

INSERT INTO `t_base_user` (`name`, `email`, `age`, `telephone`, `status`, `created_at`, `updated_at`)
VALUES ('111111', 'test@gmail.com', '111', '12345678901', '1', now(), now());

-- 查看执行计划
explain select * from t_base_user where telephone=12345678901;
-- 上述语句并没有走索引,而是全表扫描
```

当操作符与不同类型的操作数一起使用时，会发生类型转换以使操作数兼容

```sql
-- 现在语句已经走索引了
explain select * from t_base_user where telephone='12345678901';
```

如何避免隐式类型转换

```sql
-- 使用CAST函数显示转换
select 38.8, cast(38.8 as char);
select * from t_base_user where telephone=cast(12345678901 as char);
```

## cast

CAST(expr AS type), 这里需要注意的是type类型不支持所有的数据类型，而是支持特定的数据类型

```sql
mysql> SELECT CAST('1024' AS int);
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'int)' at line 1
```

example

```sql
mysql> select cast('999-11-11' as DATE);
+---------------------------+
| cast('999-11-11' as DATE) |
+---------------------------+
| 0999-11-11                |
+---------------------------+
1 row in set (0.00 sec)

mysql> select cast('01-11-11' as DATE);
+--------------------------+
| cast('01-11-11' as DATE) |
+--------------------------+
| 2001-11-11               |
+--------------------------+
1 row in set (0.00 sec)

mysql> select version();
+-----------+
| version() |
+-----------+
| 5.7.20    |
+-----------+
1 row in set (0.00 sec)

mysql> SELECT CAST('ANDYQIAN' AS DECIMAL);
+-----------------------------+
| CAST('ANDYQIAN' AS DECIMAL) |
+-----------------------------+
|                           0 |
+-----------------------------+
1 row in set, 1 warning (0.00 sec)

mysql> select cast('2017-12-14' as DATE);
+----------------------------+
| cast('2017-12-14' as DATE) |
+----------------------------+
| 2017-12-14                 |
+----------------------------+
1 row in set (0.00 sec)

mysql> select cast('12:00:00' as TIME);
+--------------------------+
| cast('12:00:00' as TIME) |
+--------------------------+
| 12:00:00                 |
+--------------------------+
1 row in set (0.00 sec)

mysql> select cast('2017-12-14 00:11:11' as DATETIME);
+-----------------------------------------+
| cast('2017-12-14 00:11:11' as DATETIME) |
+-----------------------------------------+
| 2017-12-14 00:11:11                     |
+-----------------------------------------+
1 row in set (0.00 sec)

mysql> select cast('-1024' as SIGNED);
+-------------------------+
| cast('-1024' as SIGNED) |
+-------------------------+
|                   -1024 |
+-------------------------+
1 row in set (0.00 sec)

mysql> select cast('-1024' as UNSIGNED);
+---------------------------+
| cast('-1024' as UNSIGNED) |
+---------------------------+
|      18446744073709550592 |
+---------------------------+
1 row in set, 1 warning (0.00 sec)

mysql> select cast('18.11' as DECIMAL(18,2));
+--------------------------------+
| cast('18.11' as DECIMAL(18,2)) |
+--------------------------------+
|                          18.11 |
+--------------------------------+
1 row in set (0.00 sec)
```