# mysql

[快速的 drop 掉一个超过 100G 的大表](https://blog.duhbb.com/2023/03/27922.html)
[count](https://blog.duhbb.com/2023/03/28813.html)

```shell
-- 导出数据库
mysqldump -u username -p database > /home/ps/database.sql
mysqldump -u mysql -p syncdb > /home/ps/database.sql

-- 导出某张表
mysqldump -u mysql -p syncdb ljstable > /home/ps/table.sql
mysql -u mysql -p debug < /data/usertable.sql
mysqldump -u root -p main account_token_history > /home/ps/backup-mysql/account_token_history_20200814.sql
mysqldump -u root -p main pair_price_history > /home/ps/backup-mysql/pair_price_history_20200814.sql

-- 导入数据库
mysql -u mysql -p syncdb < /home/ps/database.sql
-- 或者进入 mysql 控制台
mysql>source /home/ps/database.sql;
```

```sql
-- 显示正在执行的sql语句
show processlist;
select * from information_schema.`PROCESSLIST` where info is not null;

-- 增加列
alter table t2 add d timestamp;
alter table infos add ex tinyint not null default '0';
alter table table_name add field_name field_type;

-- 修改字段名
alter table Persons add Birthday date
-- 改变列的类型
alter table t1 change b b bigint not null;
alter table tb_article modify column name char(50);

-- 删除列
alter table t2 drop column c;
-- 重命名列
alter table t1 change a b integer;
-- 重命名表
alter table t1 rename t2;

-- 增加索引
alter table tablename add index 索引名 (字段名1[，字段名2 …]);
alter table tablename add index indexname (columnname);
alter table tablename drop index indexname;

-- mysql in 查询 根据查询的顺序返回
select * from test where id in ('4','3','2','1') order by field(id,'4','3','2','1');
```

mysql 表设计的注意事项  
在设计数据表时，一定要注意该字段存储的内容，如果允许设置表情，则一定不能使用utf8，而是使用utf8mb4  
选择合适的类型，保存手机号的字段，用varchar(20)就已经足够了，保存Boolean类型，使用tinyint就够了，而不需要设计为int，甚至bigint  

主外键类型不一致，说起来，可能会不相信，但在数据库表设计时，稍不留神，就不一致，埋下隐式类型转换的坑  
设计数据库表时，注释也非常重要！一定要记住加注释，无论是表，还是字段，索引，都必须加上注释  
状态类型用 tinyint，例如 性别等  
时间日期使用 datetime,timestamp 类型，使用datetime类型可读性高些  
尽量不要使用text和blob数据类型，特别是blob

事务的4大特性， 只有InnoDB存储引擎才支持事务  
原子性（Atomicity）：事务作为一个整体，也就是说事务中的对数据库的操作要么全部被执行，要么都不执行
一致性（Consistency）: 事务应确保数据库的状态从一个一致状态转变为另一个一致状态
隔离性（Isolation）：多个事务并发执行时，一个事务的执行不应影响其他事务的执行
持久性（Durability）：已被提交的事务对数据库的修改应该永久保存在数据库中

```sql
show variables like "autocommit";
-- 禁用自动提交
set autocommit = 0;
-- 开启自动提交
set autocommit = 1;

-- begin 开始一个新的事务
-- commit 提交事务
-- rollback 回滚事务
```

```shell
# 查看 mysql 配置文件
# 通过进程查看是否启动命令是否指定配置文件
ps -aux |grep mysql |grep 'my.cnf'
# 若无，则说明MySQL启动时采用的是默认配置文件，输入命令
mysql --help |grep 'my.cnf'
mysqld --verbose --help |grep -A 1 'Default options'
```

```sql
-- AHI 在这里就不细说，主要是当 B+tree 的层级变高时，为避免 B+tree 逐层搜索，AHI 能根据某个检索条件，直接查询到对应的数据页，跳过逐层定位的步骤。其次 AHI 会占用 1/16 的 Buffer Pool 的大小，如果线上表数据不是特别大，不是超高并发，不建议将开启 AHI，可以考虑关闭 AHI 功能
show global variables like 'innodb_adaptive_hash_index';
-- 关闭
set global innodb_adaptive_hash_index=off;

-- datadir = /data/var/lib/mysql/ 其中一个数据库是 web3_test, 执行下列命令得到表结构和索引
system ls -l /data/var/lib/mysql/web3_test
```

## show

```sql
show engines;
show variables like "%storage_engine%";
```
