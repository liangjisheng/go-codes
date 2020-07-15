package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

/*
 爬虫思路:
 1）明确目标（要知道你准备在哪个范围或网站去搜索）
 2）爬（将所有的网站的内容全部爬下来）
 3）取（去掉对我们没用处的数据）
 4）处理数据（按照我们想要的方式存储和使用）
 百度贴吧爬虫
*/

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get err=", err1)
		return
	}
	defer resp.Body.Close()

	// 读取网页Body中的内容
	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err =", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)
	//明确目标（要知道准备在哪个范围或网站去搜索）
	//下一页+50
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		//爬（将所有的网站的内容全部爬下来）
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err =", err)
			continue
		}

		// 把爬下来的内容写到文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create err =", err)
			continue
		}
		f.WriteString(result)
		f.Close()
	}
}

func crawler1() {
	var start, end int
	fmt.Printf("请输入起始页(大于等于1): ")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(大于等于起始页): ")
	fmt.Scan(&end)
	DoWork(start, end)
}
