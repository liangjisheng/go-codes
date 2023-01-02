package snowflake

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSnowFlack(t *testing.T) {
	w := NewWorker(1, 1)
	ch := make(chan uint64, 1000)
	count := 5
	wg.Add(count)
	defer close(ch)

	// 并发 count个goroutine 进行 snowFlake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < count; i++ {
				id, _ := w.NextID()
				log.Println("id", id)
				ch <- id
			}
		}()
	}

	go func() {
		m := make(map[uint64]int)
		for i := 0; i < count*count; i++ {
			id := <-ch
			// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
			_, ok := m[id]
			if ok {
				fmt.Printf("repeat id %d\n", id)
				return
			}
			// 将 id 作为 key 存入 map
			m[id] = i
		}

		// 成功生成 snowflake ID
		fmt.Println("All", len(m), "snowflake ID Get successed!")
	}()

	wg.Wait()
}
