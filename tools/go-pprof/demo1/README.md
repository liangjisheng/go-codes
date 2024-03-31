# pprof

运行这个文件你的 HTTP 服务会多出 /debug/pprof 的 endpoint 可用于观察应用程序的情况

通过web页面访问

访问 <http://127.0.0.1:8080/debug/pprof/>

通过交互式终端使用

```sh
go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=60
go tool pprof http://127.0.0.1:8080/debug/pprof/heap
go tool pprof http://127.0.0.1:8080/debug/pprof/block
go tool pprof http://127.0.0.1:8080/debug/pprof/mutex
```

## 采集数据

使用 Go 命令行工具生成性能数据文件

- data.test 测试生成的二进制文件，进行性能分析时用于解析各种符号
- cpu.profile，CPU 性能数据文件
- mem.profile，内存性能数据文件

```shell
go test -benchtime=10s -benchmem -bench=".*" -cpuprofile cpu.profile -memprofile mem.profile
```

生成采用数据后，可按以下思路分析性能

- 采样图：矩形面积最大
- 火焰图：格子最宽。
- go tool pprof：cum% 最大

## 分析采样图

以 CPU 分析为例。Go 运行时默认以 100 Hz 的频率对 CPU 使用情况采样，即每秒采样 100 次、每 10 毫秒采样一次
每次采样时记录正在运行的函数，并统计其运行时间，生成 CPU 性能数据（cpu.profile）

安装 graphviz, 见上面

### 生成调用图

```shell
# svg 格式
go tool pprof -svg cpu.profile > cpu.svg
# pdf 格式
go tool pprof -pdf cpu.profile > cpu.pdf
# png 格式
go tool pprof -png cpu.profile > cpu.png
```

### 分析火焰图

火焰图可把采样到的堆栈轨迹（Stack Trace）转化为直观图片显示。

使用 pprof 工具打开数据文件，可在浏览器中直观查看数据

```shell
go tool pprof -http="0.0.0.0:8081" cpu.profile
```

### 分析数据

交互式查看 CPU 性能数据文件

- File 二进制可执行文件名称
- Type 采样文件的类型，例如 cpu、mem 等
- Time 生成采样文件的时间
- Duration 程序执行时间。程序在采样时，会自动分配采样任务给多个核心，总采样时间可能会大于总执行时间
- (pprof), 命令行提示，表示当前在 go tool 的 pprof 工具命令行中 (还包括 cgo、doc、pprof、trace 等)

```shell
go tool pprof data.test cpu.profile
```
