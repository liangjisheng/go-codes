# go fuzzing

[article](https://mp.weixin.qq.com/s/zdsrmlwVR0bP1Q_Xg_VlpQ)
[article](https://mp.weixin.qq.com/s/5qnIUz3plQG65FVnbPZVLw)
[go-fuzz](https://github.com/dvyukov/go-fuzz)

执行 fuzzing 测试

```shell
go test -v -fuzz=FuzzParseQuery
```

fuzz testing默认会一直执行下去，直到遇到crash。如果要限制fuzz testing的执行时间，可以使用 -fuzztime

```shell
go test -v -fuzztime 10s -fuzz=FuzzParseQuery
```
