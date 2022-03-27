# slow sql

在 MySQL 中，慢查询日志默认为OFF状态，通过如下命令进行查看

```sql
show variables like "slow_query_log";
-- 开启慢查询日志
set global slow_query_log = "ON";
```

其中slow_query_log_file属性，表示慢查询日志存储位置，其日志默认名称为 host 名称

```sql
mysql> show variables like "slow_query_log_file";
+---------------------+----------------------------------------------+
| Variable_name       | Value                                        |
+---------------------+----------------------------------------------+                                     |
| slow_query_log_file | /usr/local/mysql/data/hostname.log |
+---------------------+----------------------------------------------+
2 rows in set (0.01 sec)

-- 修改日志路径
set global slow_query_log_file = ${path}/${filename}.log;
```

慢查询 查询时间，当SQL执行时间超过该值时，则会记录在slow_query_log_file 文件中，其默认为 10 ，最小值为 0，(单位:秒)
当设置值小于0时，默认为 0

```sql
mysql> show variables like "long_query_time";
+-----------------+----------+
| Variable_name   | Value    |
+-----------------+----------+
| long_query_time | 0.500000 |
+-----------------+----------+
1 row in set (0.00 sec)

-- 修改
set global long_query_time = 5;
```

通过上述设置后，退出当前会话或者开启一个新的会话，执行如下命令:
select sleep(6);
该 SQL 则会进入慢查询日志中

注意事项

在 MySQL 中，慢查询日志中默认不记录管理语句，如：
alter table, analyze table，check table等。
不过可通过以下属性进行设置：

```sql
mysql> set global log_slow_admin_statements = "ON";
Query OK, 0 rows affected (0.00 sec)
```

在 MySQL 中，还可以设置将未走索引的SQL语句记录在慢日志查询文件中(默认为关闭状态)

```sql
mysql> set global log_queries_not_using_indexes = "ON";
Query OK, 0 rows affected (0.00 sec)
```

在MySQL中，日志输出格式有支持：FILE(默认)，TABLE 两种，可进行组合使用

```sql
set global log_output = "FILE,TABLE";
```

MySQL 中还提供了一个比较好的工具 mysqldumpslow 来分析慢查询日志文件

## show profiles

show profiles 这个命令非常强大，能清晰的展示每条SQL的持续时间。通常结合show profile 命令可以更加详细的展示其耗时信息
这样就能很容易的分析出，到底慢在哪个环节了。比较遗憾的是，在MySQL中，该命令默认是关闭状态的

开启

```sql
set profiling = ON;
-- 或者
set profiling = 1;
```

查看是否生效

```sql
mysql> show variables like "profiling";
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| profiling | ON |
+---------------+-------+
1 row in set (0.00 sec)
```

show profiles 其作用为显示当前会话服务器最新收到的15条SQL的性能信息

```sql
mysql> show profiles;
+----------+------------+---------------------------------+
| Query_ID | Duration   | Query                           |
+----------+------------+---------------------------------+
|        1 | 0.00385450 | show variables like "profiling" |
|        2 | 0.00170050 | show variables like "profiling" |
|        3 | 0.00038025 | select * from t_base_user       |
+----------+------------+---------------------------------+
```

注意
show profiles 语句 默认显示的是服务端接收到的最新的15条语句。
我们可以通过以下语句进行修改默认值：
set profiling_history_size =20;
profiling_history_size最大取值取值范围为[0,100]

查看Query_ID 等于 3 的详细持续时间构成

```sql
mysql> show profile for query 3;
+----------------------+----------+
| Status               | Duration |
+----------------------+----------+
| starting             | 0.000081 |
| checking permissions | 0.000012 |
| Opening tables       | 0.000028 |
| init                 | 0.000029 |
| System lock          | 0.000017 |
| optimizing           | 0.000006 |
| statistics           | 0.000025 |
| preparing            | 0.000018 |
| executing            | 0.000004 |
| Sending data         | 0.000087 |
| end                  | 0.000007 |
| query end            | 0.000012 |
| closing tables       | 0.000013 |
| freeing items        | 0.000023 |
| cleaning up          | 0.000021 |
+----------------------+----------+
15 rows in set, 1 warning (0.00 sec)
```

也可以指定展示结果

可选参数如下:

all： 展示所有信息。
block io： 展示io的输入输出信息。
context switches： 展示线程的上线文切换信息。
cpu ：显示SQL 占用的CPU信息。
ipc： 显示统计消息的发送与接收计数信息。
page faults：显示主要与次要的页面错误。
memory：本意是显示内存信息，但目前还未实现。
swaps： 显示交换次数。
sources：显示源代码中的函数名称，以及函数发生的文件的名称和行

```sql
mysql> show profile block io,cpu for query 3;
+----------------------+----------+----------+------------+--------------+---------------+
| Status               | Duration | CPU_user | CPU_system | Block_ops_in | Block_ops_out |
+----------------------+----------+----------+------------+--------------+---------------+
| starting             | 0.000081 | 0.000036 |   0.000044 |            0 |             0 |
| checking permissions | 0.000012 | 0.000005 |   0.000006 |            0 |             0 |
| Opening tables       | 0.000028 | 0.000013 |   0.000015 |            0 |             0 |
| init                 | 0.000029 | 0.000013 |   0.000016 |            0 |             0 |
| System lock          | 0.000017 | 0.000008 |   0.000009 |            0 |             0 |
| optimizing           | 0.000006 | 0.000002 |   0.000003 |            0 |             0 |
| statistics           | 0.000025 | 0.000011 |   0.000013 |            0 |             0 |
| preparing            | 0.000018 | 0.000008 |   0.000010 |            0 |             0 |
| executing            | 0.000004 | 0.000002 |   0.000002 |            0 |             0 |
| Sending data         | 0.000087 | 0.000040 |   0.000048 |            0 |             0 |
| end                  | 0.000007 | 0.000003 |   0.000003 |            0 |             0 |
| query end            | 0.000012 | 0.000006 |   0.000007 |            0 |             0 |
| closing tables       | 0.000013 | 0.000005 |   0.000006 |            0 |             0 |
| freeing items        | 0.000023 | 0.000011 |   0.000013 |            0 |             0 |
| cleaning up          | 0.000021 | 0.000009 |   0.000011 |            0 |             0 |
+----------------------+----------+----------+------------+--------------+---------------+
15 rows in set, 1 warning (0.00 sec)
```

也可以通过 limit 选项，来显示指定的行数

```sql
mysql> show profile block io,cpu for query 3 limit 2;
+----------------------+----------+----------+------------+--------------+---------------+
| Status               | Duration | CPU_user | CPU_system | Block_ops_in | Block_ops_out |
+----------------------+----------+----------+------------+--------------+---------------+
| starting             | 0.000081 | 0.000036 |   0.000044 |            0 |             0 |
| checking permissions | 0.000012 | 0.000005 |   0.000006 |            0 |             0 |
+----------------------+----------+----------+------------+--------------+---------------+
2 rows in set, 1 warning (0.00 sec)
```