package main

import (
	"fmt"
	"sync"
	"time"
)

func syncMap() {
	var sm sync.Map
	sm.Store("name", "zhangsan")
	sm.Store("addr", "beijing")
	sm.Store("job", "programmer")

	name, ok := sm.Load("name")
	if ok {
		fmt.Println(name)
	}

	sm.Delete("nameNone") // safe delete

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, "===>", value)
		return true
	})

	// 测试并发读写使用
	go func() {
		for {
			t := time.Now().Unix()
			sm.Store("name", t)
		}
	}()

	go func() {
		for {
			name, _ := sm.Load("name")
			fmt.Println(name)
		}
	}()

	select {}
}
