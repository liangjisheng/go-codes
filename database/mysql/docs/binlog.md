# binlog

```sql
-- 二进制日志位置
show variables like '%log_bin%';
-- 查看现有 binlog
show binary logs;
-- 查看当前使用的日志格式
show variables like '%format%';
-- 查看当前使用的二进制日志文件
show master status;
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

