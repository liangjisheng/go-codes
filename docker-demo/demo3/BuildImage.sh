#!/bin/bash

# function: 
# author  : liangjisheng
# date    : 2019/02/18 16:43:51
# version : 1.0

# -t 指定要创建的目标镜像名, .:Dockerfile 文件所在目录
docker build -t demo:1.0 .

exit 0
