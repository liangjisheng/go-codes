package tokenbucket_test

import (
	"github.com/gin-gonic/gin"
	tokenbucket "go-demos/tools/rate-limit/token-bucket-2"
	"log"
	"sync"
	"testing"

	"io/ioutil"
	"net/http"
	"time"
)

func TestLimit(t *testing.T) {
	go func() {
		route := gin.Default()

		route.GET("/ping",
			// 一开始桶是空的，然后以每秒1个token的速度放入桶中, 一共可以放2个token
			tokenbucket.LimitMiddleware(1, 2),
			func (c *gin.Context) {
				c.String(http.StatusOK, "pong")
			},
		)

		route.Run("127.0.0.1:8080")
	}()

	//这里需要 sleep
	time.Sleep(2 * time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resp, err := http.Get("http://127.0.0.1:8080/ping")
			if err != nil {
				log.Println("err", err)
				return
			}
			defer resp.Body.Close()
			respBody, _ := ioutil.ReadAll(resp.Body)
			log.Println("req ", i, string(respBody))
		}(i)
	}

	wg.Wait()
}

func TestLimit2(t *testing.T) {
	var wg sync.WaitGroup
	var lr tokenbucket.TokenBucket
	lr.Set(3, 5)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		log.Println("创建请求:", i)
		go func(i int) {
			if lr.Allow() {
				log.Println("响应请求:", i)
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
