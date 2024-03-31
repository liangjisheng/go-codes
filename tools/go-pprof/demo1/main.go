package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"go-demos/tools/go-pprof/demo1/data"
)

func main() {
	go func() {
		log.Println(data.Add("https://github.com/liangjisheng"))
	}()

	log.Println("start server")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
