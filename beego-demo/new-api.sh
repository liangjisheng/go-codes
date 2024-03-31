#!/bin/bash

# function: 
# author  : liangjisheng
# date    : 2019/01/18 19:48:08
# version : 1.0

if [ $# -lt 1 ]
then
    echo -n "Usage: $0 project-name"
    exit 1
fi

if [ -d $1 ]
then
    echo -n "$1 already exist. please re-input project-name"
    exit 1
fi

if [ -f $GOPATH/bin/bee ]
then
    $GOPATH/bin/bee api $1
    cd $1
    go build main.go
else
    echo -n "$GOPATH/bin/bee executable file not exist."
    exit 1
fi

exit 0
