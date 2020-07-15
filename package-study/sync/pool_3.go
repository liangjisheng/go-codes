package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func pool4() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	fmt.Println(a)
	p.Put(1)
	p.Put(4)
	p.Put(2)
	p.Put(5)

	b := p.Get().(int)
	// runtime.GC()  // 执行回收操作后 1 0 0 0
	c := p.Get().(int)
	d := p.Get().(int)
	fmt.Println(b, c, d, p.Get())
}

func pool5() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	runtime.GOMAXPROCS(2)

	a := p.Get().(int)
	fmt.Println(a)
	p.Put(1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		p.Put(100)
	}()
	wg.Wait()

	time.Sleep(time.Second * 1)

	p.Put(4)
	p.Put(5)

	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	// fmt.Println(p.Get())

	// 有趣的输出结果
	// 1:  0  1  5  4
	// 2:  0  100  5  4
	// 3:  0   4  5  100
}
