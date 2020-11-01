#!/bin/bash

# 编译 .o 文件
gcc -Wall -c hello.c

# 将 .o 文件打包成 .a 文件
ar -rv libhello.a hello.o
