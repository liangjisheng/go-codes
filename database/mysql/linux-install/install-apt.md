# apt install

<https://dev.mysql.com/doc/mysql-apt-repo-quick-guide/en/>
<https://www.jianshu.com/p/35e7af7db96a>  
<https://www.cnblogs.com/xiaohuomiao/p/10601760.html>  
<https://blog.csdn.net/qq_34680444/article/details/86238516>

install

```shell
#增加源，有可能不是 ubuntu 18.04 bionic, 那就通过下面的方式增加源
sudo wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb
sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb
sudo apt-get update

#另一种增加源的方式 ubuntu 18.04 bionic
echo "deb https://repo.mysql.com/apt/ubuntu/ bionic mysql-apt-config
deb https://repo.mysql.com/apt/ubuntu/ bionic mysql-8.0
deb https://repo.mysql.com/apt/ubuntu/ bionic mysql-tools
deb https://repo.mysql.com/apt/ubuntu/ bionic mysql-cluster-8.0
deb-src https://repo.mysql.com/apt/ubuntu/ bionic mysql-8.0" | sudo tee /etc/apt/sources.list.d/mysql.list

#必须操作
sudo apt-get update

# 如果报签名失败, 可以增加公钥，然后再执行 update
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 8C718D3B5072E1F5
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys A8D3785C

#install mysql
sudo apt-get install mysql-server

#切到 root 用户下
sudo su -
# 直接执行 mysql 登录
msyql
```

## 改配置

修改 /etc/mysql/mysql.conf.d/mysqld.cnf, 改为监听 *

配置文件中添加 skip-grant-tables 这一样可以避免密码校验

```conf
bind-address            = *
mysqlx-bind-address     = *
max_connections         = 1000
```

重启

```shell
sudo systemctl status mysql
sudo systemctl stop mysql
sudo systemctl start mysql
sudo systemctl restart mysql
```

create user (root 用户登录)

```sql
use mysql
select user,plugin,host from user;

create user user1@'%' identified by 'pass';
create database user1_db;
grant all privileges on user1_db.* to user1;
flush privileges;
show grants for user1;
```

使用新用户连接 db, 这样就安装完 mysql 了。

将 db 目录移动到外挂磁盘上

```shell
mkdir -p /data/mysql
sudo su -
cp -r /var/lib/mysql/* /data/mysql
mv /var/lib/mysql /var/lib/mysql.bak
ln -s /data/mysql /var/lib/mysql
chown -R mysql. /data/mysql
vim /etc/apparmor.d/usr.sbin.mysqld
```

找到下面两行

```conf
# Allow data dir access
/var/lib/mysql/ r,
/var/lib/mysql/** rwk,
```

在下面增加

```conf
/data/mysql/ r,
/data/mysql/** rwk,
```

修改权限配置后重启 apparmor,mysql

```shell
systemctl restart apparmor
systemctl start mysql
```

使用 navicat 连接 mysql, 如果报下面的错误

```text
2059 - Authentication plugin 'caching_sha2_password' cannot be loaded: dlopen(../Frameworks/caching_sha2_password.so, 2): image not found
```

登录 mysql, 修改用户的密码验证方式，再次连接就可以登录了

```sql
alter user 'username'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
```

## other

切换到 root 用户下 sudo su -, 直接执行 mysql 进入，另一种方式是下载安装MySQL的过程中，系统会自动给我们创建一个用户

```shell
#里面包含 user,password
sudo cat /etc/mysql/debian.cnf
```

```sql
-- 使用mysql数据库
use mysql
select user,plugin,host from user;

-- 修改root密码
-- mysql5.x 修改密码, 下面2句都可以
update user set authentication_string=password('123456'), plugin='mysql_native_password' where user='root';
-- set password for 'root'@'localhost' = password('newpass');
-- mysql8.x 修改密码
-- root 默认密码格式为 auth_socket, 修改其为 mysql_native_password
update user set authentication_string='', plugin='mysql_native_password' where user='root';
-- 如果这句执行失败的话，可以退出重新登录后执行
alter user 'root'@'localhost' identified with mysql_native_password by 'password';
    
-- 开放远程访问权限(授权远程连接)
update user set host='%' where user='root';

-- 刷新权限
flush privileges;
```

apt-key

```shell
#删除对应的公钥
sudo apt-key del A4A9406876FCBD3C456770C88C718D3B5072E1F5
sudo apt-key del BCA43417C3B485DD128EC6D4B7B3B788A8D3785C
sudo apt-key del 859BE8D7C586F538430B19C2467B942D3A79BD29
#查看公钥
sudo apt-key list
```
