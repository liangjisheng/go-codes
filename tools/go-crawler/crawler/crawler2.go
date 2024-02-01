package main

import (
	"fmt"
	"os"
	"strconv"
)

// 并发百度贴吧爬虫

//爬取一个网页
func SpiderPage(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	//爬（将所有的网站的内容全部爬下来）
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err ", err)
		return
	}

	//把内容写入到文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err=", err1)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}

func DoWork2(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	page := make(chan int)
	//明确目标（要知道准备在哪个范围或网站去搜索）
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬去完成\n", <-page)
	}
}

func crawler2() {
	var start, end int
	fmt.Printf("请输入起始页(大于等于1): ")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(大于等于起始页): ")
	fmt.Scan(&end)
	DoWork2(start, end)
}
