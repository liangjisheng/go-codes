package main

import (
	"fmt"
	"io/ioutil"
)

func readDir() {
	rd, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", fi.Name())
		} else {
			fmt.Println(fi.Name())
		}
	}
}
