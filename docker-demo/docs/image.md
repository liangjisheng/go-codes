# image

当运行容器时，使用的镜像如果在本地中不存在，docker 就会自动从 docker 镜像仓库中下载，
默认是从 Docker Hub 公共镜像源下载

```shell
#使用 docker images 来列出本地主机上的镜像
docker images
#查看镜像详细信息
docker inspect mysql:5.7
#一个镜像是由多个层（layer）组成的，那么，我们要如何知道各个层的具体内容呢
#通过 docker history 命令，可以列出各个层（layer）的创建信息
docker history mysql:5.7
#如果想要看具体信息，可以添加 --no-trunc 参数
docker history --no-trunc mysql:5.7
#通过 docker save 命令可以导出 Docker 镜像
docker save -o redis.tar redis:latest
#可以通过 docker load 命令导入镜像
docker load -i redis.tar
docker load < redis.tar
#删除镜像
#-f, -force: 强制删除镜像，即便有容器引用该镜像；
#-no-prune: 不要删除未带标签的父镜像；
docker rmi redis
docker image rm redis:latest
#直接通过 ID 删除镜像
docker rmi ee7cbd482336
#在使用 Docker 一段时间后，系统一般都会残存一些临时的、没有被使用的镜像文件，
#可以通过以下命令进行清理：
docker image prune
```

```shell
#去 https://hub.docker.com/ 注册一个账号
#进入命令行，用我们刚刚获取的 Docker ID 以及密码登录
docker login
docker tag python:3 liangjisheng/python:3
docker push liangjisheng/python:3
```

REPOSITORY：表示镜像的仓库源
TAG：镜像的标签
IMAGE ID：镜像ID
CREATED：镜像创建时间
SIZE：镜像大小

同一仓库源可以有多个 TAG，代表这个仓库源的不同个版本，如 ubuntu 仓库源里，有 15.10、14.04
等多个不同的版本，我们使用 REPOSITORY:TAG 来定义不同的镜像
如果你不指定一个镜像的版本标签，例如你只使用 ubuntu，docker 将默认使用 ubuntu:latest 镜像

```shell
#拉取镜像
docker pull ubuntu:13.10
#搜索 httpd 来寻找适合我们的镜像
docker search httpd
```

也可以在 [docker hub](https://hub.docker.com/) 中来搜索镜像

```shell
#删除镜像
docker rmi hello-world
#更新镜像之前，我们需要使用镜像来创建一个容器
#在运行的容器内使用 apt-get update 命令进行更新
docker commit -m="has update" -a="runoob" e218edb10161 runoob/ubuntu:v2
```

-m: 提交的描述信息
-a: 指定镜像作者
e218edb10161：容器 ID
runoob/ubuntu:v2: 指定要创建的目标镜像名

使用我们的新镜像 runoob/ubuntu 来启动一个容器
docker run -t -i runoob/ubuntu:v2 /bin/bash

编写 Dockerfile 文件，在这个文件所在的目录下执行

```shell
docker build -t liangjisheng/ubuntu:v1 .
docker images
docker run -t -i liangjisheng/ubuntu:v1 /bin/bash
#为镜像添加一个新的标签
docker tag 860c279d2fec liangjisheng/ubuntu:dev
```

-t ：指定要创建的目标镜像名
. ：Dockerfile 文件所在目录，可以指定Dockerfile 的绝对路径