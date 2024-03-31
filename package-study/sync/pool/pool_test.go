package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestPool1(t *testing.T) {
	standard()
	withPool()

	withoutPool1()
	withPool1()
}

//当程序使用完对象后，可以将对象放回到 pool 中。但是需要注意的是，当对象被放回到 pool 中后，
//它并不保证立即可用，因为 pool 的策略是在池中保留一定数量的对象，超出这个数量的对象会被销毁。

func TestSyncPool1(t *testing.T) {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello,World!"
		},
	}

	value := "hello, alice"
	pool.Put(value)
	t.Log(pool.Get())
	t.Log(pool.Get())
}

func TestSyncPool2(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	p.Put("alice")
	p.Put(123456)
	t.Log(p.Get())
	t.Log(p.Get())
	t.Log(p.Get())
}

func TestSyncPool3(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	t.Log(v) // 100
	pool.Put(3)
	v1, _ := pool.Get().(int)
	t.Log(v1) // 3
}

func TestSyncPool4(t *testing.T) {
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
	c := p.Get().(int)
	d := p.Get().(int)
	fmt.Println(b, c, d, p.Get())
}

func TestSyncPool5(t *testing.T) {
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

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
