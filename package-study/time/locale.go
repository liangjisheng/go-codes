package main

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func locale1() {
	locales = make(map[string]map[string]string)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	cn := make(map[string]string)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	en["how old"] = "I am %d years old"
	cn["how old"] = "我今年%d岁了"
	fmt.Printf(msg(lang, "how old"), 30)
	fmt.Println()

	// 为了获得对应于当前locale的时间，我们应首先使用time.LoadLocation(name string)
	// 获取相应于地区的locale，比如Asia/Shanghai或America/Chicago对应的时区信息，
	// 然后再利用此信息与调用time.Now获得的Time对象协作来获得最终的时间
	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"
	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

func locale2() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UTC().Unix())

	currentTime := time.Now().UTC()
	fmt.Println(time.Date(currentTime.Year(),
		currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(),
		currentTime.Second(), currentTime.Nanosecond(),
		currentTime.Location()).Unix())

	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//loc, _ := time.LoadLocation("America/Chicago")
	//loc, _ := time.LoadLocation("Asia/Jerusalem")
	loc, _ := time.LoadLocation("UTC")
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339), t.Unix())
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339), t.Unix())
}
