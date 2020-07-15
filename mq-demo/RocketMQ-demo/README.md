# README

启动RocketMQ前需要先装java

```sh
wget http://mirrors.tuna.tsinghua.edu.cn/apache/rocketmq/4.6.0/rocketmq-all-4.6.0-bin-release.zip
unzip rocketmq-all-4.6.0-bin-release.zip
cd rocketmq-all-4.6.0-bin-release
# 修改占用内存大小 改成 128m
vim bin/runserver.sh
vim bin/runbroker.sh

# 启动 Name Server
nohup sh bin/mqnamesrv &
# 查看启动日志
tail -f ~/logs/rocketmqlogs/namesrv.log

# 创建 broker 配置文件, 使用公网ip将broker注册到 name server
sh mqbroker -m > broker.p
# brokerIP1=117.51.148.112
vim broker.p

# 启动 broker
nohup ./mqbroker -c broker.p &
# 查看启动日志
tail -f ~/logs/rocketmqlogs/broker.log
```
