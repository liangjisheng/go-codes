# user privileges

mysql 数据库用户在创建的时候就会赋予 USAGE 权限, 这个权限很小几乎为0
只能连接数据库和查询 information_schema 的权限
不过这个权限也很奇怪, 你无法 revoke

```sql
-- 创建用户只能本地访问
create user mysql@'localhost' identified by 'password';
create user mysql@'%' identified by 'password';
-- 创建user02,可以远程访问
create user readonly@'%' identified by 'password';
-- 删除用户
drop user 'username'@'host';
-- 修改密码
set password for 'user01'@'localhost'=password('anotherpassword');
set password for 'root'@'%'=password('password');
-- mysql5.x 修改密码
set password for root@localhost = password('newpass');
alter user 'root'@'localhost' identified with mysql_native_password by 'password';
alter user 'zabbix'@'localhost' identified with mysql_native_password by 'zabbix';
-- mysql8.x 修改密码
alter user 'root'@'localhost' identified by 'password';

-- 授予 select 权限， 可以将 %替换成特定的IP 或者 ip 段 192.168.1.%
grant select on dev.* to 'mysql'@'%';
-- 授予管理 test db 的全部权限
grant all privileges on test.* to user01;
grant all privileges on *.* to mysql;
-- 撤销用户权限
revoke select on dev.* from 'readonly'@'%';
```

使用 grant 命令时, 其会自动通知MySQL服务器重新加载一次权限数据。以达到即时生效的效果

改表法修改权限，以设置root用户允许远程连接为例

```sql
use mysql;
update user set host="%" where user="root";
flush privileges;

-- 通过以下语句检查是否生效
show grants for 'root'@'%';
```

但当我们使用改表法时。是没有通知重新加载权限数据的。因此会导致其不会即时生效。直至服务重启后生效。服务重启，特别是生产环境，
那几乎是灾难性的。好在MySQL为我们提供了手动通知的命令。即：flush privilege命令

```sql
-- 刷新权限
flush privileges;
```

查看用户权限

```sql
-- 查看某个用户权限
show grants for username;
-- current_user() 当前登录用户
show grants for current_user();
```

登录使用 auth_socket 插件， 首先，这种验证方式不要求输入密码，即使输入了密码也不验证。这个特点让很多人觉得很不安全，
实际仔细研究一下这种方式，发现还是相当安全的，因为它有另外两个限制；
只能用 UNIX 的 socket 方式登陆，这就保证了只能本地登陆，用户在使用这种登陆方式时已经通过了操作系统的安全验证；
操作系统的用户和 MySQL 数据库的用户名必须一致，例如你要登陆 MySQL 的 root 用户，必须用操作系统的 root 用户登陆
