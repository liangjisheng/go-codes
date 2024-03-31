# go inject tag

[go inject tag](https://www.jianshu.com/p/744d8c080d59)
[github](https://github.com/favadi/protoc-go-inject-tag)

```shell
protoc --proto_path=. --go_out=. inject.proto && protoc-go-inject-tag -input=./inject/inject.pb.go
```

注释里面可以使用 inject_tags, 也可以使用 gotags, inject_tags 是 v1.3.0 版本之前的, 现在强烈建议使用 gotags
