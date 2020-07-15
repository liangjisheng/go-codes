#!/bin/bash

protoc --go_out=plugins=grpc:./ ./client_stream.proto
