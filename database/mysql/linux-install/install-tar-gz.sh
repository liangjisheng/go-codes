#!/bin/bash

wget https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz
# wget https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz
tar xf mysql-8.0.19-linux-glibc2.12-x86_64.tar.xz
mv mysql-8.0.19-linux-glibc2.12-x86_64 /usr/local/mysql
mkdir /usr/local/mysql/data /usr/local/mysql/log
touch /usr/local/mysql/log/error.log
vim /etc/my.cnf
/usr/local/mysql/bin/mysqld --initialize
# root@localhost: 3SofffQat=m-
chown -R mysql:mysql /usr/local/mysql
/usr/local/mysql/support-files/mysql.server start
# /usr/local/mysql/support-files/mysql.server restart

# 输入初始化生成的密码, 进来后首先修改密码,不然做不了其他的操作
mysql -u root -p
# mysql5.7修改密码
set password for root@localhost = password('newpass');
# mysql8.0修改密码
alter user 'root'@'localhost' identified by 'password';

# 开放远程连接
use mysql
# ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
update user set host='%' where user='root';
flush privileges;
grant all privileges on *.* to 'root'@'%';
flush privileges;

# The server quit without updating PID file (/var/run/mysqld/mysqld.pid)
# 这个可能是因为初始化后的文件不属于 mysql 这个用户组

# 将mysql作为系统服务, 好像起不了作用
ln -s /usr/local/mysql/support-files/mysql.server /etc/init.d/mysql
systemctl daemon-reload
systemctl stauts mysql
