package main

import (
	"go-demos/microservice/movies/routes"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8001", r)
}
