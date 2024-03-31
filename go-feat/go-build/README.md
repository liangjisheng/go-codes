# README

go build -X 设置包中的变量值, 用来设置版本、时间、日期等等

```shell
#等号之间不能有空格
go build -v -ldflags "-X 'main.Version=v1.2.0-53-ga9e0819' -X 'main.BuildUser=$(id -u -n)'"
#将当前 build 用户注入到包变量 gofeat/go-build/app.BuildUser 中去
go build -v -ldflags "-X 'gofeat/go-build/app.BuildUser=$(id -u -n)' -X 'gofeat/go-build/app.BuildTime=$(date)'"
```

go build -v -ldflags "-s -w" -o xxx
-v 编译时显示包名
-p n 开启并发编译，默认情况下该值为 CPU 逻辑核数
-a 强制重新构建
-n 打印编译时会用到的所有命令，但不真正执行
-x 打印编译时会用到的所有命令
-race 开启竞态检测

-s 的作用是去掉符号信息
-w 的作用是去掉调试信息
