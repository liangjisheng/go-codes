package main

import (
	"github.com/kr/pretty"
)

type User struct {
	Name string
	Age  int
}

func main() {
	type myType struct {
		a, b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	pretty.Println(x)

	u := User{
		Name: "alice",
		Age:  18,
	}
	pretty.Println(u)
}
