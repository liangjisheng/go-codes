# 继承的 ubuntu 镜像的版本号
FROM ubuntu:18.04

# 将本地 workdir 的文件复制到镜像内部, 另外有个 ADD 指令和这个效果类似
COPY ./sources.list /etc/apt/sources.list

# 在镜像内部执行命令, 这里就是更新版本和安装 nginx
RUN apt-get update && apt-get install -y nginx

# 修改 nginx 默认的 index.html 的内容
RUN echo 'Hi, I am in your container' > /var/www/html/index.html

# 镜像对外暴露80端口
EXPOSE 80