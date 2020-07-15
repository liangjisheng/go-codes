#!/bin/bash

vim /etc/apt/sources.list.d/pgdg.list

# ubuntu16.04
deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main
# ubuntu18.04
# deb http://apt.postgresql.org/pub/repos/apt/ bionic-pgdg main

wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

sudo apt-get update

sudo apt-get install postgresql-12 postgresql-client-12 pgadmin4 libpq-dev postgresql-server-dev-12

vim /etc/postgresql/12/main/postgresql.conf

listen_addresses=’*’
port=5432

vim /etc/postgresql/12/main/pg_hba.conf

# 增加
host all all 0.0.0.0/0 md5

sudo service postgresql restart

sudo service postgresql status
sudo service postgresql start
sudo service postgresql stop
sudo service postgresql restart

# 手动启动
sudo /etc/init.d/postgresql start
sudo /etc/init.d/postgresql stop
sudo /etc/init.d/postgresql restart
sudo /etc/init.d/postgresql status
