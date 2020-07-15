#!/bin/bash

# 生成ECC私钥
openssl ecparam -genkey -name secp384r1 -out server.key

# 生成公钥 openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密
# -key指定私钥文件、-x509指输出证书、-days 3650为有效期
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
