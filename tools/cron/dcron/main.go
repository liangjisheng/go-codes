package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/libi/dcron"
	dredis "github.com/libi/dcron/driver/redis"
	"github.com/robfig/cron/v3"
	"time"
)

//type TestJob1 struct {
//	Name string
//}
//
//func (t TestJob1) Run() {
//	fmt.Println("执行 testjob ", t.Name, time.Now().Format("15:04:05"))
//}

//var testData = make(map[string]struct{})

func main() {
	drv, _ := dredis.NewDriver(&dredis.Conf{
		Host: "127.0.0.1",
		Port: 6379,
	}, redis.DialConnectTimeout(time.Second*10))

	//add recover
	dcron1 := dcron.NewDcron("server1", drv, cron.WithChain(cron.Recover(cron.DefaultLogger)))
	fmt.Println("new dcron ok")

	//panic recover test
	err := dcron1.AddFunc("s1 test1", "* * * * *", func() {
		panic("panic test")
		fmt.Println("执行 service1 test1 任务,模拟 panic", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Errorf(" s1 test1 err %+v", err)
	}

	err = dcron1.AddFunc("s1 test2", "* * * * *", func() {
		fmt.Println("执行 service1 test2 任务", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Println("add func error")
	}

	err = dcron1.AddFunc("s1 test3", "* * * * *", func() {
		fmt.Println("执行 service1 test3 任务", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Errorf(" s1 test3 err %+v", err)
	}
	dcron1.Start()

	//测试120秒后退出
	//time.Sdleep(120 * time.Second)
	select {}
}
