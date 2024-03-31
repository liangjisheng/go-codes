# README

[article](https://mp.weixin.qq.com/s/BYkOQZO73UTRAUiONnUrow)

Go 1.18 二进制文件的信息嵌入

编译 main.go 得到二进制文件之后，就可以通过以下命令获取到编译该文件时的一些元信息

```shell
go build -o main1.18
go version -m main1.18
```

如果在构建过程中，为二进制文件指定了 -tags 等，我们同样能够以上的方式获取

```shell
go build -tags version1.1 --ldflags="-X 'main.s=sss'" -o main1.18
go version -m main1.18
```
