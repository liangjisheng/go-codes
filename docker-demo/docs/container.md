# container

```shell
#拉取镜像
docker pull ubuntu
#查看有哪些容器正在运行
docker ps
#查看所有的容器
docker ps -a
#在宿主主机内使用 docker logs 命令，查看容器内的标准输出
docker logs e15f2779ea30
#停止容器, 后面也可以跟 NAMES
docker stop e15f2779ea30
#使用 docker start 启动一个已停止的容器
docker start b750bbbcfd88
#重启容器
docker restart <容器 ID>
#-t : 设置关闭容器的限制时间，若超时未能关闭，则使用 kill 命令强制关闭，
# 默认值为 10s，这个时间用于容器保存自己的状态。
docker restart -t=5 redis
#进入容器，如果从这个容器退出，容器不会停止， attach 命令则会停止容器
docker exec -it 243c32535da7 /bin/bash
#导出容器 1e560fca3906 快照到本地文件 ubuntu.tar
docker export 1e560fca3906 > ubuntu.tar
#导入容器快照，将快照文件 ubuntu.tar 导入到镜像 test/ubuntu:v1
cat docker/ubuntu.tar | docker import - test/ubuntu:v1
#也可以通过指定 URL 或者某个目录来导入
docker import http://example.com/exampleimage.tgz example/imagerepo
#删除容器使用 docker rm 命令
docker container rm <容器 ID>
docker rm <容器 ID>
#-f 可强制删除一个正在运行的容器
docker rm -f <容器 ID>
#删除停止的容器
docker container rm e15f2779ea30 c4fc39f526b4
#清理掉所有处于终止状态的容器
docker container prune
#看已经停止运行的容器
docker container ls -a
```

```shell
#运行 hello world
docker run hello-world
#运行 ubuntu, 并进入容器
#-t: 在新容器内指定一个伪终端或终端
#-i: 允许你对容器内的标准输入 (STDIN) 进行交互
docker run -it ubuntu /bin/bash
#上面都是一次性运行的命令，退出后就容器就结束了。使用以下命令创建一个以进程方式运行的容器
docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
#在大部分的场景下，我们希望 docker 的服务是在后台运行的，我们可以过 -d 指定容器的运行模式
docker run -itd --name ubuntu-test ubuntu /bin/bash
```

运行一个 web 应用容器

```shell
docker pull training/webapp
#-d:让容器在后台运行
#-P:将容器内部使用的网络端口随机映射到我们使用的主机上
docker run -d -P training/webapp python app.py
docker ps
#也可以通过 -p 参数来设置不一样的端口
#容器内部的 5000 端口映射到我们本地主机的 5001 端口上
docker run -d -p 5001:5000 training/webapp python app.py
#使用 docker port 可以查看指定容器的某个确定端口映射到宿主机的端口号
docker port bdf16fc374aa
#通过名字查看
docker port bold_sanderson

#可以指定容器绑定的网络地址，比如绑定 127.0.0.1
docker run -d -p 127.0.0.1:5001:5000 training/webapp python app.py
#上面的例子中，默认都是绑定 tcp 端口，如果要绑定 UDP 端口，可以在端口后面加上 /udp
docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py
#当我们创建一个容器的时候，docker 会自动对它进行命名。
#另外，我们也可以使用 --name 标识来命名容器
docker run -d -P --name runoob training/webapp python app.py
```

```shell
#docker logs [ID或者名字] 可以查看容器内部的标准输出
docker logs bdf16fc374aa
#-f: 让 docker logs 像使用 tail -f 一样来输出容器内部的标准输出
docker logs -f bdf16fc374aa
#使用 docker top 来查看容器内部运行的进程
docker top bdf16fc374aa
docker top bold_sanderson
#使用 docker inspect 来查看 Docker 的底层信息
docker inspect bdf16fc374aa

#容器连接
#创建一个新的 Docker 网络
#-d：参数指定 Docker 网络类型，有 bridge、overlay
#其中 overlay 网络类型用于 Swarm mode
docker network create -d bridge test-net
#运行一个容器并连接到新建的 test-net 网络
docker run -itd --name test1 --network test-net ubuntu /bin/bash
#打开新的终端，再运行一个容器并加入到 test-net 网络
docker run -itd --name test2 --network test-net ubuntu /bin/bash
```

```shell
#进入2个容器，安装 ping
apt-get update
apt install iputils-ping
#进入容器 test1, 可以 ping 通 test2
ping test2
#容器 test2
ping test1
```

## dns

可以在宿主机的 /etc/docker/daemon.json 文件中增加以下内容来设置全部容器的 DNS

```json
{
  "dns" : [
    "114.114.114.114",
    "8.8.8.8"
  ]
}
```

设置后，启动容器的 DNS 会自动配置为 114.114.114.114 和 8.8.8.8。

配置完，需要重启 docker 才能生效。

查看容器的 DNS 是否生效可以使用以下命令，它会输出容器的 DNS 信息

```shell
docker run -it --rm  ubuntu  cat etc/resolv.conf
#如果只想在指定的容器设置 DNS，则可以使用以下命令：
docker run -it --rm -h host_ubuntu  --dns=114.114.114.114 --dns-search=test.com ubuntu
```

参数说明：

--rm：容器退出时自动清理容器内部的文件系统。

-h HOSTNAME 或者 --hostname=HOSTNAME： 设定容器的主机名，它会被写到容器内的 /etc/hostname 和 /etc/hosts。

--dns=IP_ADDRESS： 添加 DNS 服务器到容器的 /etc/resolv.conf 中，让容器用这个服务器来解析所有不在 /etc/hosts 中的主机名。

--dns-search=DOMAIN： 设定容器的搜索域，当设定搜索域为 .example.com 时，在搜索一个名为 host 的主机时，DNS 不仅搜索 host，还会搜索 host.example.com。
