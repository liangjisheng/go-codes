package main

import (
	"fmt"
	"time"
)

func show() {
	layout := "01__02-2006 3.04.05 PM" // 指定自己的时间格式
	fmt.Println(time.Now().Format(layout))
	fmt.Println(time.Now().Format(time.ANSIC))
	fmt.Println(time.Now().Format(time.UnixDate))
	fmt.Println(time.Now().Format(time.RubyDate))
	fmt.Println(time.Now().Format(time.RFC822))
	fmt.Println(time.Now().Format(time.RFC822Z))
	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(time.Now().Format(time.RFC1123))
	fmt.Println(time.Now().Format(time.RFC1123Z))
	fmt.Println(time.Now().Format(time.RFC3339Nano))
	fmt.Println(time.Now().Format(time.Kitchen))
	fmt.Println()

	// Handy time stamps
	fmt.Println(time.Now().Format(time.Stamp))
	fmt.Println(time.Now().Format(time.StampMilli))
	fmt.Println(time.Now().Format(time.StampMicro))
	fmt.Println(time.Now().Format(time.StampNano))
}
