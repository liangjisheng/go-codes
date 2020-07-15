package main

import (
	"fmt"
	"sync"
)

func pool1() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello,World!"
		},
	}
	value := "hello, ljs"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

func pool2() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	p.Put("jiangzhou")
	p.Put(123456)
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
