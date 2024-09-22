package main

import (
	"fmt"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

func fileOnly() {
	var dbPath = "ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	// do the search
	// 韩国|0|首尔|首尔|0
	//var ip = "58.140.41.6"

	var ip = "117.129.58.230"

	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Println("region:", region)
	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))

	// 备注：并发使用，每个 goroutine 需要创建一个独立的 searcher 对象。
}

func vectorIndex() {
	var dbPath = "ip2region.xdb"
	// 1、从 dbPath 加载 VectorIndex 缓存，把下述 vIndex 变量全局到内存里面。
	vIndex, err := xdb.LoadVectorIndexFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load vector index from `%s`: %s\n", dbPath, err)
		return
	}

	// 2、用全局的 vIndex 创建带 VectorIndex 缓存的查询对象。
	searcher, err := xdb.NewWithVectorIndex(dbPath, vIndex)
	if err != nil {
		fmt.Printf("failed to create searcher with vector index: %s\n", err)
		return
	}

	// do the search
	var ip = "58.140.41.6"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Println("region:", region)
	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))

	// 备注：并发使用，全部 goroutine 共享全局的只读 vIndex 缓存，每个 goroutine 创建一个独立的 searcher 对象
}

func memorySearch() {
	var dbPath = "ip2region.xdb"
	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
		return
	}

	// do the search
	//韩国|0|首尔|首尔|0
	//var ip = "58.140.41.6"

	//中国|0|北京|北京市|移动
	//var ip = "117.129.58.230"

	//中国|0|香港|0|谷歌
	var ip = "35.220.164.227"

	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return
	}

	fmt.Println("region:", region)
	fmt.Printf("{region: %s, took: %s}\n", region, time.Since(tStart))

	// 备注：并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。
}

func main() {
	//fileOnly()
	//vectorIndex()
	memorySearch()
}
