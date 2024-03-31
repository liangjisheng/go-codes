package main

import "fmt"

// T ...
type T struct {
	n int
}

// 当使用for-range遍历一个容器时 其实遍历的是此容器的一个副本

func demo3_1() {
	ts := [2]T{}
	for i, t := range ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
			// Output
			// 0
		}
	}
	fmt.Println(ts)
	// Output
	// {0} {9}]
}

func demo3_2() {
	ts := [2]T{}

	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
			// Output
			// 9
		}
	}
	fmt.Println(ts)
	// Output
	// [{0} {9}]
}

func demo3_3() {
	ts := [2]T{}
	for i := range ts[:] {
		t := &ts[i]
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
			// Output
			// 9
		}
	}
	fmt.Println(ts)
	// Output
	// [{3} {9}]
}
