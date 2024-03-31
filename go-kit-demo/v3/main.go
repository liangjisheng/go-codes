package main

// curl -X POST http://127.0.0.1:8080/login -d '{"account":"alice","password":"123456"}' -H 'content-type: application/json'
// curl -X GET "http://127.0.0.1:8080/sum?a=1&b=1" -H 'Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoibGlhbmdqaXNoZW5nIiwiRGNJRCI6MSwiZXhwIjoxNTgzNDk5NzkwLCJpYXQiOjE1ODM0OTk3NjAsImlzcyI6ImdvLWtpdC12MyIsIm5iZiI6MTU4MzQ5OTc2MCwic3ViIjoibG9naW4ifQ.ZLDVfLs7t3Vt-5p73XgVgflbCW9mDQ8wVef2HVcMhhQ' -H 'content-type: application/json'

import (
	"net/http"
	"v3/utils"
	"v3/v3endpoint"
	"v3/v3service"
	"v3/v3transport"
)

func main() {
	utils.NewLoggerServer()
	server := v3service.NewService(utils.GetLogger())
	endpoints := v3endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v3transport.NewHTTPHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 127.0.0.1:8080")
	_ = http.ListenAndServe("127.0.0.1:8080", httpHandler)
}
