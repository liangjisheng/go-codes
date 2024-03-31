#!/bin/bash

# 编译 .o 文件
gcc -c hello.c

# 生成静态库
ar -cru libhello.a hello.o
