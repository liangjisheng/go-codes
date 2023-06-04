//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

//wire.Bind 将 Struct 和接口进行绑定了，表示这个结构体实现了这个接口

//在上面 NewPostService  代码，我们可以发现有很多 Struct  的初始化其实就是填充里面的属性，没有其他的逻辑，
//这种情况我们可以偷点懒直接使用 wire.Struct  方法直接生成 provider
// structType: 结构体类型
// fieldNames: 需要填充的字段，使用 "*" 表示所有字段都需要填充
//Struct(structType interface{}, fieldNames ...string)

//除了依赖某一个类型之外，有时候我们还会依赖一些具体的值，这时候我们就可以使用 wire.Value
//或者是 wire.InterfaceValue ，为某个类型绑定具体的值

// wire.Value 为某个类型绑定值，但是不能为接口绑定值
//Value(interface{}) ProvidedValue
// wire.InterfaceValue 为接口绑定值
//InterfaceValue(typ interface{}, x interface{}) ProvidedValue

//func GetPostService() (*PostService, func(), error) {
//	panic(wire.Build(
//		//NewPostService,
//		wire.Struct(new(PostService), "*"),
//		wire.Value(10),
//		wire.InterfaceValue(new(io.Reader), os.Stdin),
//		wire.Bind(new(IPostUsecase), new(*PostUsecase)),
//		NewPostUsecase,
//		NewPostRepo,
//	))
//}

func GetPostService() (*PostService, func(), error) {
	panic(wire.Build(
		PostServiceSet,
	))
}
