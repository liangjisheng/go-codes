package main

import "fmt"

func sliceMap() {
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
