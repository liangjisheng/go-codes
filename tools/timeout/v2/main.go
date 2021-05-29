package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func hardWork(job interface{}) error {
	panic("oops")
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// goroutine泄露了，让我们看看为啥会这样呢？首先 requestWork
	// 函数在2秒钟超时后就退出了，一旦 requestWork 函数退出，那么
	// done channel 就没有goroutine接收了，等到执行 done <- hardWork(job)
	// 这行代码的时候就会一直卡着写不进去，导致每个超时的请求都会一直占用掉一个
	// goroutine，这是一个很大的bug，等到资源耗尽的时候整个服务就失去响应了

	done := make(chan error)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <- done:
		return err
	case <- ctx.Done():
		return ctx.Err()
	}
}

func requestWorkFix(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// 把 buffer size 设为1
	// 此时可能有人会问如果这时写入一个已经没goroutine接收的channel会不会有问题
	// 在Go里面channel不像我们常见的文件描述符一样，不是必须关闭的，只是个对象而已
	// close(channel) 只是用来告诉接收者没有东西要写了，没有其它用途

	done := make(chan error, 1)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <- done:
		return err
	case <- ctx.Done():
		return ctx.Err()
	}
}

func requestWorkFixPanicChan(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// 把 buffer size 设为1
	// 此时可能有人会问如果这时写入一个已经没goroutine接收的channel会不会有问题
	// 在Go里面channel不像我们常见的文件描述符一样，不是必须关闭的，只是个对象而已
	// close(channel) 只是用来告诉接收者没有东西要写了，没有其它用途
	done := make(chan error, 1)

	//解决方法是在 requestWork 里加上 panicChan 来处理，同样，需要 panicChan 的 buffer size 为1
	panicChan := make(chan interface{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		done <- hardWork(job)
	}()

	select {
	case err := <- done:
		return err
	case p := <- panicChan:
		panic(p)
	case <- ctx.Done():
		return ctx.Err()
	}
}

func main() {
	const total = 1
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			// 此时执行一下就会发现panic是无法被捕获的，原因是因为在 requestWork
			// 内部起的goroutine里产生的panic其它goroutine无法捕获
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("oops, panic")
				}
			}()

			defer wg.Done()
			//requestWork(context.Background(), "any")
			//requestWorkFix(context.Background(), "any")
			requestWorkFixPanicChan(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(time.Second * 5)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}
