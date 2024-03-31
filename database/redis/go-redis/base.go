package main

import (
	"context"
	"fmt"
	"time"
)

func base() {
	ctx := context.TODO()

	// redis 全局命令
	// 获取redis所有的键,返回包含所有键的slice
	keys := client.Keys(ctx, "*").Val()
	fmt.Println(keys)

	// 获取redis中的有多少个键,返回整数
	size := client.DBSize(ctx).Val()
	fmt.Println(size)

	// 判断一个键是否存在,有一个存在返回整数1,有两个存在返回整数2...
	exist := client.Exists(ctx,"age", "name").Val()
	fmt.Println(exist)

	// 删除键,删除成功返回删除的数,删除失败返回0
	del := client.Del(ctx, "unknownKey").Val()
	fmt.Println(del)

	// 查看键的有效时间
	ttl := client.TTL(ctx, "age").Val()
	fmt.Println(ttl)

	// 给键设置有效时间,设置成功返回true,失败返回false
	expire := client.Expire(ctx, "age", time.Second*86400).Val()
	fmt.Println(expire)

	// 查看键的类型(string,hash,list,set,zset...)
	Rtype := client.Type(ctx, "store:finish:bill:list").Val()
	fmt.Println(Rtype)

	// 给键重命令,成功返回true,失败false
	Rname := client.RenameNX(ctx, "age", "newAge").Val()
	fmt.Println(Rname)

	// 从redis中随机返回一个键
	key := client.RandomKey(ctx, ).Val()
	fmt.Println(key)
}
