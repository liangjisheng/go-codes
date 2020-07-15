package main

import (
	"errors"
	"fmt"
	"time"
)

func error1() error {
	// 显示返回一个错误
	return errors.New("something wrong")
}

// 如果你的程序需要记录更多的错误信息，比如时间、数值等信息，
// 可以声明一个自定义的 error 类型

type myerror struct {
	err   string
	time  time.Time
	count int
}

// 实现了error接口
func (m *myerror) Error() string {
	return fmt.Sprintf("%s %d 次。时间：%v", m.err, m.count, m.time)
}

func newErr(s string, i int) *myerror {
	return &myerror{
		err:   s,
		time:  time.Now(),
		count: i,
	}
}

var count int

func error2() error {
	if true {
		count++
		return newErr("self define error", count)
	}
	return nil
}

func selfError() {
	if err := error1(); err != nil {
		fmt.Println(err)
	}
	if err := error2(); err != nil {
		fmt.Println(err)
	}
}
