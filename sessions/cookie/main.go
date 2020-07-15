package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	cookieName = "my-cookie"
)

func main() {
	// 在浏览器上测试
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	log.Println("server start.")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: "some value",
	})
	fmt.Fprintf(w, "COOKIE WRITTEN - CHECK YOUR BROWSER\n")
}

func get(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "YOUR COOKIE: %+v\n", c)
}
