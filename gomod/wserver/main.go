package main

import (
	"fmt"
	"wserver/route"

	"github.com/jinzhu/configor"
)

func main() {
	fmt.Println("user outside package test:", configor.Config{})
	route.Name()
}
