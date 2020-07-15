#!/bin/bash

go test -run TestAdd -v
go test -run TestMul -v

go test -run TestMul1/pos -v
go test -run TestMul1/neg -v

go test -run TestMul2 -v
go test -run TestMul3 -v

go test -run TestConn -v
go test -run TestConn1 -v

go test -run BenchmarkHello -benchmem -bench .
go test -run BenchmarkParallel -benchmem -bench .
