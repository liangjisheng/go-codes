package main

import (
	"fmt"
	"maps"
)

func main() {
	//map1()
	//map2()
	//genericMap()
	map3()
}

func map3() {
	m1 := make(map[int]int, 0)
	m1[0] = 0
	m1[1] = 1
	m1[2] = 2

	m2 := make(map[int]int, 0)
	maps.Copy(m2, m1)
	fmt.Println(m2)

	equal := maps.Equal(m1, m2)
	fmt.Println(equal)

	m3 := maps.Clone(m1)
	fmt.Println(m3)
}

// map 是无序的,key不能重复，如果重复，相当于覆盖

func map1() {
	var v1 map[string]map[string]string

	v1 = make(map[string]map[string]string)

	v1["no1"] = make(map[string]string)
	v1["no1"]["name"] = "alice1"
	v1["no1"]["hobby"] = "soccer"
	v1["no1"]["age"] = "20"

	v1["no2"] = make(map[string]string)
	v1["no2"]["name"] = "alice2"
	v1["no2"]["hobby"] = "soccer"
	v1["no2"]["age"] = "21"

	v1["no2"] = make(map[string]string)
	v1["no2"]["name"] = "alice3"
	v1["no2"]["hobby"] = "soccer"
	v1["no2"]["age"] = "22"

	fmt.Println(v1)
}

func map2() {
	var monsters []map[string]string
	// 给切片分配空间
	monsters = make([]map[string]string, 3)

	// 给第一个妖怪的map分配空间
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "红孩儿"
		monsters[0]["age"] = "10"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "牛魔王"
		monsters[1]["age"] = "500"
	}

	if monsters[2] == nil {
		monsters[2] = make(map[string]string, 2)
		monsters[2]["name"] = "白骨精"
		monsters[2]["age"] = "400"
	}

	fmt.Println(monsters)
}
