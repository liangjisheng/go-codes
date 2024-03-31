package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	cf := "../ca/server.crt"
	ck := "../ca/server.key"
	log.Println("server start")
	http.ListenAndServeTLS("127.0.0.1:8080", cf, ck, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.RemoteAddr)
	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}
