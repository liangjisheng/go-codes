# mysql

导出数据库

```sql
mysqldump -u username -p database > /home/ps/database.sql
mysqldump -u mysql -p syncdb > /home/ps/database.sql
```

导出某张表

```sql
mysqldump -u mysql -p syncdb ljstable > /home/ps/table.sql
mysql -u mysql -p debug < /data/usertable.sql
mysqldump -u root -p main account_token_history > /home/ps/backup-mysql/account_token_history_20200814.sql
mysqldump -u root -p main pair_price_history > /home/ps/backup-mysql/pair_price_history_20200814.sql
```

导入数据库

```sql
mysql -u mysql -p syncdb < /home/ps/database.sql
```

或者进入 mysql 控制台

```sql
mysql>source /home/ps/database.sql;
```

## binlog

```sql
-- 查看现有 binlog
show binary logs;
-- 设置 binlog 过期时间
set global expire_logs_days=7;
-- 下面操作会触发过期日志清除操作,开始新的 binlog
flush logs;
mysqladmin flush-logs
```

mysql 从库开启只读, mysql 的 replication 可以做这两个的参数设定

```sql
SET GLOBAL super_read_only = ON;
SET GLOBAL read_only = ON;
```

设置 binlog 过期时间为30天

```sql
set global expire_logs_days=30;
flush logs;
```

将bin.000055之前的binlog清掉:

```sql
mysql>purge binary logs to 'bin.000055';
```

将指定时间之前的binlog清掉:

```sql
mysql>purge binary logs before '2019-09-13 23:59:59';
```

注意, 不要轻易手动去删除 binlog, 会导致 binlog.index 和真实存在的 binlog 不匹配, 而导致 expire_logs_day 失效

redo log 事务日志ib_logfile文件是InnoDB存储引擎产生的  
重做日志是在崩溃恢复期间用于纠正由未完成事务写入的数据的基于磁盘的数据结构  
默认情况下, 重做日志在磁盘上物理表示为一组文件, 名为ib_logfile0和ib_logfile1 MySQL以循环方式写入重做日志文件  
innodb_log_files_in_group 确定ib_logfile文件个数, 命名从 ib_logfile0 开始
如果最后1个 ib_logfile 被写满, 而第一个ib_logfile中所有记录的事务对数据的变更已经被持久化到磁盘中, 将清空并重用之  
为了最大程度避免数据写入时io瓶颈带来的性能问题, MySQL采用了这样一种缓存机制: 当query修改数据库内数据时
InnoDB先将该数据从磁盘读取到内存中, 修改内存中的数据拷贝, 并将该修改行为持久化到磁盘上的事务日志  
(先写redo log buffer, 再定期批量写入), 而不是每次都直接将修改过的数据记录到硬盘内, 等事务日志持久化完成之后  
内存中的脏数据可以慢慢刷回磁盘, 称之为Write-Ahead Logging. 事务日志采用的是追加写入, 顺序io会带来更好的性能优势

从服务器 mysql 终端执行

```sql
mysql>change master to master_host='172.20.101.23',master_user='mysql',master_password='password',master_log_file='mysql-bin.000001',master_log_pos=154;
show master status\G
show slave status\G
```

slave status 这2个字段状态都是 Yes 同步成功  

```txt
Slave_IO_Running: Yes  
Slave_SQL_Running: Yes
```

显示正在执行的sql语句

```sql
show processlist;
select * from information_schema.`PROCESSLIST` where info is not null;
```

增加列

```sql
alter table t2 add d timestamp;
alter table infos add ex tinyint not null default '0';
alter table table_name add field_name field_type;
```

```sql
-- 修改字段名
alter table Persons add Birthday date
-- 改变列的类型
alter table t1 change b b bigint not null;
alter table tb_article modify column name char(50);
```

```sql
-- 删除列
alter table t2 drop column c;
-- 重命名列
alter table t1 change a b integer;
-- 重命名表
alter table t1 rename t2;
```

增加索引

```sql
alter table tablename add index 索引名 (字段名1[，字段名2 …]);
alter table tablename add index indexname (columnname);
alter table tablename drop index indexname;
```

mysql in 查询 根据查询的顺序返回

```sql
select * from test where id in ('4','3','2','1') order by field(id,'4','3','2','1');
```

concat 将多个字符串连接成一个字符串, 返回结果为连接参数产生的字符串，如果有任何一个参数为null，则返回值为null

```sql
select concat(id, address, mission_name) as info from address_nft;
-- 可以加一个逗号作为分隔符
select concat(id, ',', address, ',', mission_name) as info from address_nft;
```

