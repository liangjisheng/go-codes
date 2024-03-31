# mkcert

mkcert 工具 [release](https://github.com/FiloSottile/mkcert/releases/latest)

[post](https://islishude.github.io/blog/2019/01/30/blockchain/%E7%BB%99-localhost-%E7%AD%BE%E5%8F%91-https-%E8%AF%81%E4%B9%A6/)

安装完后启动服务

```shell
$ node node_https.js
#或者
$ go run main_https.go
```

测试

```shell
$ curl https://localhost:8000
# hello,world
$ curl https://127.0.0.1:8000
# hello,world
```
