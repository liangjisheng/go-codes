#!/bin/sh

protoc -I.\
 -I$GOPATH/src\
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:.\
 $1
