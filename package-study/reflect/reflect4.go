package main

import (
	"fmt"
	"reflect"
)

type ss4 struct {
	s struct {
		B int
		b int
	}
	A int
	a int
}

func reflect4() {
	var v = reflect.ValueOf(ss4{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Println("字段:", field.Type().String())
		if field.Type().Kind() == reflect.Struct {
			for j := 0; j < field.NumField(); j++ {
				subField := field.Field(j)
				fmt.Println("    嵌套字段:", subField.Type().String())
			}
		}
	}
}
