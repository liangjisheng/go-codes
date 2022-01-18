package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

var name = "test"
var number int
var result string

func main() {
	runtime.GOMAXPROCS(1)
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		Timeout:               2000, // 超时时间设置,单位毫秒
		MaxConcurrentRequests: 8,    // 最大请求数
		// 请求阈值 熔断器是否打开首先要满足这个条件;这里的设置表示至少有5个请求才进行
		// ErrorPercentThreshold 错误百分比计算
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  30, // 错误率
		SleepWindow:            1,  // 过多长时间,熔断器再次检测是否开启,单位毫秒
	})

	cbs, _, _ := hystrix.GetCircuit(name)
	defer hystrix.Flush()

	for i := 0; i < 50; i++ {
		start1 := time.Now()
		number = i
		// hystrix.Go(name, run, getFallBack)
		hystrix.Do(name, run, getFallBack)
		time.Sleep(25 * time.Millisecond)
		fmt.Println(
			"请求次数:", i+1,
			";用时:", time.Now().Sub(start1),
			";请求状态:", result,
			";熔断器开启状态:", cbs.IsOpen(),
			"请求是否允许:", cbs.AllowRequest())
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
}

func run() error {
	result = "running"
	if number >= 10 {
		return nil
	}
	if number%2 == 0 {
		return nil
	}
	return errors.New("请求失败")
}

func getFallBack(err error) error {
	result = "fallback"
	return nil
}
