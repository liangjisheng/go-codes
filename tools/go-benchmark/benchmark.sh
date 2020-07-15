#!/bin/bash

go test --run TestInsertionSort -v
go test --run TestQuickSort -v

go test -run BenchmarkInsertionSort10 -benchmem -bench .
go test -run BenchmarkQuickSort10 -benchmem -bench .

go test -run BenchmarkInsertionSort100 -benchmem -bench .
go test -run BenchmarkQuickSort100 -benchmem -bench .

go test -run BenchmarkInsertionSort1000 -benchmem -bench .
go test -run BenchmarkQuickSort1000 -benchmem -bench .

go test -run BenchmarkInsertionSort10000 -benchmem -bench .
go test -run BenchmarkQuickSort10000 -benchmem -bench .

go test -run BenchmarkInsertionSort100000 -benchmem -bench .
go test -run BenchmarkQuickSort100000 -benchmem -bench .

go test -run BenchmarkInsertionSort1000000 -benchmem -bench .
go test -run BenchmarkQuickSort1000000 -benchmem -bench .

go test -benchtime=5s -benchmem -bench=BenchmarkQuickSort1000000 -run=none
