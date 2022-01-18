package main

import (
	"github.com/didip/tollbooth/v6"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	lmt := tollbooth.NewLimiter(2, nil)
	http.Handle("/", tollbooth.LimitFuncHandler(lmt, HelloHandler))
	log.Println("serer listen on :8080")
	http.ListenAndServe(":8080", nil)
}
