package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	High  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

// 如果我们想直接解析到struct Class对象中，可以将接受体对应
// 的项定义为该struct类型
type StuRead1 struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	High  interface{}
	sex   interface{}
	Class Class `json:"class"`
	Test  interface{}
}

type StuRead2 struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	High  interface{}
	sex   interface{}
	Class *Class `json:"class"`
	Test  interface{}
}

// 如果不想指定Class变量为具体的类型，仍想保留interface{}类型，
// 但又希望该变量可以解析到struct Class对象中
// 我们可以将该变量定义为json.RawMessage类型
type StuRead3 struct {
	Name  interface{}
	Age   interface{}
	HIgh  interface{}
	Class json.RawMessage `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

// 利用反射，打印变量类型
func printType(stu *StuRead) {
	nameType := reflect.TypeOf(stu.Name)
	ageType := reflect.TypeOf(stu.Age)
	highType := reflect.TypeOf(stu.High)
	sexType := reflect.TypeOf(stu.sex)
	classType := reflect.TypeOf(stu.Class)
	testType := reflect.TypeOf(stu.Test)

	fmt.Println("nameType:", nameType)
	fmt.Println("ageType:", ageType)
	fmt.Println("highType:", highType)
	fmt.Println("sexType:", sexType)
	fmt.Println("classType:", classType)
	fmt.Println("testType:", testType)
}

func rawmessage1() {
	// 使用原生字符串
	data := `{
		"name":"张三",
		"Age":18,
		"high":true,
		"sex":"男",
		"CLASS":{"naME":"1班","Grade": 3}
		}`
	str := []byte(data)

	stu := StuRead{}
	fmt.Println(stu)
	// interface{}类型变量在json解析前，打印出的类型都为nil，
	// 就是没有具体类型，这是空接口（interface{}类型）的特点
	printType(&stu)

	// json串中key为CLASS的value是个复合结构，不是可以直接解析
	// 的简单类型数据（如“张三”，18，true等）。所以解析时，
	// 由于没有指定变量Class的具体类型，json自动将value为复合
	// 结构的数据解析为map[string]interface{}类型的项。也就是
	// 说，此时的struct Class对象与StuRead中的Class变量没有
	// 半毛钱关系，故与这次的json解析没有半毛钱关系
	err := json.Unmarshal(str, &stu)
	fmt.Println("----------------json解析后-----------")

	if err != nil {
		fmt.Println(err)
		return
	}

	printType(&stu)
	fmt.Println(stu)
	fmt.Println()

	stu1 := StuRead1{}
	err = json.Unmarshal(str, &stu1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu1)
	fmt.Println()

	stu2 := StuRead2{}
	err = json.Unmarshal(str, &stu2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu2)
	fmt.Println()

	stu3 := StuRead3{}
	err = json.Unmarshal(str, &stu3)
	// 二次解析
	cla := new(Class)
	// 接收体中，被声明为json.RawMessage类型的变量在json解析时，
	// 变量值仍保留json的原值，即未被自动解析为
	// map[string]interface{}类型
	err = json.Unmarshal(stu3.Class, cla)
	fmt.Println("stu3:", stu3)
	fmt.Println("string(stu.Class):", string(stu3.Class))
	fmt.Println("class:", cla)
}
