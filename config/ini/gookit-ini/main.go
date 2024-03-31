package main

import (
	"fmt"
	"github.com/gookit/ini/v2"
)

func main() {
	// config, err := ini.LoadFiles("testdata/tesdt.ini")
	// LoadExists will ignore not exists file
	err := ini.LoadExists("testdata/test.ini", "not-exist.ini")
	if err != nil {
		panic(err)
	}

	// load more, will override prev data by key
	err = ini.LoadStrings(`
age = 100
[sec1]
newK = newVal
some = change val
`)
	fmt.Printf("%v\n", ini.Data())
}
