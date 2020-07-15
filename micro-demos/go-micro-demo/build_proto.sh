#!/bin/bash

protoc --proto_path=./proto --micro_out=./src/share/pb --go_out=./src/share/pb ./proto/*.proto
