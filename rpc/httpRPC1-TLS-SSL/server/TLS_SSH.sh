#!/bin/bash

# 需要在linux上运行,windows下会报错

# 生成私钥
openssl genrsa -out server.key 2048

# 生成自签名的证书
openssl req -new -x509 -key server.key -subj "/CN=localhost" -out server.crt -days 3650

# 只读权限
chmod 400 server.key
