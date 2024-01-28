package main

import (
	"fmt"
	"time"

	"github.com/xxjwxc/gowp/workpool"
)

func test1() {
	wp := workpool.New(5) //设置最大线程数
	fmt.Println(wp.IsDone())
	wp.DoWait(func() error {
		for j := 0; j < 10; j++ {
			fmt.Println(fmt.Sprintf("%v->\t%v", 000, j))
		}

		return nil
		// time.Sleep(1 * time.Second)
		// return errors.New("my test err")
	})
}

func test2() {
	wp := workpool.New(5) //设置最大线程数
	for i := 0; i < 50; i++ {
		ii := i
		wp.Do(func() error {
			for j := 0; j < 1; j++ {
				time.Sleep(time.Millisecond * 500)
				fmt.Println(fmt.Sprintf("%v->\t%v", ii, j))
				// if ii == 1 {
				// 	return errors.Cause(errors.New("my test err")) //有err 立即返回
				// }
			}
			return nil
		})

		//fmt.Println("is done", wp.IsDone())
	}

	wp.Wait()
	fmt.Println(wp.IsDone())
	//fmt.Println(wp.IsClosed())
	//fmt.Println("down")
}

func main() {
	//test1()
	test2()
}
