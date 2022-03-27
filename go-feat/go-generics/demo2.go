package main

import (
	"fmt"
	"strconv"
)

//泛型中的接口本身对范型进行约束

type MyStringer interface {
	String() string
}

type StringInt int
type myString string

func (i StringInt) String() string {
	return strconv.Itoa(int(i))
}

func (str myString) String() string {
	return string(str)
}

//在泛型方法中，我们声明了泛型的类型为：任意实现了MyStringer接口的类型；只要实现了这个接口，那么你就可以直接使用
func stringify[T MyStringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

func case5() {
	fmt.Println(stringify([]StringInt{1, 2, 3, 4, 5}))
	fmt.Println(stringify([]myString{"1", "2", "3", "4", "5"}))

	//直接使用 int 会报错
	//只有实现了 MyStringer 接口的类型才会被允许作为实参传递给 stringify 泛型函数的类型参数并成功实例化
	//int does not implement MyStringer (missing String method)
	//fmt.Println(stringify([]int{1, 2, 3, 4, 5}))
}

//MySignedStringer 当然也可以将MyStringer接口写成如下的形式
//表示只有int, int8, int16, int32, int64，这样类型参数的实参类型既要在MySignedStringer的类型列表中，
//也要实现了MySignedStringer的String方法，才能使用。像这种不在里面的type StringInt uint就会报错
type MySignedStringer interface {
	int | int8 | int16 | int32 | int64
	String() string
}
