#!/bin/bash

# 由于apt官方库里的docker版本可能比较旧，所以先卸载可能存在的旧版本
sudo apt-get remove docker docker-engine docker-ce docker.io containerd runc

sudo apt-get update

# 安装以下包以使apt可以通过HTTPS使用存储库(repository)
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common

# 添加Docker官方的GPG密钥
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
# 验证现在是否拥有带有指纹的密钥
sudo apt-key fingerprint 0EBFCD88

# 使用下面的命令来设置stable存储库
#sudo add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/ $(lsb_release -cs) stable"
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

sudo apt-get update

# 安装最新版本的Docker CE
sudo apt-get install -y docker-ce docker-ce-cli

# 登录个人阿里云 https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors 使用镜像加速器地址
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://xxx.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

#检查加速器是否生效
docker info

# 列出可用的版本
# sudo apt-cache madison docker-ce

# 安装指定版本
# sudo apt-get install docker-ce=<VERSION>

# 查看docker服务是否启动
# sudo systemctl status docker
# 若未启动，则启动docker服务
# sudo systemctl start docker

# 经典的hello world
# sudo docker run hello-world

#卸载 docker
#sudo apt-get purge docker-ce
#删除镜像、容器、配置文件等内容：
#sudo rm -rf /var/lib/docker
