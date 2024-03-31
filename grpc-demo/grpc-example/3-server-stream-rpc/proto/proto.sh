#!/bin/bash

protoc --go_out=plugins=grpc:./ ./server_stream.proto
