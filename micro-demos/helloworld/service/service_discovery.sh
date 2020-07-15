#!/bin/bash

# install consul
# go get -u -v github.com/hashicorp/consul
# cd $GOPATH/src/github.com/hashicorp/consul
# make tools
# make dev

# run consul
consul agent -dev -advertise=127.0.0.1