# bimg

[github](https://github.com/h2non/bimg)

先安装 libvips，再获取 h2non/bimg 包

mac install

```shell
$ brew install vips
```

再次尝试获取 h2non/bimg 包，通常会提示 invalid flag in pkg-config --cflags: -Xpreprocessor ;
此时 CGO_FLAGS 授权，即执行一下命令：

```shell
$ export CGO_CFLAGS_ALLOW=-Xpreprocessor
```
