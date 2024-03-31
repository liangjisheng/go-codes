# mongo

[mongodb](https://www.mongodb.com/)
[mongodb](https://www.runoob.com/mongodb/mongodb-tutorial.html)

## install

```shell
cd /usr/local
sudo wget https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.17.tgz
#sudo curl -O https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.17.tgz
sudo tar -zxvf mongodb-macos-x86_64-6.0.6.tgz
sudo mv mongodb-macos-x86_64-6.0.6 mongodb
cd mongodb

vim ~/.zshrc
export PATH=$PATH:/usr/local/mongodb/bin
source ~/.zshrc

sudo mkdir log data

#--fork 表示后台运行
sudo mongod -f mongodb.conf 
```

```shell
mongod -f mongodb.conf
```

ubuntu 22.04

```shell
wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu2204-6.0.6.tgz
```

## command

```shell
show dbs;
use database;
show collections;
db.docs.find();
```
