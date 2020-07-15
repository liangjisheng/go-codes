#!/bin/bash

# 由于apt官方库里的docker版本可能比较旧，所以先卸载可能存在的旧版本
sudo apt-get remove docker docker-engine docker-ce docker.io

sudo apt-get update

# 安装以下包以使apt可以通过HTTPS使用存储库(repository)
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common

# 添加Docker官方的GPG密钥
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# 使用下面的命令来设置stable存储库
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
# sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main"

sudo apt-get update

# 安装最新版本的Docker CE
sudo apt-get install -y docker-ce


# 列出可用的版本
# apt-cache madison docker-ce

# 安装指定版本
# sudo apt-get install docker-ce=<VERSION>


# 查看docker服务是否启动
# systemctl status docker

# 若未启动，则启动docker服务
# sudo systemctl start docker

# 经典的hello world
# sudo docker run hello-world
