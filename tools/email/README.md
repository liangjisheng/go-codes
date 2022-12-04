# README

[send email](https://gitee.com/zhucheer/orange/tree/master/mailer)
[send email doc](https://www.kancloud.cn/chase688/orange_framework/1481191)

[golang send email](https://pkg.go.dev/gopkg.in/gomail.v2#example-package)
[gmail 邮箱开启 IMAP&SMTP](https://support.google.com/mail/answer/7126229)

gmail 需要去 [这里](https://support.google.com/accounts/answer/6010255?authuser=1#zippy=%2Cif-less-secure-app-access-is-on-for-your-account) 设置下允许低安全等级的 client 通过 username 和 password 访问邮箱

2022.05.30 之后就不能第三方(代码)通过用户名密码登录 gmail 邮箱发邮件了
[link](https://support.google.com/accounts/answer/6010255?authuser=1#zippy=%2Cif-less-secure-app-access-is-on-for-your-account)

QQ 邮箱
POP3 服务器地址：qq.com（端口：995）
SMTP 服务器地址：smtp.qq.com（端口：465/587）

163 邮箱
POP3 服务器地址：pop.163.com（端口：110）
SMTP 服务器地址：smtp.163.com（端口：25）

126 邮箱
POP3 服务器地址: pop.126.com (端口: 110)
SMTP 服务器地址: smtp.126.com (端口: 25)

gmail
smtp.gmail.com
Requires SSL: Yes
Requires TLS: Yes (if available)
Requires Authentication: Yes
Port for SSL: 465
Port for TLS/STARTTLS: 587

POP3 服务器地址: pop.gmail.com (端口: 995)
SMTP 服务器地址: smtp.126.com (端口: 25)
