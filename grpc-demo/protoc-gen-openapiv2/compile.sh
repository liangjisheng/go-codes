#!/usr/bin/env sh

protoc --proto_path=. --go_out=paths=source_relative:. --openapiv2_out=:. \
  --openapiv2_opt=output_format=yaml,version=true,generate_unbound_methods=true \
  --openapiv2_opt=logtostderr=true,json_names_for_fields=false,enums_as_ints=true,allow_delete_body=true \
  example.proto