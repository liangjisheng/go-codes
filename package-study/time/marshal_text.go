package main

import (
	"fmt"
	"time"
)

func marshalText() {
	// 可见Marshal类的函数只是提供一个将时间t序列化为[]byte数组的功能，
	// 利用UnMarshal类的函数可以获取到原来的时间t
	a := time.Now()
	fmt.Println(a)
	b, _ := a.MarshalText()
	fmt.Println(b)
	var c = new(time.Time)
	fmt.Println(c)
	c.UnmarshalText(b)
	fmt.Println(c)
	fmt.Println()

	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	// func (t Time) Round(d Duration) Time
	// 将时间ｔ以d Duration为单位进行四舍五入求近似值
	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
	fmt.Println()

	// 截断
	t, _ = time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")
	trunc := round
	for _, d := range trunc {
		fmt.Printf("t.Truncate(%6s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
}
