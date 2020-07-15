package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

var gcount, gerror int
var name = "test"

func init() {
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		Timeout:               100, // cmd 的超时时间,一旦超时则返回失败 毫秒
		MaxConcurrentRequests: 1,   // 最大并发请求数
		// RequestVolumeThreshold 为1,则第二次调用开始探测,如果第一次生成的r随机数
		// 小于1或者大于1但是sleep 200ms,超过上面的Timeout的100ms,引发超时失败,
		// 则由于第二次之前的所有调用(即第一次调用)全部失败(占比100%,总共调用一次,
		// 失败一次),则熔断立马被触发
		RequestVolumeThreshold: 1,   // 熔断探测前的调用次数
		SleepWindow:            100, // 熔断发生后多少毫秒后就开始探测是否恢复
		// 失败次数占总次数的比例,触发熔断之后清空,待恢复之后重新计算
		ErrorPercentThreshold: 50, // 失败占比
	})
}

func testHystrix() error {
	query := func() error {
		var err error
		// rand.Float64()返回[0.0, 1.0),所以r的值肯定小于1,所以乘以2
		rand.Seed(time.Now().Unix())
		r := rand.Float64() * 2
		gcount++
		if r < 1 {
			err = errors.New("bad luck")
			gerror++
			return err
		}
		time.Sleep(200 * time.Millisecond)
		return nil
	}

	var err error
	err = hystrix.Do(name, query, nil)
	// err = hystrix.Go(name, query, nil)
	return err
}

func main() {
	for i := 0; i < 100; i++ {
		err := testHystrix()
		if err != nil {
			log.Printf("testHystrix error:%v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	log.Printf("gcount:%d gerror:%d", gcount, gerror)
}
