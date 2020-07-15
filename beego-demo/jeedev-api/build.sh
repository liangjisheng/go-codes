#!/bin/bash

# function: 
# author  : liangjisheng
# date    : 2019/01/20 16:39:55
# version : 1.0

bee api jeedev-api -conn="root:password@tcp(127.0.0.1:3306)/jeedev"

cd jeedev-api

bee generate docs

bee run watchall true

exit 0
