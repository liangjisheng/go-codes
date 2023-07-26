# apt install

<https://www.jianshu.com/p/35e7af7db96a>  
<https://www.cnblogs.com/xiaohuomiao/p/10601760.html>  
<https://blog.csdn.net/qq_34680444/article/details/86238516>

install

```shell
#可以直接执行, 如果没有找到源的话再执行下面的语句更新源
sudo apt-get install mysql-server

sudo wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb
sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb
sudo apt-get update
```

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

退出后重新登录

```shell
mysql -u root -p
```

## 改配置

修改 /etc/mysql/mysql.conf.d/mysqld.cnf, 改为监听 *

配置文件中添加 skip-grant-tables 这一样可以避免密码校验

```conf
bind-address            = *
mysqlx-bind-address     = *
```

重启

```shell
sudo systemctl status mysql
sudo systemctl stop mysql
sudo systemctl start mysql
sudo systemctl restart mysql
```
