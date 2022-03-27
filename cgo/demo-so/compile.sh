#!/bin/bash

#编译生成 so 文件
go build -buildmode=plugin -o aplugin.so aplugin.go
