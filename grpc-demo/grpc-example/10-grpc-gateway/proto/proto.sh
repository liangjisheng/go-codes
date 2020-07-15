#!/bin/bash

# 编译api下的http.pb.go
protoc --go_out=plugins=grpc:./ ./google/api/http.proto

# 编译api下annotations.pb.go
protoc --proto_path=. --go_out=plugins=grpc:./ ./google/api/annotations.proto

# 编译options下openapiv2.proto
protoc --proto_path=. --go_out=plugins=grpc:./ ./google/options/openapiv2.proto

# 编译options下annotations.proto
protoc --proto_path=. --go_out=plugins=grpc:./ ./google/options/annotations.proto

# 编译simple.validator.pb.go和simple.pb.go
protoc --proto_path=${GOPATH}/src --proto_path=. --govalidators_out=. --go_out=plugins=grpc:. ./simple.proto

# 编译simple.pb.gw.go
protoc --proto_path=${GOPATH}/src --proto_path=. --grpc-gateway_out=logtostderr=true:. ./simple.proto

# 编译simple.swagger.json
protoc --proto_path=${GOPATH}/src --proto_path=. --swagger_out=logtostderr=true:./ ./simple.proto
