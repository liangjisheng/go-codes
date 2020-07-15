package main

import (
	"go-demos/microservice/showtimes/routes"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8002", r)
}
