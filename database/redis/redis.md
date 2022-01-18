# redis

```shell
# 查询所有key
127.0.0.1:6379> keys *
# 显示当前库key的数量
127.0.0.1:6379> dbsize
# 所有库key的数量
127.0.0.1:6379> info keyspace

# 关闭 redis server, 如果是默认端口则去掉 -p
# 下面2种方式都可以关闭
redis-cli -p port shutdown
127.0.0.1:6379> shutdown
```
