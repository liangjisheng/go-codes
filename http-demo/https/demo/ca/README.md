# CA

首先我们来建立我们自己的CA，需要生成一个CA私钥和一个CA的数字证

```sh
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 5000 -out ca.crt
```

接下来，生成server端的私钥，生成数字证书请求，并用我们的ca私钥签发server的数字证书

```sh
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
```

生成客户端的私钥与证书

```sh
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=localhost" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
```
