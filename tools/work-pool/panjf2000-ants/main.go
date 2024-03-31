package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	//fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	//fmt.Println("Hello World!")
}

func d1() {
	defer ants.Release()
	runTimes := 10

	// Use the common pool.
	//没有限制 goroutine 数量
	var wg sync.WaitGroup
	task := func() {
		fmt.Println("hello")
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(task)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}

func d2() {
	//同时允许2个并发任务
	p, _ := ants.NewPool(2)
	defer p.Release()

	var wg sync.WaitGroup
	task := func() {
		defer wg.Done()

		fmt.Println("hello")
		time.Sleep(time.Second)
	}

	runTimes := 10
	for i := 0; i < runTimes; i++ {
		wg.Add(1)

		_ = p.Submit(task)
	}
	wg.Wait()
}

func d3() {
	runTimes := 1000
	var wg sync.WaitGroup

	task := func(i interface{}) {
		myFunc(i)
		wg.Done()
	}

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, task)
	defer p.Release()

	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}

func multiPool1() {
	runTimes := 10
	var wg sync.WaitGroup
	task := func() {
		defer wg.Done()

		fmt.Println("hello")
		time.Sleep(time.Second)
	}

	// Use the MultiPool and set the capacity of the 10 goroutine pools to unlimited.
	// If you use -1 as the pool size parameter, the size will be unlimited.
	// There are two load-balancing algorithms for pools: ants.RoundRobin and ants.LeastTasks.
	mp, _ := ants.NewMultiPool(10, -1, ants.RoundRobin)
	defer mp.ReleaseTimeout(5 * time.Second)

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mp.Submit(task)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", mp.Running())
	fmt.Printf("finish all tasks.\n")
}

func multiPool2() {
	runTimes := 1000
	var wg sync.WaitGroup

	// Use the MultiPoolFunc and set the capacity of 10 goroutine pools to (runTimes/10).
	mpf, _ := ants.NewMultiPoolWithFunc(10, runTimes/10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	}, ants.LeastTasks)

	defer mpf.ReleaseTimeout(5 * time.Second)

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = mpf.Invoke(int32(i))
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", mpf.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
	if sum != 499500 {
		panic("the final result is wrong!!!")
	}
}

func main() {
	//d1()
	//d2()
	//d3()
	//multiPool1()
	multiPool2()
}
