# pprof

使用 go tool pprof 命令获取指定的 profile 文件，采集 60s 的 CPU 使用情况
会将采集的数据下载到本地，之后进入交互模式，可以使用命令行查看运行信息

```shell
go tool pprof http://127.0.0.1:8080/debug/pprof/profile -seconds 60
```

使用命令行进入交互式模式查看

```shell
go tool pprof pprof.samples.cpu.001.pb.gz
```

也可以打开浏览器查看 cpu 使用火焰图

```shell
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```
