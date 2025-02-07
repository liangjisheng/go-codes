# mongo

MongoDB是一个流行的开源文档型数据库，它使用类似 JSON 的文档模型存储数据，这使得数据存储变得非常灵活。

MongoDB 是一个基于文档的 NoSQL 数据库，由 MongoDB Inc. 开发。

MongoDB 旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。

MongoDB 是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。

MongoDB Atlas: MongoDB 提供的云服务，允许在云中托管 MongoDB 数据库

[mongodb](https://www.mongodb.com/)
[docs](https://docs.mongodb.com/manual/)
[download](https://www.mongodb.com/try/download/community)
[github](https://github.com/mongodb/mongo)
[mongodb 入门](https://www.runoob.com/mongodb/mongodb-tutorial.html)

## install

```shell
cd /usr/local
wget https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.17.tgz
wget https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-7.0.15.tgz
wget https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-6.0.19.tgz
#curl -O https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.17.tgz

wget https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-8.0.3.tgz
tar -zxvf mongodb-macos-x86_64-8.0.3.tgz
mv mongodb-macos-x86_64-8.0.3 mongodb
cd mongodb

mkdir log data
# 编辑配置文件
vim mongodb.conf
# 前台启动
# -f [ --config ] Configuration file specifying additional options
./bin/mongod -f mongodb.conf
# 查看 mongod 服务是否启动
ps aux | grep -v grep | grep mongod

vim ~/.zshrc
export PATH=$PATH:/usr/local/mongodb/bin
source ~/.zshrc
```

## mongosh

MongoDB6.0 以后做出了重大改变，MongoDB 已经不再默认安装 shell 工具，你需要安装一个额外的 shell 工具。用来连接 mongodb

[install](https://www.mongodb.com/zh-cn/docs/mongodb-shell/install/)

MongoDB Shell 是 MongoDB 提供的官方交互式界面，允许用户与 MongoDB 数据库进行交互、执行命令和操作数据库。

MongoDB Shell 是基于 JavaScript 的，允许用户直接在命令行或者脚本中使用 JavaScript 语言来操作 MongoDB 数据库。

[mongosh 使用文档](https://www.mongodb.com/docs/mongodb-shell/)

```shell
wget https://downloads.mongodb.com/compass/mongosh-2.3.3-darwin-x64.zip
unzip mongosh-2.3.3-darwin-x64.zip

vim ~/.zshrc
export PATH=$PATH:/usr/local/mongoos/bin
source ~/.zshrc
```

使用 mongosh

```shell
mongosh --version
# 使用默认值连接数据库, 启动 MongoDB Shell
mongosh
mongosh -u "username" -p "password" --authenticationDatabase "dbname"

test> show dbs
test> use study
study> db.mycollection.insertOne({ name: "Alice", age: 30 })
study> db.mycollection.find()
study> db.mycollection.updateOne({ name: "Alice" }, { $set: { age: 31 } })
study> db.mycollection.deleteOne({ name: "Alice" })
# 下面 2 个命令都可以退出 mongosh
study> exit
study> quit
```

## mongodb compass

mac 上安装可视化 mongodb 工具, 类似 navicat

```shell
wget https://downloads.mongodb.com/compass/mongodb-compass-1.44.6-darwin-x64.dmg
open mongodb-compass-1.44.6-darwin-x64.dmg
```

## mongos

mongs 作用：

为应用程序提供一个集群单一入口，一般情况下只使用一个 mongos，其它的 mongos 作为高可用存在。你的数据可能存在多个分片（shard）上，mongos 能够转发你的请求到相应的分片服务器上，并且把结果进行合并再返回给应用端。mongs 至少有两个，目的是做高可用，更多的 mongos 可以做 Load Balance 的效果。

## linux ubuntu 22.04

```shell
wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2204-6.0.6.tgz
```
