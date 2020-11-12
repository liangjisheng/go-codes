# pprof

运行这个文件你的 HTTP 服务会多出 /debug/pprof 的 endpoint 可用于观察应用程序的情况

1.通过web页面访问

访问 <http://127.0.0.1:6060/debug/pprof/>

2.通过交互式终端使用

```sh
go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=60
go tool pprof http://127.0.0.1:6060/debug/pprof/heap
go tool pprof http://127.0.0.1:6060/debug/pprof/block
go tool pprof http://127.0.0.1:6060/debug/pprof/mutex
```

3.执行test用例

```sh
go test -bench=. -cpuprofile=cpu.prof
```
