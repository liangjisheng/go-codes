#!/usr/bin/env sh

protoc library.proto -I=. --openapi_out=naming=proto,fq_schema_naming=1,version=1.2.3,default_response=true:.
