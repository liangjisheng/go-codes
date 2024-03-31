# volume

简单来说，数据卷是一个可供一个或多个容器使用的特殊目录，用于持久化数据以及共享容器间的数据，它以
正常的文件或目录的形式存在于宿主机上。 另外，其生命周期独立于容器的生命周期，即当你删除容器时，
数据卷并不会被删除。

Docker 镜像由多个文件系统（只读层）叠加而成。Docker 会加载只读镜像层，并在镜像栈顶部添加一个读写
层。当运行容器后，如果修改了某个已存在的文件，那么该文件将会从下面的只读层复制到上面的读写层，同时，
该文件在只读层中仍然存在。当我们删除 Docker 容器，并通过镜像重新启动容器时，之前的更改的文件将会丢失。

数据卷可以在容器之间共享和重用；
对数据卷的修改会立刻生效；
更新数据卷不会影响镜像；
数据卷默认一直存在，即使容器被删除；

Docker 提供了 3 种不同的方式将数据从宿主机挂载到容器中。

volume : Docker 管理宿主机文件系统的一部分，默认位于 /var/lib/docker/volumes 目录下, 也是最
常用的方式。
所有的 Docker 容器数据都保存在 /var/lib/docker/volumes 目录下。若容器运行时未指定数据卷， Docker
创建容器时会使用默认的匿名卷（名称为一堆很长的 ID）。

bind mount: 意为可以存储在宿主机中的任意位置。需要注意的是，bind mount 在不同的宿主机系统时不可移植的
比如 Windows 和 Linux 的目录结构是不一样的，bind mount 所指向的 host 目录也不一样。这也是为什么 
bind mount 不能出现在 Dockerfile 中的原因所在，因为这样 Dockerfile 就不可移植了。

tmpfs mount : 挂载存储在宿主机的内存中，而不会写入宿主机的文件系统，一般不用此种方式。

```shell
#创建数据卷
docker volume create test-vol
#查看所有的数据卷
docker volume ls
# 查看数据卷名为 test-vol 的信息
docker volume inspect test-vol
#运行容器时挂载数据卷
docker run -d -it --name=test-nginx -p 8011:80 -v test-vol:/usr/share/nginx/html nginx:1.13.12
docker run -d -it --name=test-nginx -p 8011:80 --mount source=test-vol,target=/usr/share/nginx/html nginx:1.13.12

#由于数据卷的生命期独立于容器，想要删除数据卷，就需要我们手动来操作, 执行命令如下
docker volume rm test-vol
#如果你需要在删除容器的同时移除数据卷，请使用 docker rm -v 命令。
#对于那些没有被使用的数据卷，可能会占用较多的磁盘空间，你可以通过如下命令统一删除：
docker volume prune
```

-v 和 --mount 有什么区别
都是挂载命令，使用 -v 挂载时，如果宿主机上没有指定文件不会报错，会自动创建指定文件；当使用 --mount时，
如果宿主机中没有这个文件会报错找不到指定文件，不会自动创建指定文件。

挂载成功后，我们不论是修改 /var/lib/docker/volumes 下的数据，还是进入到容器中修改 
/usr/share/nginx/html 下的数据，都会同步修改对应的挂载目录，类似前端开发中双向绑定的作用。

## 数据卷容器

如果你有一些需要持续更新的数据需要在容器之间共享，最佳实践是创建数据卷容器。数据卷容器，其实就是一个正常的
Docker 容器，专门用于提供数据卷供其他容器挂载的。

运行一个容器，并创建一个名为 dbdata 的数据卷

docker run -d -v /dbdata --name dbdata training/postgres echo Data-only container for postgres

容器运行成功后，会发现该数据卷容器处于停止运行状态，这是因为数据卷容器并不需要处于运行状态，只需用于提供数据卷挂载即可。

--volumes-from 命令支持从另一个容器挂载容器中已创建好的数据卷

```shell
docker run -d --volumes-from dbdata --name db1 training/postgres
docker run -d --volumes-from dbdata --name db2 training/postgres
docker ps
CONTAINER ID       IMAGE                COMMAND                CREATED             STATUS              PORTS               NAMES
7348cb189292       training/postgres    "/docker-entrypoint.   11 seconds ago      Up 10 seconds       5432/tcp            db2
a262c79688e8       training/postgres    "/docker-entrypoint.   33 seconds ago
```

如果删除了挂载的容器（包括 dbdata、db1 和 db2），数据卷并不会被自动删除。如果想要删除一个数据卷，
必须在删除最后一个还挂载着它的容器时使用 docker rm -v 命令来指定同时删除关联的容器。
