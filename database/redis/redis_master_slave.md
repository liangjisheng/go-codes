# redis cluster

## master slave

下面使用3台机器做个redis1主2从集群搭建

3台机器: 172.20.101.23(主), 172.20.101.24(从), 172.20.101.25(从)

先在三台机器上安装redis到/usr/local/redis

主节点修改配置文件 redis.conf

```conf
daemonize yes
appendonly yes
bind *
requirepass 123456
dir /usr/local/redis
pidfile /usr/local/redis/redis.pid
logfile /usr/local/redis/redis.log
```

2个从节点修改配置 redis.conf

```conf
daemonize yes
appendonly yes
bind *
requirepass 123456
dir /usr/local/redis
pidfile /usr/local/redis/redis.pid
logfile /usr/local/redis/redis.log
# master 节点密码
masterauth 123456
slaveof 172.20.101.23 6379
slave-read-only yes
```

启动3台机器的 redis
sudo /usr/local/redis/bin/redis-server /usr/local/redis/redis.conf

登录主节点查看主从状态, 如果没有连上从节点, 则查看防火墙是否关闭了6379端口
127.0.0.1:6379> info replication

到此时一个简单的1主2从的redis集群便已经搭建完成,如果只是想在一台机器上测试redis集群的话
可以拷贝多份配置文件,修改不同的端口和数据目录等配置,使用不同的配置文件启动redis-server

## sentinel

哨兵模式的redis集群搭建是在主从的基础上进行的,所以如果想搭建哨兵模式的集群
也需要执行上面的步骤

apt 安装哨兵
sudo apt-get install redis-sentinel

3个节点修改 sentinel.conf 配置文件

```conf
bind 0.0.0.0
port 26379
daemonize yes
pidfile /usr/local/redis/sentinel.pid
logfile /usr/local/redis/sentinel.log
dir /usr/local/redis
# 数量(2代表只有两个或两个以上的哨兵认为主服务器不可用的时候，才会进行failover操作)
sentinel monitor mymaster 172.20.101.23 6379 2
sentinel auth-pass mymaster 123456
# 多长时间没有响应认为主观下线(SDOWN)
sentinel down-after-milliseconds mymaster 60000
# 表示如果15秒后, mysater仍没活过来, 则启动failover, 从剩下从节点选取新的主节点
sentinel failover-timeout mymaster 15000
# 指定了在执行故障转移时, 最多可以有多少个从服务器同时对新的主服务器进行同步, 这个数字越小完成故障转移所需的时间就越长
sentinel parallel-syncs mymaster 1
```

3个节点启动 redis-sentinel
sudo /usr/local/redis/bin/redis-sentinel /usr/local/redis/sentinel.conf

到这里redis的哨兵模式就搭建完成了,如果master挂了,那么其中的1个slave将自动升级为master
修改旧master配置文件redis.conf masterauth password 为新的master节点密码
重启旧master的redis,重新加入集群
