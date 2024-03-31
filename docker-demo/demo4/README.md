# README

[docker go](https://mp.weixin.qq.com/s/773INmwebAIy6zDtGHOEoQ)
[multi stage build](https://zhuanlan.zhihu.com/p/535414655)
[multi stage build](https://blog.csdn.net/MyySophia/article/details/121138073)

docker build -f Dockerfile1 -t test-go-docker:latest .

docker build -f Dockerfile2 -t test-go-docker2:latest .

对于最后一个 Dockerfile，我们只将 alpine 基础镜像更改为 scratch. Scratch 是一个空镜像，
所以一旦容器运行，我们就无法执行到容器中，因为它没有 shell 命令
docker build -f Dockerfile3 -t test-go-docker3:latest .
