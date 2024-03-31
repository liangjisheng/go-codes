package main

import (
	"fmt"
	"time"
)

func timeStamp() {
	nowTime := time.Now().Unix()
	fmt.Println("current timestamp:", nowTime)
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 时间截转str格式化时间
	strTime := time.Unix(nowTime, 0).Format("2006-01-02 15:04:05")
	fmt.Println(strTime)

	// str格式化时间转换为时间截
	theTime := time.Date(2018, 8, 22, 20, 06, 20, 0, time.Local)
	unixTime := theTime.Unix()
	fmt.Println(unixTime)

	// 返回固定时间的时间截
	t1, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err != nil {
		fmt.Println("parse time err:", err)
		return
	}
	unixTime = t1.Unix()
	fmt.Println(unixTime)
}
