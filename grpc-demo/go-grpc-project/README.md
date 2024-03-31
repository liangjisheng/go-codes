# go grpc project

整理自[Kevin Vogel](https://medium.com/gitconnected/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30)的文章。

[part-1](https://alanhou.org/microservices-with-go-grpc-api-gateway-and-authentication-part-1/)
[part-2](https://alanhou.org/microservices-with-go-grpc-api-gateway-and-authentication-part-2/)

## 项目结构
* [网关go-grpc-api-gateway](./go-grpc-api-gateway/)
* [认证微服务go-grpc-auth-svc](./go-grpc-auth-svc/)
* [商品微服务go-grpc-product-svc](./go-grpc-product-svc/)
* [订单微服务go-grpc-order-svc](./go-grpc-order-svc/)

注册新用户

```shell
curl --request POST \
  --url http://localhost:3000/auth/register \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "elon@musk.com",
 "password": "1234567"
}'
```

用户登录

```shell
curl --request POST \
  --url http://localhost:3000/auth/login \
  --header 'Content-Type: application/json' \
  --data '{
 "email": "elon@musk.com",
 "password": "1234567"
}'
```

```json
{"status":200,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDIxMjgzMjIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.WdBxR11kirOwnghbyqihe6k3w7goa4RLEtkq9tsR3YY"}
```

创建商品

```shell
curl --request POST \
  --url http://localhost:3000/product/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDIxMjgzMjIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.WdBxR11kirOwnghbyqihe6k3w7goa4RLEtkq9tsR3YY' \
  --header 'Content-Type: application/json' \
  --data '{ "name": "Product A", "stock": 5, "price": 15 }'
```

查找商品

```shell
curl --request GET \
  --url http://localhost:3000/product/1 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDIxMjgzMjIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.WdBxR11kirOwnghbyqihe6k3w7goa4RLEtkq9tsR3YY'
```

创建订单

```shell
curl --request POST \
  --url http://localhost:3000/order/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDIxMjgzMjIsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MSwiRW1haWwiOiJlbG9uQG11c2suY29tIn0.WdBxR11kirOwnghbyqihe6k3w7goa4RLEtkq9tsR3YY' \
  --header 'Content-Type: application/json' \
  --data '{
 "productId": 1,
 "quantity": 1
}'
```
