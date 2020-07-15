package main

import (
	"fmt"
	"time"
)

func strings() {
	// 设置一组键值对,并设置有效期
	set1 := client.Set("age", 10, time.Hour*24).Val()
	fmt.Println(set1) // OK

	// 设置一组键值对,设置的键不存在的时候才能设置成功
	set2 := client.SetNX("age", "20", time.Hour*12).Val()
	fmt.Println(set2) // false

	// 设置一组键值对,设置的键必须存在的时候才能设置成功
	set3 := client.SetXX("age", "30", time.Second*86400).Val()
	fmt.Println(set3) // true

	// 批量设置
	set4 := client.MSet("age1", "40", "age2", "50").Val()
	fmt.Println(set4) // OK

	// 获取一个键的值
	get1 := client.Get("age2").Val()
	fmt.Println(get1) // 50

	// 批量获取,获取成功返回slice类型的结果数据
	get2 := client.MGet("age", "age1", "age2").Val()
	fmt.Println(get2) // [30 40 50]

	// 对指定的键进行自增操作
	incr1 := client.Incr("age").Val()
	fmt.Println(incr1) // 31

	// 对指定键进行自减操作
	decr1 := client.Decr("age1").Val()
	fmt.Println(decr1) // 39

	// 自增指定的值
	incr2 := client.IncrBy("age", 10).Val()
	fmt.Println(incr2) // 41

	// 自减指定的值
	decr2 := client.DecrBy("age1", 10).Val()
	fmt.Println(decr2) // 29

	// 在key后面追加指定的值,返回字符串的长度
	append1 := client.Append("age2", "abcd").Val()
	fmt.Println(append1) // 6

	// 获取一个键的值得长度
	strlen1 := client.StrLen("age2").Val()
	fmt.Println(strlen1) //6

	// 设置一个键的值,并返回原有的值
	getset1 := client.GetSet("age2", "hello golang").Val()
	fmt.Println(getset1) // 50abcd

	// 设置键的值,在指定的位置
	_ = client.SetRange("age2", 0, "H")
	fmt.Println(client.Get("age2").Val()) // Hello golang
	// 截取一个键的值的部分,返回截取的部分
	newStr := client.GetRange("age2", 6, 11).Val()
	fmt.Println(newStr) //golang
}
