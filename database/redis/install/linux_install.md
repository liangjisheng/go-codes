# redis

[redis](https://redis.io/download)

```shell
cd /usr/local
sudo wget https://github.com/redis/redis/archive/refs/tags/7.2.4.tar.gz
sudo tar zxf 7.2.4.tar.gz
cd redis-7.2.4
sudo make PREFIX=/usr/local/redis-7.2.4 install
```

sudo vim /usr/local/redis-7.2.4/redis.conf

```conf
daemonize yes     # (如果用supervisor监控的话需要改成no)
port 6379         # 默认是6379 需要安全组开放端口
bind 127.0.0.1    # 远程访问改成 *
dir /usr/local/redis-7.0.11
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