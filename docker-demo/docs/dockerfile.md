# docker file

Dockerfile 是一个被用来构建 Docker 镜像的文本文件，该文件中包含了一行行的指令
（Instruction）这些指令对应着修改、安装、构建、操作的命令，每一行指令构建一层（layer）
层层累积，于是有了一个完整的镜像。

```shell
#在 Dockerfile 文件的存放目录下，执行构建动作
$ docker build -t nginx:v3 .
```

上一节中，有提到指令最后一个 . 是上下文路径，那么什么是上下文路径呢？
上下文路径，是指 docker 在构建镜像，有时候想要使用到本机的文件（比如复制）
docker build 命令得知这个路径后，会将路径下的所有内容打包。
由于 docker 的运行模式是 C/S。我们本机是 C，docker 引擎是 S。实际的构建过程是在
docker 引擎下完成的，所以这个时候无法用到我们本机的文件。这就需要把我们本机的指定
目录下的文件一起打包提供给 docker 引擎使用。
如果未说明最后一个参数，那么默认上下文路径就是 Dockerfile 所在的位置。
注意：上下文路径下不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。
当然，我们也可以像编写 .gitignore 一样的语法写一个 .dockerignore, 通过
它可以忽略上传一些不必要的文件给 Docker 引擎。

## from

FROM	指定基础镜像，用于后续的指令构建。
制作镜像必须要先声明一个基础镜像，基于基础镜像，才能在上层做定制化操作。
通过 FROM指令可以指定基础镜像，在 Dockerfile 中，FROM 是必备指令，且必须是第一条指令。
比如，上面编写的 Dockerfile 文件第一行就是 FROM nginx, 表示后续操作都是基于 Ngnix 镜像之上。

## run

RUN 是在 docker build。
作用：为启动的容器指定默认要运行的程序，程序运行结束，容器也就结束。CMD 指令指定的
程序可被 docker run 命令行参数中指定要运行的程序所覆盖。
注意：如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效。

## copy

复制指令，支持从上下文目录中复制文件或者文件夹到容器里的指定路径。

```shell
COPY hom* /mydir/
COPY hom?.txt /mydir/
```

注意： 如果源路径为文件夹，复制的时候不是直接复制该文件夹，而是将文件夹中的内容复制到目标路径。
目标路径 : 复制到容器内的指定路径，无需提前创建好，路径不存在的话，会自动创建。
注意：目标路径可以是容器内的绝对路径，也可以是相对于工作目录的相对路径（工作目录可以用 WORKDIR 指令来指定）。

## add

ADD 指令与 COPY 指令功能类似，都可以复制文件或文件夹（同样的需求下，官方推荐使用 COPY 指令）

## cmd

CMD 指令用于启动容器时，指定需要运行的程序以及参数。
CMD 类似于 RUN 指令，用于运行程序，但二者运行的时间点不同:
CMD 在docker run 时运行。
RUN 是在 docker build 时运行；
如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效

```shell
CMD ["<可执行文件>", "<参数1>", "<参数2>", ...]
CMD [ "sh", "-c", "echo $HOME" ]
```

## entrypoint

ENTRYPOINT 的功能和 CMD 一样，都用于指定容器启动程序以及参数
如果 Dockerfile 中存在多个 ENTRYPOINT 指令，仅最后一个生效。

对于 CMD 指令， 执行 docker run 命令如果有传递参数，这些参数是可以覆盖
Dockerfile 中的CMD 指令参数，但是对于 ENTRYPOINT 来说，这些参数会被
传递给 ENTRYPOINT，而不是覆盖。

```shell
ENTRYPOINT ["<executeable>","<param1>","<param2>",...]
```

假设我们需要一个打印当前公网 IP 的镜像，那么可以先用 CMD 来实现, Dockerfile 文件如下

```dockerfile
FROM ubuntu:18.04
RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*
CMD [ "curl", "-s", "http://myip.ipip.net" ]
```

```shell
#构建镜像
docker build -t myip .
#启动容器，打印 ip
docker run myip
```

