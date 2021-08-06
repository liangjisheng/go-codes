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
# 修改root密码
# mysql5.x 修改密码
set password for 'root'@'localhost' = password('newpass');
# mysql8.x 修改密码
alter user 'root'@'localhost' identified by 'password';

# 开放远程访问权限(授权远程连接)
update user set host='%' where user='root';
grant all privileges on *.* to 'root'@'%';

# 执行刷新权限
flush privileges;
