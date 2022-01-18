#!/usr/bin/bash

# 根据swagger规范 创建 swagger.json 规范文档
swagger generate spec -o ./swagger.json --scan-models

# 执行该命令后，将使用 Petstore 托管的 SwaggerUI 打开一个新选项卡
# 启动一个http 服务同时将json文档放入http://petstore.swagger.io 执行
swagger serve -F=swagger swagger.json

# 如果使用Redoc flavor(-F=redoc)文档,将托管在您自己的计算机上(localhost:port/docs)
swagger serve -F=redoc swagger.json
swagger serve swagger.json

# statik是自动生成的,用于初始化文件目录,作为文件服务器
# statik -src=/c/Users/12948/go/src/go-demos/tools/swaggerui