concat_ws 和 concat()一样，将多个字符串连接成一个字符串，但是可以一次性指定分隔符

```sql
select concat_ws(',', id, address, mission_name) as info from address_nft;
```

group_concat 将group by产生的同一个分组中的值连接起来，返回一个字符串结果
group_concat([distinct] 要连接的字段 [order by 排序字段 asc/desc  ] [separator '分隔符'])

```sql
select address, group_concat(mission_name) from address_nft group by address;
select address, group_concat(concat_ws(',', mission_name, status)) from address_nft group by address;
```

mysql 常用的命令

```sql
-- 显示指定数据库的创建语句
show create database zkswapv3;
show create schema zkswapv3;
-- 显示表中所有列信息
show full columns tables_name;
-- 查看MySQL版本
select version();
-- 查看当前用户
select current_user();
-- 显示单表信息
show table status like "table_name";
-- 显示正在操作数据库的进程数
show processlist;
-- 显示表中的所有索引
show index from t_base_data;
-- 查看查询语句的执行情况，常用于SQL优化
-- explain 查询语句
-- 显示当前时间
select now();
-- 显示指定字符长度
select char_length('test');
-- 格式化日期
select date_format(now(), '%y-%m-%d');
select DATE_FORMAT(now(), '%y-%m-%d %H:%i:%s');
-- 添加/减少日期时间 
-- DATE_ADD(date,interval expr unit) DATE_SUB(date,interval expr unit)
-- unit：表示单位，支持毫秒(microsecond)，秒(second)，小时(hour)，天(day)，周(week)，年(year)等
select date_add(now(), interval 1 day);
-- 类型转换
select cast(18700000000 as char);
-- md5(data)
select md5("test");
-- 字符串连接
select concat("test","big");
-- 获取系统当前时间的时间戳，类型: long 单位: s
select unix_timestamp(now()), unix_timestamp(current_timestamp()), unix_timestamp(sysdate());
-- 此时时间精度是s，也可以增加精度，给函数加上参数，表示s后面的小数位数，例如参数3，此时为ms
select unix_timestamp(now(3)), unix_timestamp(current_timestamp(3)), unix_timestamp(sysdate(3));
-- 如果直接输出毫秒单位的时间戳，就是去掉上面中间的小数点，可以借助 replace 函数
select replace(unix_timestamp(now(3)), '.', ''), replace(unix_timestamp(current_timestamp(3)), '.', ''),
       replace(unix_timestamp(sysdate(3)), '.', '');

-- 获取系统当前时间，类型：timestamp 格式yyyy-MM-dd HH:mm:ss
-- 2019-01-04 20:37:19
-- 三者基本没有区别，稍微一点的区别在于：NOW(),CURRENT_TIMESTAMP()都表示SQL开始执行的时间；SYSDATE()表示执行此SQL时的当前时间
select now(), current_timestamp(), sysdate();
select NOW(),CURRENT_TIMESTAMP(),SYSDATE(),SLEEP(2),NOW(),CURRENT_TIMESTAMP(),SYSDATE();
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

mysql time

```sql
select unix_timestamp('2022-03-23 18:00:00');
select unix_timestamp();
select from_unixtime(1648029600, '%Y-%m-%d %H:%i:%S');
select from_unixtime(1648029600);
```

查看 mysql 配置文件

```shell
# 通过进程查看是否启动命令是否指定配置文件
ps -aux |grep mysql |grep 'my.cnf'
# 若无，则说明MySQL启动时采用的是默认配置文件，输入命令
mysql --help |grep 'my.cnf'
mysqld --verbose --help |grep -A 1 'Default options'
```

清除 binlog 自动删除

```conf
; 查看binlog文件
show binary logs
; 查看binlog文件过期时间
show variables like 'expire_logs_days';

; 永久生效：修改MySQL配置文件my.cnf，配置binlog的过期时间，重启生效
expire_logs_days=30

; -- 设置过期时间为30天 临时生效：即时生效，重启后失效 mysql 终端
set global expire_logs_days=30;
```

手动删除
手动删除前需要先确认主从库当前在用的 binlog 文件
主库: show master status;
从库: show slave status;

假设当前在用的binlog文件为master-bin.000277，现需要删除master-bin.000277之前的所有binlog日志文件(不删master-bin.000277):

```shell
# mysql 终端
purge master logs to 'master-bin.000277';
purge master logs before '2017-05-01 13:09:51';
purge binary logs to 'master-bin.000277';
purge binary logs before '2017-05-01 13:09:51';
```
