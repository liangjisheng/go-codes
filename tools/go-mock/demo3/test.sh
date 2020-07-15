#!/bin/bash

go test --run=TestReturn -v
go test --run=TestReturnDynamic -v
go test --run=TestTimes -v
go test --run=TestOrder -v
