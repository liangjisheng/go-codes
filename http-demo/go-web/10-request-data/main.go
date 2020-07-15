package main

import (
	"log"
	"net/http"

	"requestdata/routes"
)

func main() {
	r := routes.NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at 127.0.0.1:8080")
	err := http.ListenAndServe("127.0.0.1:8080", nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
