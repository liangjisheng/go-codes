#!/bin/bash

protoc --proto_path=${GOPATH}/src --proto_path=. --govalidators_out=. --go_out=plugins=grpc:. *.proto
