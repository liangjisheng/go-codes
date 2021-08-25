# redis

[redis](https://redis.io/download)

```shell
wget http://download.redis.io/releases/redis-5.0.7.tar.gz
tar zxf redis-5.0.7.tar.gz
mkdir /usr/local/redis
make PREFIX=/usr/local/redis install
cp redis.conf /usr/local/redis/
cp sentinel.conf /usr/local/redis/
```

vim /usr/local/redis/redis.conf

```conf
daemonize yes     # (如果用supervisor监控的话需要改成no)
port 6379         # 默认是6379 需要安全组开放端口
bind 127.0.0.1    # 远程访问改成 *
dir /usr/local/redis
appendonly yes
requirepass 123456
pidfile /usr/local/redis/redis.pid
logfile /usr/local/redis/redis.log
```

启动 redis
sudo /usr/local/redis/bin/redis-server /usr/local/redis/redis.conf

```shell
# 处理中文乱码问题
redis-cli --raw
# 关闭redis进程
redis-cli shutdown
# 登录redis
redis-cli -a '123456'
```
