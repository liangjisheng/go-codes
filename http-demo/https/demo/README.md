# https

[study](https://studygolang.com/articles/2946)
[github](https://github.com/bigwhite/experiments/gohttps)

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

## demo1

[基于x509的认证授权技术](https://islishude.github.io/blog/2020/09/22/crypto/%E5%9F%BA%E4%BA%8Ex509%E7%9A%84%E8%AE%A4%E8%AF%81%E6%8E%88%E6%9D%83%E6%8A%80%E6%9C%AF/)

其实就是服务端也要认证客户端的证书, [参考代码](https://github.com/islishude/grpc-mtls-example)
