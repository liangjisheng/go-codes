package main

import (
	"fmt"
	"reflect"
)

// 发送数据（会阻塞），v 值必须是可写通道。
// func (v Value) Send(x reflect.Value)

// 接收数据（会阻塞），v 值必须是可读通道。
// func (v Value) Recv() (x reflect.Value, ok bool)

// 尝试发送数据（不会阻塞），v 值必须是可写通道。
// func (v Value) TrySend(x reflect.Value) bool

// 尝试接收数据（不会阻塞），v 值必须是可读通道。
// func (v Value) TryRecv() (x reflect.Value, ok bool)

// 关闭通道，v 值必须是通道。
// func (v Value) Close()

func reflect5() {
	ch := make(chan int, 2)
	v := reflect.ValueOf(ch)

	a := reflect.ValueOf(1)
	b := reflect.ValueOf(2)

	v.Send(a)
	if ok := v.TrySend(b); ok {
		fmt.Println("尝试发送成功！") // 尝试发送成功！
	}
	if i, ok := v.Recv(); ok {
		fmt.Println("接收成功：", i) // 接收成功： 1
	}
	if i, ok := v.TryRecv(); ok {
		fmt.Println("尝试接收成功：", i) // 尝试接收成功： 2
	}
}
