# go test

[article](https://geektutu.com/post/quick-go-test.html)

## 生成测试模板

可以使用工具 [cweill/gotests](https://github.com/cweill/gotests) 生成测试模板，再填充测试用例

```shell
go get -u github.com/cweill/gotests/...
gotests -all -w .
gotests -only='HelloWorld' -w .
```

## 测试并生成覆盖率数据

```shell
go test -race -cover -coverprofile=./coverage.out -timeout=10m -short -v .
go test -race -cover -coverprofile=./coverage.out -timeout=10m -short -v calc.go calc_test.go
```

覆盖率分析

```shell
go tool cover -func ./coverage.out
#生成 HTML 文件
go tool cover -html=coverage.out -o coverage.html
```

## Example

Godoc 将在函数文档旁提供其示例。示例测试函数名称必须以 Example 开头，无参数，无返回值。

函数的结尾可能包含以 Output: 或 Unordered output: 开头的注释。Unordered output: 开头的注释会忽略输出行的顺序

参考 [errors/example_test.go](https://github.com/marmotedu/errors/blob/v1.0.2/example_test.go)

在执行 go test 时会自动执行这些测试，将示例测试输出到标准输出的内容与注释作对比（忽略行前后的空格）。相等通过，否则不通过

## TestMain

测试用例在执行时会先执行 TestMain 函数，可以在 TestMain 中调用 m.Run() 函数执行普通的测试函数。

参考 [iam/user_test.go](https://github.com/marmotedu/iam/blob/v1.0.8/internal/apiserver/service/v1/user_test.go)

可在 m.Run() 函数前面编写准备逻辑，在 m.Run() 后面编写清理逻辑

```txt
Before all tests
=== RUN   ExampleHello
--- PASS: ExampleHello (0.00s)
PASS
After all tests
```

## Fake 测试

对于比较复杂的接口，可以 Fake 一个接口实现来进行测试。即针对接口实现假的实例。Fake 实例需要根据业务自行实现。

参考 [iam/internal/apiserver/store/fake](https://github.com/marmotedu/iam/tree/v1.0.8/internal/apiserver/store/fake) 

通过 TestMain 初始化 fake 实例 [iam/store.go](https://github.com/marmotedu/iam/blob/v1.0.8/internal/apiserver/store/store.go#L12-L17)

## 性能测试

安装

```shell
git clone https://github.com/wg/wrk

cd wrk
make
sudo cp ./wrk /usr/local/bin
```

基本使用：线程数为 CPU 核数 2~4 倍即可，避免切换过多造成效率降低

```shell
# 指定线程数、并发数、持续时间、超时时间、打印延迟，也可以指定 Lua 脚本实现复杂的请求。
wrk -t144 -c30000 -d30s -T30s --latency http://127.0.0.1:8080/index
```

API 性能参考

| 指标      | 要求                              |
|----------|---------------------------------|
| 响应时间   | <500ms，否则需要优化                   |
| 请求成功率 | 99.95%                          |
| QPS      | 满足预期的情况下，服务器状态稳定，单机 QPS 在 1000+ |
