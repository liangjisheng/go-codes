#!/bin/bash

# function: 
# author  : liangjisheng
# date    : 2019/01/19 10:59:56
# version : 1.0

# 生成一个以1.11版本Go环境为基础镜像的Dockerfile,并暴露9000端口
bee dockerize -image="library/golang:1.11" -expose=9000

exit 0
