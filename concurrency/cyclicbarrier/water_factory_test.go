package cyclicbarrier

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestWaterFactory(t *testing.T) {
	// 存放水分子结果。
	var ch chan string
	releaseHydrogen := func() {
		ch <- "H"
	}
	releaseOxygen := func() {
		ch <- "O"
	}

	// 每个 goroutine 并发产生一个原子。
	var N = 100
	ch = make(chan string, N*3)
	h2o := New()

	// 等待所有 goroutine 完成。
	var wg sync.WaitGroup
	wg.Add(N * 3)

	// 氢原子 goroutine。
	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}
	// 氧原子 goroutine。
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}

	//等待所有的 goroutine 执行完
	wg.Wait()
	if len(ch) != N*3 {
		t.Fatalf("expect %d atom but got %d", N*3, len(ch))
	}

	// 分组检查结果。
	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)
		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			t.Fatalf("expect a water molecule but got %s", water)
		}
	}
}
