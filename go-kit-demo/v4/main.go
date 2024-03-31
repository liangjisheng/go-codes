package main

// curl -X POST http://127.0.0.1:8080/login -d '{"account":"alice","password":"123456"}' -H 'content-type: application/json'
// curl -X GET "http://127.0.0.1:8080/sum?a=1&b=1" -H 'Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoibGlhbmdqaXNoZW5nIiwiRGNJRCI6MSwiZXhwIjoxNTgzNDk5NzkwLCJpYXQiOjE1ODM0OTk3NjAsImlzcyI6ImdvLWtpdC12MyIsIm5iZiI6MTU4MzQ5OTc2MCwic3ViIjoibG9naW4ifQ.ZLDVfLs7t3Vt-5p73XgVgflbCW9mDQ8wVef2HVcMhhQ' -H 'content-type: application/json'

import (
	"net/http"

	"v4/utils"
	"v4/v4endpoint"
	"v4/v4service"
	"v4/v4transport"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1) // 每秒产生10个令牌,令牌桶的可以装1个令牌
	uberLimit := ratelimit.New(1)         // 一秒请求一次
	server := v4service.NewService(utils.GetLogger())
	endpoints := v4endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	httpHandler := v4transport.NewHTTPHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 127.0.0.1:8080")
	_ = http.ListenAndServe("127.0.0.1:8080", httpHandler)
}
