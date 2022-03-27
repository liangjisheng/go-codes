package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	name string
}

type EmptyUser struct {
	name struct{}
	age  struct{}
}

//Address 空结构体地址相同 大小为 0
func Address() {
	var user1 User
	var user2 User
	fmt.Printf("%p\n", &user1)
	fmt.Printf("%p\n", &user2)

	var v1 struct{}
	var v2 struct{}
	fmt.Printf("%p\n", &v1)
	fmt.Printf("%p\n", &v2)

	fmt.Println(unsafe.Sizeof(v1))
	fmt.Println(unsafe.Sizeof(v2))

	var u EmptyUser
	fmt.Printf("%p\n", &u)
	fmt.Println(unsafe.Sizeof(u))
}

func set() {
	set := NewSet()
	set.Add("hello")
	set.Add("world")
	fmt.Println(set.Contains("hello"))

	set.Remove("hello")
	fmt.Println(set.Contains("hello"))
}

func main() {
	//Address()
	//set()
	Channel()
}
