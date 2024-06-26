package main

import (
	"fmt"
	"github.com/armon/go-radix"
)

func main() {
	r := radix.New()
	r.Insert("foo", 1)
	r.Insert("bar", 2)
	r.Insert("foobar", 2)
	r.Insert("foobartest", "vaule")

	// Find the longest prefix match
	m, _, _ := r.LongestPrefix("foozip")
	if m != "foo" {
		panic("should be foo")
	}

	fmt.Println(r.Len(), r.ToMap())

	r.Walk(func(k string, v interface{}) bool {
		fmt.Println(k)
		return false
	})
}
