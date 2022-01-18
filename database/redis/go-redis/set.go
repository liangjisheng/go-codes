package main

import (
	"context"
	"fmt"
	"strconv"
)

func set() {
	ctx := context.TODO()

	key := "sdemo"
	key1 := "sdemo1"
	client.Del(ctx, key)
	client.Del(ctx, key1)
	for i := 0; i < 6; i++ {
		// set 类型添加元素
		client.SAdd(ctx, key, "ele-"+strconv.Itoa(i))
	}
	for i := 3; i < 9; i++ {
		// set 类型添加元素
		client.SAdd(ctx, key1, "ele-"+strconv.Itoa(i))
	}

	// 计算key中的元素个数
	n1 := client.SCard(ctx, key).Val()
	fmt.Println(n1) // 5
	// 判断元素是否在集合中
	e1 := client.SIsMember(ctx, key, "ele-0").Val()
	fmt.Println(e1) // true
	// 随机从集合中返回一个元素
	v1 := client.SRandMember(ctx, key).Val()
	fmt.Println(v1)
	// 随机返回指定个数的元素,返回包含元素的slice
	v2 := client.SRandMemberN(ctx, key, 3).Val()
	fmt.Println(v2)
	// 获取集合中的所有元素,无序的slice
	v3 := client.SMembers(ctx, key).Val()
	fmt.Println(v3) // [ele-1 ele-0 ele-3 ele-2 ele-4]
	// 从集合中随机弹出一个元素
	v4 := client.SPop(ctx, key).Val()
	fmt.Println(v4)
	// 从集合中删除元素
	n2 := client.SRem(ctx, key, "ele-5").Val()
	fmt.Println(n2) // 1
	// 求多个集合的交集
	s1 := client.SInter(ctx, key, key1).Val()
	fmt.Println(s1) // [ele-3 ele-4]
	// 求多个集合的并集
	s2 := client.SUnion(ctx, key, key1).Val()
	fmt.Println(s2) // [ele-3 ele-5 ele-1 ele-2 ele-4 ele-8 ele-6 ele-0 ele-7]
	// 求多个集合的差集
	s3 := client.SDiff(ctx, key, key1).Val()
	fmt.Println(s3) // [ele-0 ele-1]
	// 将多个交集结果存为一个新的集合
	s4 := client.SInterStore(ctx, "sdemo2", key, key1).Val()
	fmt.Println(s4)                              // 2
	fmt.Println(client.SMembers(ctx, "sdemo2").Val()) // [ele-4 ele-3]
	// 将多个交集的并集结果存为新的集合
	s5 := client.SUnionStore(ctx, "sdemo3", key, key1).Val()
	fmt.Println(s5)                              // 8
	fmt.Println(client.SMembers(ctx, "sdemo3").Val()) // [ele-3 ele-5 ele-1 ele-4 ele-7 ele-8 ele-6 ele-0]
	// 将多个差集的结果存为新的集合
	s6 := client.SDiffStore(ctx, "sdemo4", key, key1).Val()
	fmt.Println(s6)                              // 2
	fmt.Println(client.SMembers(ctx, "sdemo4").Val()) // [ele-0 ele-1]
}
