package main

// curl "http://127.0.0.1:8080/sum?a=1&b=1"

import (
	"net/http"

	"v2/utils"
	"v2/v2endpoint"
	"v2/v2service"
	"v2/v2transport"
)

func main() {
	utils.NewLoggerServer()
	server := v2service.NewService(utils.GetLogger())
	endpoints := v2endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v2transport.NewHTTPHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 127.0.0.1:8080")
	_ = http.ListenAndServe("127.0.0.1:8080", httpHandler)
}
