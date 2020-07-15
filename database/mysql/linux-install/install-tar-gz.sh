#!/bin/bash

# https://www.jianshu.com/p/276d59cbc529

wget https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.19-linux-glibc2.12-x86_64.tar.xz
# wget https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-8.0.19-linux-glibc2.12-x86_64.tar.xz
tar xvf mysql-8.0.19-linux-glibc2.12-x86_64.tar.xz
mv mysql-8.0.19-linux-glibc2.12-x86_64 /usr/local/mysql
cd /usr/local/mysql
mkdir data
cd ..
chown -R mysql:mysql /usr/local/mysql
chmod -R 755 /usr/local/mysql
cd /usr/local/mysql/bin
./mysqld --initialize --user=mysql --datadir=/usr/local/mysql/data --basedir=/usr/local/mysql
vim /etc/my.cnf

/usr/local/mysql/support-files/mysql.server start

# 输入初始化生成的密码
mysql -u root -p

# mysql5.7修改密码
set password for root@localhost = password('newpass');
# mysql8.0修改密码
alter user 'root'@'localhost' IDENTIFIED BY 'password';

# 开放远程连接
use mysql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
update user set host='%' where user='root';
grant all privileges on *.* to 'root'@'%';
flush privileges;


# 将mysql作为系统服务
ln -s /usr/local/mysql/support-files/mysql.server /etc/init.d/mysql
ln -s /usr/local/mysql/bin/mysql /usr/bin/mysql

sudo systemctl daemon-reload
