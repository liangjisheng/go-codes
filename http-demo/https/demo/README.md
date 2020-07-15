# https

[study](https://studygolang.com/articles/2946)
[github](https://github.com/bigwhite/experiments/tree/master/gohttps)

使用curl访问

```sh
curl -k https://localhost:8080/
```

使用 curl 如果不加-k, curl会验证不通过

客户端访问,忽略对服务端证书的校验

```sh
cd client-ignore-verify-server-cert
go run client.go
```

或者服务启动后在浏览器中输入 <https://localhost:8080/> 访问,受信任服务端证书
