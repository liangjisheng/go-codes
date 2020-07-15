package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func judge(v interface{}) {
	fmt.Printf("%p %+v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p %+v\n", &v, v)
		fmt.Printf("nil type[%T] %+v\n", v, v)

	case student:
		fmt.Printf("%p %+v\n", &v, v)
		fmt.Printf("student type[%T] %+v\n", v, v)

	case *student:
		fmt.Printf("%p %+v\n", &v, v)
		fmt.Printf("*student type[%T] %+v\n", v, v)

	default:
		fmt.Printf("%p %+v\n", &v, v)
		fmt.Printf("unknown\n")
	}
}

func main() {
	// var i interface{} = new(student)
	// var i interface{} = (*student)(nil)
	var i interface{}

	fmt.Printf("%p %+v\n", &i, i)

	judge(i)
}
