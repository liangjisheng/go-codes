package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

type TestJob struct {
}

func (t TestJob) Run() {
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (t Test2Job) Run() {
	fmt.Println("testJob2...")
}

func main() {
	i := 0
	c := newWithSeconds()
	//spec := "0 */1 * * * ?" // 一分钟运行一次
	spec := "*/5 * * * * ?" // 5s运行一次
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	//AddJob方法
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
