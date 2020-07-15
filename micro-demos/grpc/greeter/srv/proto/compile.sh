#!/bin/bash

protoc --proto_path=. --micro_out=. --go_out=plugins=grpc:. hello.proto