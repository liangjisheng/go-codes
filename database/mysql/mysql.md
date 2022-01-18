# mysql

导出数据库
mysqldump -u username -p database > /home/ps/database.sql
mysqldump -u mysql -p syncdb > /home/ps/database.sql

导出某张表
mysqldump -u mysql -p syncdb ljstable > /home/ps/table.sql
mysql -u mysql -p debug < /data/usertable.sql
mysqldump -u root -p main account_token_history > /home/ps/backup-mysql/account_token_history_20200814.sql
mysqldump -u root -p main pair_price_history > /home/ps/backup-mysql/pair_price_history_20200814.sql

导入数据库
mysql -u mysql -p syncdb < /home/ps/database.sql
进入 mysql 控制台
mysql>source /home/ps/database.sql;

复制表
CREATE TABLE targetTable LIKE sourceTable;
INSERT INTO targetTable SELECT * FROM sourceTable;

查看现有 binlog
show binary logs;

设置 binlog 过期时间
set global expire_logs_days=7;

会触发过期日志清除操作,开始新的 binlog
flush logs;
mysqladmin flush-logs

创建用户只能本地访问
create user mysql@'localhost' identified by 'qtum100$';
create user mysql@'%' identified by 'qtum100$';

创建user02,可以远程访问
create user readonly@'%' identified by '346ce5ba6024';

删除用户
drop user 'username'@'host';

修改密码
set password for 'user01'@'localhost'=password('anotherpassword');
set password for 'root'@'%'=password('qtum100$');

mysql5.x 修改密码
set password for root@localhost = password('newpass');
mysql8.x 修改密码
alter user 'root'@'localhost' identified by 'password';
使用5.x版本的密码加密方式
alter user 'root'@'localhost' identified with mysql_native_password by 'qtum100$';
alter user 'zabbix'@'localhost' identified with mysql_native_password by 'zabbix';

授予 select 权限
grant select on dev.* to 'mysql'@'%';

授予管理 test db 的全部权限
grant all privileges on test.* to user01;

grant all privileges on *.* to mysql;

撤销用户权限
revoke select on dev.* from 'readonly'@'%';

刷新权限
flush privileges;

查看某个用户权限
show grants for username;

mysql数据库用户在创建的时候就会赋予USAGE权限, 这个权限很小几乎为0
只能连接数据库和查询information_schema的权限
不过这个权限也很奇怪, 你无法revoke

mysql从库开启只读
mysql 的 replication 可以做这两个的参数设定
SET GLOBAL super_read_only = ON;
SET GLOBAL read_only = ON;

设置 binlog 过期时间为30天
set global expire_logs_days=30;
flush logs;

将bin.000055之前的binlog清掉:
mysql>purge binary logs to 'bin.000055';

将指定时间之前的binlog清掉:
mysql>purge binary logs before '2019-09-13 23:59:59';

注意，不要轻易手动去删除binlog, 会导致binlog.index和真实存在的binlog不匹配, 而导致expire_logs_day失效

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
mysql>change master to master_host='172.20.101.23',master_user='mysql',master_password='qtum100$',master_log_file='mysql-bin.000001',master_log_pos=154;

show master status\G
show slave status\G

slave status 这2个字段状态都是 Yes 同步成功
Slave_IO_Running: Yes
Slave_SQL_Running: Yes

显示正在执行的sql语句
show processlist;
select * from information_schema.`PROCESSLIST` where info is not null;
