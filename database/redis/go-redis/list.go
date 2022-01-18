package main

import (
	"context"
	"fmt"
	"strconv"
)

func list() {
	ctx := context.TODO()

	key := "demo"
	client.Del(ctx, key)
	for i := 0; i < 5; i++ {
		client.LPush(ctx, key, "e-"+strconv.Itoa(i))
	}

	// 获取list 长度
	length := client.LLen(ctx, key).Val()
	fmt.Println(length) // 5
	// 获取指定下标元素
	value1 := client.LIndex(ctx, key, 0).Val()
	fmt.Println(value1) // e-4
	// 获取所有元素
	vs := client.LRange(ctx, key, 0, -1).Val()
	fmt.Println(vs) // [e-4 e-3 e-2 e-1 e-0]
	// 修改指定下标的元素值
	status := client.LSet(ctx, key, 0, "golang").Val()
	fmt.Println(status) // ok
	// 从右边插入元素
	client.RPush(ctx, key, "e-right")
	// 从左边插入元素
	client.LPush(ctx, key, "e-left")
	// 从list最右边弹出元素
	v1 := client.RPop(ctx, key).Val()
	fmt.Println(v1) // e-right
	// 从list最左边弹出元素
	v2 := client.LPop(ctx, key).Val()
	fmt.Println(v2) // e-left
	// 删除指定元素
	n := client.LRem(ctx, key, 0, "e-3").Val()
	fmt.Println(n) // 1
	status1 := client.LTrim(ctx, key, 0, 1).Val()
	fmt.Println(status1) // ok
}
