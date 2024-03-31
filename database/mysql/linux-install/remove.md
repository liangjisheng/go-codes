# remove

```shell
#remove mysql
sudo apt-get -y remove mysql-server
sudo apt-get -y remove mysql-server-5.7 mysql-common
sudo apt-get -y autoremove
#查看通过 apt 安装的 mysql 包，有的话指定包名移除
dpkg -l | grep mysql | grep ii
```