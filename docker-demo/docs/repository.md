# repository

目前 Docker 官方维护了一个公共仓库 Docker Hub。

大部分需求都可以通过在 Docker Hub 中直接下载镜像来实现。

注册
在 https://hub.docker.com 免费注册一个 Docker 账号

```shell
docker login
docker logout
docker pull ubuntu:18.04 
docker tag ubuntu:18.04 liangjisheng/ubuntu:18.04
docker push liangjisheng/ubuntu:18.04
docker search liangjisheng/ubuntu:18.04
```