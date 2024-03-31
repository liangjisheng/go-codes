#!/bin/bash

# 使用-k, 是不对服务器的证书进行检查，这样就不必关心服务器证书的导出问题了
# --cert ../tls/server.pem --cert-type PEM --key ../tls/server.key
curl -k -X "POST" https://localhost:8080/v1/example/route -d '{"some_integer":99, "some_float":1}' -H "Content-Type: application/json" -H "Authorization: Bearer grpc.auth.token"

curl -k -X "POST" https://localhost:8080/v1/example/route -d '{"some_integer":99, "some_float":1}' -H "Content-Type: application/json" -H "Authorization: Bearer grpc.auth.token"
