# redis cluster

[cluster](http://www.redis.cn/topics/cluster-tutorial.html)
[cluster](https://www.cnblogs.com/mafly/p/redis_cluster.html)
[cluster](https://www.jianshu.com/p/8045b92fafb2)

下载redis并编译完,复制6份到 /usr/local/redis-cluster
修改6个配置文件, 端口分别为 7000-7005

```conf
port 7005
# 允许集群启动
cluster-enabled yes
# 集群配置文件名,集群启动后自动生成
cluster-config-file nodes7005.conf
# 集群节点之间多少毫秒无法连接后判定节点挂掉
cluster-node-timeout 10000
pidfile /var/run/redis_7005.pid
logfile /usr/local/redis-cluster/redis_7005.log
dir /usr/local/redis-cluster/7005/data/
appendonly yes
```

启动 redis

```shell
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7000/redis.conf
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7001/redis.conf
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7002/redis.conf
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7003/redis.conf
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7004/redis.conf
sudo /usr/local/redis-cluster/bin/redis-server /usr/local/redis-cluster/7005/redis.conf
```

创建集群(注意, 一次创建永久使用, 以后不需要再创建)
执行下面命令后会将16384个槽位平均分配给三组节点(3主3从), 输入Y确认
如果有密码可以加上-a参数

```shell
# 创建集群
/usr/local/redis-cluster/bin/redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1
# 增加节点
redis-cli --cluster add-node 127.0.0.1:7006 127.0.0.1:7007
# 移除节点, 第一个参数是任意一个节点的地址, 第二个参数是你想要移除的节点ID
# 如果是移除主节点, 需要确保这个节点是空的, 如果不是空的则需要将这个节点上的数据重新分配到其他节点上
redis-cli --cluster del-node 127.0.0.1:7001 <nodeId>
```

从某个实例登录集群, -c 表示集群模式

```shell
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7000
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7001
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7002
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7003
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7004
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7005
```

登录后查看集群信息

```shell
127.0.0.1:7000> cluster info
127.0.0.1:7000> cluster nodes
```

关闭节点

```shell
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7000 shutdown
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7001 shutdown
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7002 shutdown
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7003 shutdown
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7004 shutdown
/usr/local/redis-cluster/bin/redis-cli -c -h 127.0.0.1 -p 7005 shutdown
```
