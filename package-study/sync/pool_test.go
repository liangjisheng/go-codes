package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool1(t *testing.T) {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello,World!"
		},
	}

	value := "hello, ljs"
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

	p.Put("ljs")
	p.Put(123456)
	t.Log(p.Get())
	t.Log(p.Get())
	t.Log(p.Get())
}

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	t.Log(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	t.Log(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
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
