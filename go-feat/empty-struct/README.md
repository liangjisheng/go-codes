# go empty struct

我们说不包含任何字段的结构体叫做空结构体，可以通过如下的方式定义空结构体

原生定义

```go
var a struct{}
```

类型别名

```go
type empty struct{}
var e empty
```

使用场景，如果某个场景使用 map 只关心 key，而不关心 value 的话，可以使用 struct{} 作为 value

总结

1. 空结构体是一种特殊的结构体，不包含任何元素
2. 空结构体的大小都为0
3. 空结构体的地址都相同
4. 由于空结构体不占用空间，从节省内存的角度出发，适用于实现Set结构、在 channel 中传输信号等
