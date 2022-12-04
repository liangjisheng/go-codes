# README

http 服务测试

```shell
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'
grpcurl -plaintext -d '{"name":"alice"}' -import-path ./helloworld -proto hello_world.proto localhost:50051 proto.Greeter/SayHello
```
