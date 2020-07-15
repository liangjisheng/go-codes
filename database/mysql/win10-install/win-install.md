# install mysql on win10

## mysql-8.0.17-winx64.zip 安装说明

1. 解压到任一目录,进入mysql-8.0.17-winx64目录
2. 创建my.ini文件
3. 打开cmd,进入bin目录下执行 mysqld --initialize --console
4. 注意后续的几句话里面，root@localhost后面有初始密码，复制下来
5. 使用cmd（以管理员身份运行）在bin目录下执行mysqld --install
6. 启动MySql服务 net start mysql
7. mysql -u root -p  输入之前复制的密码
8. ALTER user 'root'@'localhost' IDENTIFIED BY '新密码'
