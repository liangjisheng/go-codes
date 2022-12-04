#!/usr/bin/env sh

protoc -I . --go_out=":." --validate_out="lang=go:." example.proto
