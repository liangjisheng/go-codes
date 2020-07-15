package main

import (
	"fmt"

	"github.com/prashantv/gostub"
)

var counter = 100

// 为一个全局变量打桩
func stubGlobalVar() {
	stubs := gostub.Stub(&counter, 200)
	// Reset方法将全局变量的值恢复为原值
	defer stubs.Reset()
	fmt.Println("counter:", counter)
}
