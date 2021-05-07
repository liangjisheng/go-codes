# 调用流程

1. 启动 hellogrpc  
2. 启动 grpcgateway
3. 调用接口

```sh
curl -X POST -d '{"name":"ljs"}' "http://127.0.0.1:8080/v1/example/echo"
```

