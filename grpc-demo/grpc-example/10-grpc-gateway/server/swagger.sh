#!/bin/bash

# 运行指令把Swagger UI转成Go代码
go-bindata --nocompress -pkg swagger -o swagger/datafile.go swagger/swagger-ui/...
