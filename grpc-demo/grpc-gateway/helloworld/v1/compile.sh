#!/bin/bash

#2种方式生成对应的go文件

#实现 server 接口的 struct 不需要内置 pb.UnimplementedGreeterServer
#protoc --go_out=plugins=grpc:./ ./helloworld.proto
#protoc --grpc-gateway_out=./ ./helloworld.proto

#实现 server 接口的 struct 需要内置 pb.UnimplementedGreeterServer
#protoc --go_out=./ ./hello_world.proto
#protoc --go-grpc_out=./ ./hello_world.proto
#protoc --grpc-gateway_out=./ ./hello_world.proto
#protoc --openapiv2_out=./ --openapiv2_opt logtostderr=true ./hello_world.proto

protoc --go_out=./ --go-grpc_out=./ --grpc-gateway_out=./ --openapiv2_out=./ --openapiv2_opt logtostderr=true ./hello_world.proto

