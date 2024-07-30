package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
)

func demo() {
	m := map[string]string{
		"alice": "alice",
	}
	m = nil

	res, ok := m["alice"]
	if !ok {
		fmt.Println("not ok")
		return
	}
	fmt.Println(res)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}

var (
	sMap = sync.Map{}
)

func demo1() {
	key := "key"
	value := []string{"a", "b"}

	sMap.Store(key, value)

	v1I, ok := sMap.Load(key)
	if !ok {
		return
	}

	v1, ok := v1I.([]string)
	if !ok {
		return
	}
	fmt.Println(v1)

	v1 = append(v1, "c")
	sMap.Store(key, v1)

	v1I, ok = sMap.Load(key)
	if !ok {
		return
	}

	v1, ok = v1I.([]string)
	if !ok {
		return
	}
	fmt.Println(v1)
}

func testU128() {
	str := "ffffffffffffffffffffffffffffffff"
	value, ok := new(big.Int).SetString(str, 16)
	if !ok {
		fmt.Println("not ok")
		return
	}
	fmt.Println(hex.EncodeToString(value.Bytes()))
	fmt.Println(value.String())
}

func main() {
	testU128()
}
