# mysql cluster install

[article](https://database.51cto.com/art/201505/475376_all.htm)
[多管理节点](http://www.jizhuomi.com/software/527.html)

## 概念

1. sql节点(mysqld): 分布式数据库, 包括自身数据和查询中心结点数据
2. 数据结点(Data node – ndbd): 集群共享数据(内存中)
3. 管理服务器(Management Server – ndb_mgmd): 管理集群 SQL node,Data node

不管是 management server, 还是Data node、SQL node, 都需要先安装MySQL集群版本
然后根据不用的配置来决定当前服务器有哪几个角色

## 所有节点

linux 新建mysql 组和用户,管理节点不需要

```shell
groupadd mysql
useradd -g mysql mysql
```

下载mysql集群文件

```shell
wget https://dev.mysql.com/get/Downloads/MySQL-Cluster-8.0/mysql-cluster-8.0.26-linux-glibc2.12-x86_64.tar.gz
tar zxf mysql-cluster-8.0.26-linux-glibc2.12-x86_64.tar.gz
mv mysql-cluster-8.0.26-linux-glibc2.12-x86_64 /usr/local/mysql
mkdir /usr/local/mysql/data /usr/local/mysql/log /usr/local/mysql/sock
```

如果防火墙开着, 需要关闭防火墙, 或者打开 1186,3306 这2个端口

## 配置管理节点

```shell
mkdir /var/lib/mysql-cluster
# 管理节点
vim /var/lib/mysql-cluster/config.ini
cp /usr/local/mysql/bin/ndb_mgm* /usr/local/bin
chmod +x /usr/local/bin/ndb_mgm*
```

## 配置数据节点

```shell
vim /etc/my.cnf
# 初始化(创建系统数据库), 然后会得到root的默认初始密码,记录下来
/usr/local/mysql/bin/mysqld --initialize
chown -R mysql:mysql /usr/local/mysql
```

## 配置sql节点

```shell
vim /etc/my.cnf
# 初始化(创建系统数据库), 然后会得到root的默认初始密码,记录下来
/usr/local/mysql/bin/mysqld --initialize
chown -R mysql:mysql /usr/local/mysql
```

## cluster环境启动

先是管理节点,然后是数据节点,最后是sql节点

启动管理节点

```shell
ndb_mgmd --config-file=/var/lib/mysql-cluster/config.ini --configdir=/var/lib/mysql-cluster
# 使用客户端登录
ndb_mgm
```

启动数据节点  
首次启动需要添加--initial参数, 以便进行NDB节点的初始化工作, 在以后的启动过程中  
则是不能添加该参数的, 否则ndbd程序会清除在之前建立的所有用于恢复的数据文件和日志文件

```shell
# 首次启动
/usr/local/mysql/bin/ndbd --initial
# 非首次启动
/usr/local/mysql/bin/ndbd
```

启动sql节点

```shell
/usr/local/mysql/support-files/mysql.server start
# 使用root默认密码登录,登录后先修改root密码
/usr/local/mysql/bin/mysql -u root –p
# 修改root密码
# mysql5.x 修改密码
set password for 'root'@'localhost' = password('newpass');
# mysql8.x 修改密码
alter user 'root'@'localhost' identified by 'password';
```

## 集群测试

到这里一个简单版本的集群就已经搭建好了,下面进行验证

到第一个sql节点登录mysql

```shell
mysql>create database dbtest;
mysql>use dbtest;
# 这里必须指定数据库表的引擎为NDB,否则同步失败
mysql>create table `country` (`id` int, `name` varchar(128)) ENGINE=NDB;
mysql>insert into `country` values (1, 'China');
mysql>select * from country;
```

到其他的sql节点登录mysql看数据是否同步过来, 经过测试, 在非master上创建数据, 可以同步到master上

## 关闭集群

关闭管理节点和数据节点, 只需要在管理节点执行

```shell
/usr/local/mysql/bin/ndb_mgm -e shutdown
```

关闭sql节点

```shell
/usr/local/mysql/support-files/mysql.server stop
```

要再次启动集群, 这次启动数据节点和sql节点的时候就不需要加"-initial"参数了
