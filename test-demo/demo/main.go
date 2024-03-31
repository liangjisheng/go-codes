package main

import (
	"fmt"
	"sync"
)

func demo() {
	m := map[string]string{
		"alice": "alice",
	}
	m = nil

	res, ok := m["alice"]
	if !ok {
		fmt.Println("not ok")
		return
	}
	fmt.Println(res)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}

var (
	sMap = sync.Map{}
)

func demo1() {
	key := "key"
	value := []string{"a", "b"}

	sMap.Store(key, value)

	v1I, ok := sMap.Load(key)
	if !ok {
		return
	}

	v1, ok := v1I.([]string)
	if !ok {
		return
	}
	fmt.Println(v1)

	v1 = append(v1, "c")
	sMap.Store(key, v1)

	v1I, ok = sMap.Load(key)
	if !ok {
		return
	}

	v1, ok = v1I.([]string)
	if !ok {
		return
	}
	fmt.Println(v1)
}

func main() {
	demo1()
}
