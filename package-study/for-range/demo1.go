package main

import (
	"fmt"
	"time"
)

type user struct {
	id   int
	name string
}

// 测试的时候发现一个有意思的地方，就是go始终利用同一块内存来接收集合中
// 的一个值，只是在每次循环的时候重新赋值而已
func demo1() {
	ii := []int{1, 2, 3, 4, 5}
	for idx, i := range ii {
		fmt.Printf("[%d]: [%d]@ %p --> %p\n", idx, i, &i, &(ii[idx]))
	}

	uu := []user{user{1, "aaa"}, user{2, "bbb"}, user{3, "ccc"}, user{4, "ddd"}}
	for idx, i := range uu {
		// 结论是，&i始终不变！而且，值正常，说明空间是重复利用
		fmt.Printf("[%d]: [%v]@ %p -->%p\n", idx, i, &i, &(uu[idx]))
	}
	fmt.Println()

	var u *user
	go func() {
		var id int
		var name string
		for {
			if nil == u {
				continue
			}

			// 通过这个的变化，来观察是否有清零过程 结论是没有
			if id != u.id {
				fmt.Println("id changed!", id, "to", u.id)
				id = u.id
			}
			// 通过这个的变化，来观察是否有清零过程 结论是没有
			if name != u.name {
				fmt.Println("name changed!", name, "to", u.name)
				name = u.name
			}
		}
	}()

	for idx, i := range uu {
		if idx == 0 {
			u = &i
		}
		// 结论是: &i始终不变！而且值正常 说明空间是复利用
		fmt.Printf("[%d]: [%v]@ %p -->%p\n", idx, i, &i, &(uu[idx]))
		time.Sleep(time.Second)
	}
}
