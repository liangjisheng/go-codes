package main

import (
	"fmt"
	"github.com/liyue201/gostl/ds/list/simplelist"
)

func listTest() {
	l := simplelist.New()
	l.PushBack(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushBack(4)
	for n := l.FrontNode(); n != nil; n = n.Next() {
		fmt.Printf("%v ", n.Value)
	}
}
