#!/bin/bash

# https://www.jianshu.com/p/35e7af7db96a
# https://www.cnblogs.com/xiaohuomiao/p/10601760.html
# https://blog.csdn.net/qq_34680444/article/details/86238516

sudo wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb
sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb
sudo apt-get update
sudo apt-get install mysql-server

# 开放远程访问
# 使用mysql数据库
use mysql
# mysql8.0需要更改加密方式
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';

# 开放远程访问权限(授权远程连接)
update user set host='%' where user='root';
grant all privileges on *.* to 'root'@'%';

# 执行刷新权限
flush privileges;
