# limit

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

-- 新增记录
insert into `t_base_user` (`name`, `email`, `age`, `telephone`, `status`, `created_at`, `updated_at`)
values ('111111', 'test@gmail.com', '111', '12345678901', '1', now(), now());

-- 复制数据
insert into t_base_user(name,email,age,telephone,status) select name,email,age,telephone,status from t_base_user;
```

```sql
-- 从起始为0的位置,往后获取100条记录
LIMIT 0,100; -- 等价于下面语句
LIMIT 100;
-- 从起始角标为10的位置,往后获取100条记录
LIMIT 10,100;
```

问: LIMIT 0,100 与 LIMIT 1000,100 的执行效率是一样吗？ 一样为什么？不一样又为什么？

```sql
explain select * from t_base_user order by oid desc limit 0,100;
explain select * from t_base_user order by oid desc limit 1000,100;
```

我们会发现扫描的行数是完全不一样的, 第二条语句其实MySQL实际扫描 1100 行记录,然后只返回 100 条记录,将前面的 1000 条记录都丢弃了
在offset比较大的时候,效率会非常低,所以,对LIMIT优化,要么限制分页的数量,要么降低offset的大小, 下面是2种优化方式

```sql
select * from t_base_user where oid between 1000 and 1100;
select * from t_base_user where oid >= 1000 and oid <=1100;
```