如果我们希望显示 HTTP 头信息，就需要加上 -i 参数。那么我们可以直接加 -i 参数给 
docker run myip 么？这个是不行的，上面已经说到，docker run 传递的参数会替换 
CMD 的默认值 ，因此这里的 -i 替换了原来的 CMD，而不是添加在原来的 
curl -s http://myip.ipip.net 后面， 而 -i 根本不是命令，所以自然找不到报错。

那么如果希望支持 -i 这参数，我们就必须重新完整的输入这个命令：

```shell
docker run myip curl -s http://myip.ipip.net -i
```

这显然不是一个优雅的解决方案，这个时候 ENTRYPOINT 就上场了, 因为它可以传递参数，
而不是覆盖。现在我们重新用 ENTRYPOINT 来实现这个镜像：

```dockerfile
FROM ubuntu:18.04
RUN apt-get update \
    && apt-get install -y curl \
    && rm -rf /var/lib/apt/lists/*
ENTRYPOINT [ "curl", "-s", "http://myip.ipip.net" ]
```

这次我们再来尝试, 可以看到，这次成功了

```shell
docker run myip -i
```

## env

ENV 设置环境变量，定义了环境变量，那么在后续的指令中，就可以使用这个环境变量

```shell
ENV <key> <value>
ENV <key1>=<value1> <key2>=<value2>...
```

```dockerfile
ENV NODE_VERSION 7.2.0

RUN curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.xz" \
  && curl -SLO "https://nodejs.org/dist/v$NODE_VERSION/SHASUMS256.txt.asc"
```

## arg

ARG 指令用于指定构建参数，与 ENV 功能一样，都是设置环境变量。不同点在于作用域不一样, 
ARG 声明的环境变量仅对 Dockerfile 内有效，也就是说仅对 docker build 的时候有效，
将来容器运行的时候不会存在这些环境变量的。

```shell
ARG <参数名>[=<默认值>]
```

不要使用 ARG 保存密码之类的敏感信息，因为通过 docker history 可以看到所有数据
ARG 指令有生效范围，如果在 FROM 指令之前指定，那么只能用于 FROM 指令中。

```dockerfile
ARG DOCKER_USERNAME=library

FROM ${DOCKER_USERNAME}/alpine

RUN set -x ; echo ${DOCKER_USERNAME}
```

使用上述 Dockerfile 会发现无法输出 ${DOCKER_USERNAME} 变量的值，要想正常输出，
你必须在 FROM 之后再次指定 ARG ：

```dockerfile
# 只在 FROM 中生效
ARG DOCKER_USERNAME=library

FROM ${DOCKER_USERNAME}/alpine

# 要想在 FROM 之后使用，必须再次指定
ARG DOCKER_USERNAME=library

RUN set -x ; echo ${DOCKER_USERNAME}
```

对于多阶段构建，尤其要注意这个问题:

```dockerfile
# 这个变量在每个 FROM 中都生效
ARG DOCKER_USERNAME=library

FROM ${DOCKER_USERNAME}/alpine

RUN set -x ; echo 1

FROM ${DOCKER_USERNAME}/alpine

RUN set -x ; echo 2
```

对于上述 Dockerfile 两个 FROM 指令都可以使用 ${DOCKER_USERNAME}，
对于在各个阶段中使用的变量都必须在每个阶段分别指定：

```dockerfile
ARG DOCKER_USERNAME=library

FROM ${DOCKER_USERNAME}/alpine

# 在FROM 之后使用变量，必须在每个阶段分别指定
ARG DOCKER_USERNAME=library

RUN set -x ; echo ${DOCKER_USERNAME}

FROM ${DOCKER_USERNAME}/alpine

# 在FROM 之后使用变量，必须在每个阶段分别指定
ARG DOCKER_USERNAME=library

RUN set -x ; echo ${DOCKER_USERNAME}
```

## expose

EXPOSE 指令用于暴露容器运行时提供服务的端口。注意，这仅仅是一个声明，容器实际运行时，
并不会开启这个声明的端口。

