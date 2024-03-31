package main

import (
	"fmt"
	"log"
	"net/http"
)

// MyHander ...
type MyHander struct {
}

func (handler *MyHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloGolang(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

// HelloHander ...
type HelloHander struct {
}

func (handler *HelloHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sayHelloGolang(w, r)
}

func sayHelloGolang(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang!")
}

// WorldHander ...
type WorldHander struct {
}

func (handler *WorldHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// curl "http://127.0.0.1:8080/"
// curl "http://127.0.0.1:8081/hello"
// curl "http://127.0.0.1:8081/world"

func main() {
	go func() {
		hello := HelloHander{}
		world := WorldHander{}
		server := http.Server{
			Addr: "127.0.0.1:8081",
		}
		http.Handle("/hello", &hello)
		http.Handle("/world", &world)
		log.Println("http server start on 127.0.0.1:8081")
		server.ListenAndServe()
	}()

	handler := MyHander{}
	log.Println("http server start on 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", &handler)
}
