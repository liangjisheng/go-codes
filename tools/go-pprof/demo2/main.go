package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	case1()
	// case2()
}

// 运行这个函数后, cpu,内存占用很高, 说明发生了内存泄露
func case1() {
	ch := make(chan string, 100)

	go func() {
		for {
			ch <- "alice"
		}
	}()

	go func() {
		ip := "127.0.0.1:8080"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()

	duration := 3 * time.Minute
	timeout := time.NewTimer(duration)
	defer timeout.Stop()

	for {
		timeout.Reset(duration)
		select {
		case <-ch:
		// 这里我们的定时时间设置的是3分钟, 在for循环每次select的时候, 都会实例化
		// 一个一个新的定时器. 该定时器在3分钟后, 才会被激活, 但是激活后已经跟select无引用关系
		// 被gc给清理掉. 这里最关键的一点是在计时器触发之前, 垃圾收集器不会回收 Timer
		// 换句话说, 被遗弃的time.After定时任务还是在时间堆里面, 定时任务未到期之前
		// 是不会被gc清理的, 所以这就是会造成内存泄漏的原因. 每次循环实例化的新定时器对象
		// 需要3分钟才会可能被GC清理掉, 如果我们把上面代码中的3分钟改小点, 会有所改善, 但是仍存在风险
		// case <-time.After(time.Minute * 3):
		case <-timeout.C:
		}
	}
}

func case2() {
	// 开启pprof，监听请求
	ip := "127.0.0.1:8080"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
