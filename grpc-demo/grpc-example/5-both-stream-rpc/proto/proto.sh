#!/bin/bash

protoc --go_out=plugins=grpc:./ ./both_stream.proto
