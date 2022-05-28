#!/bin/bash

go test -v -run='Test*' -count=2

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

#-benchmem 输出内存分配统计，还可以指定测试时间（-benchtime）、超时时间（-timeout）
go test -bench=".*" -benchmem