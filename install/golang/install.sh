#!/bin/bash

wget https://studygolang.com/dl/golang/gogo1.20.3.linux-amd64.tar.gz
wget https://golang.google.cn/dl/go1.20.3.linux-amd64.tar.gz
tar zxf go1.20.3.linux-amd64.tar.gz

vim ~/.bashrc

export GO111MODULE=auto
export GOPROXY=https://goproxy.io
export GOROOT=/data/go
export GOPATH=/data/gopath
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOBIN:$PATH

source ~/.bashrc

# 通过命令行直接安装
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt-get install golang-go
