package main

import (
	"sync"
)

type Token struct {
	ID int64
}

type Overview struct {
	Count int64
}

type Cache struct {
	Tokens    []*Token
	TokensMap map[string]*Token
	Overview  *Overview
}

var (
	data = Cache{
		Tokens:    make([]*Token, 0),
		TokensMap: make(map[string]*Token, 0),
		Overview:  nil,
	}

	dataMap = map[string]string{}

	rwMutex = sync.RWMutex{}
)

func main() {
	go func() {
		var count int64
		for {
			count++
			rwMutex.Lock()
			tokens := []*Token{
				{ID: 1},
				{ID: 2},
			}

			tokensMap := map[string]*Token{
				"1": {ID: 1},
				"2": {ID: 2},
			}

			overview := &Overview{
				Count: 1,
			}

			data.Tokens = tokens
			data.TokensMap = tokensMap
			data.Overview = overview

			//直接修改会导致下面读的时候报错
			//dataMap["alice"] = "alice"

			//重新赋值则不会报错
			//dataMap = map[string]string{
			//	"alice": "alice",
			//}

			//直接修改会导致下面读的时候报错
			//data.TokensMap["alice"] = &Token{
			//	ID: 3,
			//}

			//t1Map := map[string]*Token{
			//	"3": {ID: 3},
			//	"4": {ID: 4},
			//}
			//这样修改也会报错
			//for k, v := range t1Map {
			//	data.TokensMap[k] = v
			//}

			//重新赋值则不会报错
			data.TokensMap = map[string]*Token{
				"3": {ID: 3},
				"4": {ID: 4},
			}

			if count > 0 && count%100 == 0 {
				//fmt.Println("write", count)
			}

			rwMutex.Unlock()
			//time.Sleep(10 * time.Millisecond)
		}
	}()

	var count int64
	for {
		count++
		//rwMutex.RLock()
		for _, _ = range data.TokensMap {
			if count > 0 && count%100 == 0 {
				//fmt.Println("read", count)
			}
			//time.Sleep(10 * time.Millisecond)
		}
		//rwMutex.RUnlock()

		//不加读锁的话
		//如果直接修改 map, 则程序会直接 panic 报下面的错误
		//fatal error: concurrent map iteration and map write
		//如果将一个新的对象赋值给 map, 则不会报错

		//加读锁的话，无论是直接修改还是赋值新对象都不会报错
	}
}