在 Dockerfile 中写入 EXPOSE 指令好处如下：
1、帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射；
2、在运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口

docker run -p，是映射宿主端口和容器端口，换句话说，就是将容器的对应端口服务公开给外界访问，
而 EXPOSE 仅仅是声明容器打算使用什么端口而已，并不会自动在宿主进行端口映射。

## workdir

WORKDIR 指令用于指定工作目录，后面各层的当前目录即为 WORKDIR 指定的目录，如果该目录
不存在，WORKDIR 会自动建立目录。

```shell
WORKDIR <工作目录路径>
```

```dockerfile
RUN cd /app
RUN echo "hello" > world.txt
```

通过这个 Dockerfile 构建镜像运行后，会发现找不到 /app/world.txt 文件，或者内容不是 hello
为什么会发生这种情况呢？因为在 Shell 中，连续两行命令执行是在同一个进程中；而在 docker build 
构建镜像过程中，每一个 RUN 命令都会新建一层，执行环境是完全不同的。

定义了 WORKDIR /app 后，就可以指定后面每层的工作目录了

## user

USER 用于指定执行后续命令的用户和用户组，这边只是切换后续命令执行的用户
（用户和用户组必须提前已经存在）

```shell
USER <用户名>[:<用户组>]
```

```dockerfile
RUN groupadd -r redis && useradd -r -g redis redis
USER redis
RUN [ "redis-server" ]
```

## health check

EALTHCHECK 指令用于设置 Docker 要如何判断容器状态是否正常

```shell
HEALTHCHECK [选项] CMD <命令> #设置检查容器健康状况的命令
HEALTHCHECK NONE #如果基础镜像有健康检查指令，使用这行可以屏蔽掉其健康检查指令
```

当一个镜像指定了 HEALTHCHECK 后，初启动容器时，初始状态为 starting, 当 HEALTHCHECK 
指令检查成功后，容器状态会变为 healthy，如果重试多次失败，则会变为 unhealthy。

--interval=<间隔>：两次健康检查的间隔，默认为 30 秒；
--timeout=<时长>：健康检查命令运行超时时间，如果超过这个时间，本次健康检查就被视为失败，默认 30 秒；
--retries=<次数>：健康检查重试次数，若指定次数依然失败，则将容器状态视为 unhealthy，默认 3 次。

```dockerfile
FROM nginx
RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*
HEALTHCHECK --interval=5s --timeout=3s \
CMD curl -fs http://localhost/ || exit 1
```

## on build

ONBUILD 用于延迟构建命令的执行。简单的说，就是 Dockerfile 里用 ONBUILD 指定的命令，
在本次构建镜像的过程中不会执行（假设镜像为 test-build）。当有新的 Dockerfile 使用了
之前构建的镜像 FROM test-build ，这时执行新镜像的 Dockerfile 构建时候，会执行 
test-build 的 Dockerfile 里的 ONBUILD 指定的命令。

```dockerfile
FROM node:slim
RUN mkdir /app
WORKDIR /app
ONBUILD COPY ./package.json /app
ONBUILD RUN [ "npm", "install" ]
ONBUILD COPY . /app/
CMD [ "npm", "start" ]
```

## label

MAINTAINER	指定Dockerfile的作者/维护者。（已弃用，推荐使用LABEL指令）
LABEL 指令用来给镜像添加一些元数据（metadata），以键值对的形式

```shell
LABEL <key>=<value> <key>=<value> <key>=<value> ...
```

```dockerfile
LABEL org.opencontainers.image.authors="alice"
LABEL org.opencontainers.image.documentation="https://www.alice.com"
```

## shell

SHELL 指令可以指定 RUN 、ENTRYPOINT、CMD 指令执行的 shell 命令
Linux 中默认为 ["/bin/sh", "-c"]

```dockerfile
SHELL ["/bin/sh", "-cex"]

# /bin/sh -cex "nginx"
CMD nginx
```

## scratch

特殊的镜像：scratch, 它表示一个空白的镜像
以 scratch 为基础镜像，表示你不以任何镜像为基础。
