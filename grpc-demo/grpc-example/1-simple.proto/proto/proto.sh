#!/bin/bash

protoc --go_out=plugins=grpc:./ ./simple.proto
