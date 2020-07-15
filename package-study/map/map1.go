package main

import "fmt"

// map是无序的,key不能重复，如果重复，相当于覆盖

func test1() {
	var map1 map[string]map[string]string

	map1 = make(map[string]map[string]string)

	map1["no1"] = make(map[string]string)
	map1["no1"]["name"] = "ljs1"
	map1["no1"]["hobby"] = "soccer"
	map1["no1"]["age"] = "20"

	map1["no2"] = make(map[string]string)
	map1["no2"]["name"] = "ljs2"
	map1["no2"]["hobby"] = "soccer"
	map1["no2"]["age"] = "21"

	map1["no2"] = make(map[string]string)
	map1["no2"]["name"] = "ljs3"
	map1["no2"]["hobby"] = "soccer"
	map1["no2"]["age"] = "22"

	fmt.Println(map1)
}
