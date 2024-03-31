#!/bin/bash

# https://www.cnblogs.com/linkstar/p/10341685.html

wget https://github.com/nsqio/nsq/releases/download/v1.2.0/nsq-1.2.0.linux-amd64.go1.12.9.tar.gz
tar zxf nsq-1.2.0.linux-amd64.go1.12.9.tar.gz
cd nsq-1.2.0.linux-amd64.go1.12.9/bin

# tcp client 4160 
# http client 4161
nohup ./nsqlookupd > nsqlookupd_nohup 2>&1 &

# tcp client 4150
# http client 4151
# https client 4152
nohup ./nsqd --lookupd-tcp-address=127.0.0.1:4160 > nsqd_nohup 2>&1 &

# http client 4171
nohup ./nsqadmin --lookupd-http-address=127.0.0.1:4161 > nsqadmin_nohup 2>&1 &

# 会创建一个test主题，并发送一个hello world消息
# curl -d 'hello world' 'http://127.0.0.1:4151/pub?topic=test'

# 浏览器访问
# http://117.51.148.112:4171/

# 消费test中刚才的消息，并输出到服务器/tmp目录中
# ./nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
