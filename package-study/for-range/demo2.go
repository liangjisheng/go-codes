package main

import (
	"fmt"
	"time"
)

func demo2_0() {
	str := []string{"I", "like", "Golang"}
	for _, v := range str {
		v += "good"
	}
	for k, v := range str {
		fmt.Println(k, v)
	}
}

func demo2_1() {
	str := []string{"I", "like", "Golang"}
	for k, v := range str {
		fmt.Println(&k, &v)
	}
}

func demo2_2() {
	str := []string{"I", "like", "Golang"}
	for k, v := range str {
		str = append(str, "good")
		fmt.Println(str)
		fmt.Println(k, v)
	}
	fmt.Println(str)
}

func demo2_3() {
	str := []string{"I", "like", "Golang"}
	for k, v := range str {
		go func(i int, s string) {
			fmt.Println(i, s, k, v)
		}(k, v)
	}
	time.Sleep(1e9)
}

func demo2_4() {
	str := []string{"I", "like", "Golang"}
	for k, v := range str {
		go func(i int, s string) {
			fmt.Println(i, s, k, v)
		}(k, v)
		time.Sleep(1e9)
	}
	time.Sleep(5e9)
}
