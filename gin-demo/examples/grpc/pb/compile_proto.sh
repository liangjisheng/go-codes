#!/bin/bash

protoc -I ./ $1 --go_out=plugins=grpc:.
# protoc --go_out=plugins=grpc:. $1
