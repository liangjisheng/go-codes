# redis

在 Redis 中的一个命令执行过程期间 ，所有服务器接收到的其他命令都必须等待被处理。
因此，对于生产环境的性能来说，调用KEYS 命令是一个危险的操作。 对于这个问题，
可以使用此前案例中所提到的 SCAN 类命令，如SCAN 或 SSCAN ，以在不阻塞服务器的
情况下在 Redis 服务器上遍历键)

```shell
# 查询所有key
127.0.0.1:6379> keys *
127.0.0.1:6379> scan 0
# 显示当前库key的数量
127.0.0.1:6379> dbsize
# 所有库key的数量
127.0.0.1:6379> info keyspace
#查看 db 数量, 默认 16 个 db
127.0.0.1:6379> config get databases
#切换到数据库序号(从 0 开始)为 3 的 db
127.0.0.1:6379> select 3
#清空当前数据库
127.0.0.1:6379> flushdb
#清空所有数据库
127.0.0.1:6379> flushall

# 关闭 redis server, 如果是默认端口则去掉 -p
# 下面2种方式都可以关闭
redis-cli -p port shutdown
127.0.0.1:6379> shutdown
```
