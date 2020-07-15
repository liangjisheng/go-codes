package main

import (
	"fmt"
	"regexp"
)

func regexp2() {
	// regular expression pattern
	re := regexp.MustCompile("/oid/([\\d]+)/")

	// first convert string to byte for Find() function
	searchByte := []byte("/oid/1/")
	matchSlice := re.Find(searchByte)
	// if found, return leftmost match
	fmt.Printf("%s\n", matchSlice) // /oid/1

	matchSlice2 := re.FindAll(searchByte, 500)
	// if found, return all successive matches
	fmt.Printf("%s\n", matchSlice2)

	oid := re.NumSubexp()
	fmt.Printf("OID is %d\n", oid)

	// search by string
	matchSlice3 := re.FindAllString(string(searchByte), -1)
	fmt.Printf("%s\n", matchSlice3)
}
