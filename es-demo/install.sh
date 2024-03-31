#!/bin/bash

# https://www.jianshu.com/p/1d2ddb92f6fb

wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.6.0-linux-x86_64.tar.gz
tar zxf elasticsearch-7.6.0-linux-x86_64.tar.gz

# 修改 es 所占内存为 200m
vim elasticsearch-7.6.0/config/jvm.options
# -Xms200m
# -Xmx200m

# 修改绑定地址
vim elasticsearch-7.6.0/config/elasticsearch.yml
# network.host: 0.0.0.0
# 同时需要修改 es 所在机器的内网ip和外网ip
# discovery.seed_hosts: ["127.0.0.1", "10-255-20-235", "117.51.148.112"]
# 允许跨域
# http.cors.enabled: true
# http.cors.allow-origin: "*"


# 不能root账户来启动 不然会报错 需要为elasticsearch新建一个系统运行账号
# 新建一个elasticsearch的用户组
groupadd elasticsearch

# 在elasticsearch用户组下面建立一个elasticsearch的用户
useradd -g elasticsearch elasticsearch

# 将elasticsearch目录的所有者给刚刚建立的账号
chown -R elasticsearch:elasticsearch elasticsearch-7.6.0/

# 切换到刚刚的账号启动elasticsearch
su elasticsearch

# 启动 es
# ./bin/elasticsearch

# 后台启动 es
./bin/elasticsearch -d

# 浏览器访问
# http://117.51.148.112:9200/

# 安装 elasticsearch-head 非必须, 如果要安装,需要先安装 node
git clone git://github.com/mobz/elasticsearch-head.git
cd elasticsearch-head
npm install

# 修改绑定地址 4374行
sudo vi _site/app.js
# 后台启动
nohup npm run start > nohup 2>&1 &
# 浏览器访问
# http://117.51.148.112:9100/

# 安装 kibana
curl -O https://artifacts.elastic.co/downloads/kibana/kibana-7.6.0-linux-x86_64.tar.gz
tar -xzf kibana-7.6.0-linux-x86_64.tar.gz
chown -R elasticsearch:elasticsearch kibana-7.6.0-linux-x86_64/
su elasticsearch
cd kibana-7.6.0-linux-x86_64/

vim config/kibana.yml
# 修改监听主机
# server.host: "0.0.0.0"

./bin/kibana

# 后台启动
nohup ./bin/kibana > nohup 2>&1 &

# 浏览器访问
# http://117.51.148.112:5601/app/kibana#/home?_g=()
