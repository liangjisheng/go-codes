# common images

[minio](https://www.quanxiaoha.com/docker/docker-install-minio.html)
[jenkins](https://www.quanxiaoha.com/docker/docker-install-jenkins.html)

## ubuntu

```shell
#拉取最新版的 Ubuntu 镜像
docker pull ubuntu
docker pull ubuntu:latest
#启动容器
docker run -itd --name ubuntu-test ubuntu
#进入容器
docker exec -it ubuntu-test /bin/bash
```

## centos

```shell
docker pull centos:centos7
docker images
docker run -itd --name centos-test centos:centos7
docker exec -it centos-test /bin/bash
docker ps
```

## nginx

[配置 ssl](https://www.quanxiaoha.com/docker/docker-nginx-install-ssl.html)

```shell
docker search nginx
docker pull nginx:latest
docker images
docker run --name nginx-test -p 8080:80 -d nginx
#访问 8080 端口的 nginx 服务
http://127.0.0.1:8080/
```

## nodejs

```shell
docker search node
docker pull node:latest
docker run -itd --name node-test node
#进入容器查看运行的 node 版本
docker exec -it node-test /bin/bash
node -v
```

## php

```shell
docker search php
docker pull php:5.6-fpm
#-v ~/nginx/www:/www : 将主机中项目的目录 www 挂载到容器的 /www
docker run --name myphp-fpm -v ~/nginx/www:/www -d php:5.6-fpm

mkdir ~/nginx/conf/conf.d 
#在该目录下添加 ~/nginx/conf/conf.d/alice-test-php.conf 文件，内容如下：

#--link myphp-fpm:php: 把 myphp-fpm 的网络并入 nginx，并通过修改
# nginx 的 /etc/hosts，把域名 php 映射成 127.0.0.1，让 nginx 通过
# php:9000 访问 php-fpm
docker run --name alice-php-nginx -p 8083:80 -d \
    -v ~/nginx/www:/usr/share/nginx/html:ro \
    -v ~/nginx/conf/conf.d:/etc/nginx/conf.d:ro \
    --link myphp-fpm:php \
    nginx
```

alice-test-php.conf

```conf
server {
    listen       80;
    server_name  localhost;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm index.php;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    location ~ \.php$ {
        fastcgi_pass   php:9000;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  /www/$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

接下来我们在 ~/nginx/www 目录下创建 index.php，代码如下：

```php
<?php
echo phpinfo();
?>
```

浏览器打开 [http://127.0.0.1:8083/index.php](http://127.0.0.1:8083/index.php)

## mysql

```shell
docker search mysql
docker pull mysql:latest
#MYSQL_ROOT_PASSWORD=123456：设置 MySQL 服务 root 用户的密码
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
docker ps
```

## tomcat

```shell
docker search tomcat
docker pull tomcat
docker images | grep tomcat
docker run --name tomcat -p 8080:8080 -v $PWD/test:/usr/local/tomcat/webapps/test -d tomcat
```

浏览器打开 [http://127.0.0.1:8080/index.php](http://127.0.0.1:8080/index.php)

## python

```shell
docker search python
docker pull python:3.5
#-w /usr/src/myapp: 指定容器的 /usr/src/myapp 目录为工作目录
docker run -v $PWD/myapp:/usr/src/myapp -w /usr/src/myapp python:3.5 python helloworld.py
```

在 ~/python/myapp 目录下创建一个 helloworld.py 文件，代码如下：

```python
#!/usr/bin/python

print("Hello, World!");
```

## redis

```shell
docker search redis
docker pull redis:latest
docker run -itd --name redis-test -p 6379:6379 redis
docker exec -it redis-test /bin/bash
```

## mongodb

```shell
docker search mongo
docker pull mongo:latest
docker run -d -p 27017:27017 --name my-mongo-container mongo
#连接 db
mongosh --host 127.0.0.1 --port 27017

docker exec -it my-mongo-container bash
docker stop my-mongo-container
docker rm my-mongo-container
```

## apache

```shell
docker search httpd
docker pull httpd
docker run -p 80:80 -v $PWD/www/:/usr/local/apache2/htdocs/ -v $PWD/conf/httpd.conf:/usr/local/apache2/conf/httpd.conf -v $PWD/logs/:/usr/local/apache2/logs/ -d httpd
```

浏览器打开 [http://127.0.0.1](http://127.0.0.1)
