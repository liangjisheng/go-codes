# README

[go-sso](https://blog.csdn.net/guyan0319/article/details/106697940)
[go-sso](https://github.com/guyan0319/go-sso)

## 注册手机号

```shell
curl -i -X POST "http://127.0.0.1:8080/signup/mobile" -d '{"mobile":"13522227564", "passwd":"123456", "code": "111111"}'
```

## 手机号+密码 登录

```shell
curl -i -X POST "http://127.0.0.1:8080/login" -d '{"mobile":"13522227564", "passwd":"123456"}'
```

## 手机号+验证码登录

```shell
curl -i -X POST "http://127.0.0.1:8080/login/mobile" -d '{"mobile":"13522227564","code":"123456"}'
```

## 获取用户信息

使用上面登录获取的Access_Token获取用户信息

```shell
curl -i "http://127.0.0.1:8080/my/info" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYzNjEwNjMyOH0.n59WjHkYASfdv9IhrH1ngTN6ECBMi6o1t05NrzgHnxw"
```

## 查看手机号是否存在

```shell
curl -i -X POST "http://127.0.0.1:8080/signup/mobile/exist" -d '{"mobile":"13522227564"}'
```

## access token 续期

```shell
curl -i -X POST "http://127.0.0.1:8080/renewal" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYzNjEwNjMyOH0.n59WjHkYASfdv9IhrH1ngTN6ECBMi6o1t05NrzgHnxw;Refresh_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYzNjEwODEyOH0.35GZNL-H0Ky5KwDM3JadmkXRFNtLP4ltR34mg7nMkno"
```

## 登出系统

```shell
curl -i -X POST "http://127.0.0.1:8080/logout" -b "Access_Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYzNjEwNjMyOH0.n59WjHkYASfdv9IhrH1ngTN6ECBMi6o1t05NrzgHnxw"
```

## 小结

仅使用jwt实现单点登录会遇到两个问题

- 用户无法主动登出，即服务端发出token后，无法主动销毁token，用户还可以用通过token访问系统，本项目增加了缓存登出用户token到黑名单的方式，变相实现登出。
- token续期问题，access_token携带有效期，有效期过了无法自动续期。本项目提供了续期接口（renewal），服务端在生成access_token同时还会生成refresh_token（有效期比access_token长），用户可以通过有效的refresh_token和access_token访问renewal接口重新获取新的refresh_token和access_token。
