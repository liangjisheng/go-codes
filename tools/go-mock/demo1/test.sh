#!/bin/bash

go test --run TestGetFromDB . -cover -v
go test --run TestGetFromDB1 . -cover -v
go test --run TestGetFromDB2 . -cover -v
