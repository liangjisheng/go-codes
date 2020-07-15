#!/bin/bash

protoc product_service.proto --go_out=plugins=grpc:.