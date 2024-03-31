#!/bin/bash

COMMIT_ID=`git log |head -n 1| awk '{print $2;}'`
#COMMIT_ID=`git rev-parse --short HEAD`
AUTHOR=`git log |head -n 3| grep Author| awk '{print $2;}'`
# 匹配以 * 开头的行
BRANCH_NAME=`git branch | awk '/\*/ { print $2; }'`
SERVICE_INFO="$COMMIT_ID,$AUTHOR,$BRANCH_NAME"
echo $SERVICE_INFO
go build -ldflags "-X main.BuildUser=$SERVICE_INFO"
