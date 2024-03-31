package main

import (
	"container/list"
	"fmt"
)

func list1() {
	l := list.New()
	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	fmt.Println("original list:")
	prtList(l)

	// for只循环了一次
	// for e := l.Front(); e != nil; e = e.Next() {
	// 	l.Remove(e)
	// }

	var next *list.Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}

	fmt.Println("remove all elements:")
	prtList(l)
}

func prtList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
