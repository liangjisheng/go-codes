#!/bin/sh

protoc -I.\
 -I$GOPATH/src\
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\
 --grpc-gateway_out=logtostderr=true:. $1
 # --swagger_out=logtostderr=true:. $1
