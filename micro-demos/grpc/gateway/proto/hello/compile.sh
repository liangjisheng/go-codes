#!/bin/bash

# Generate gRPC stub hello.pb.go
protoc -I. \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
  --go_out=plugins=grpc:. hello.proto

# Generate reverse-proxy hello.pb.gw.go
protoc -I. \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
  --grpc-gateway_out=logtostderr=true:. hello.proto
