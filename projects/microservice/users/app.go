package main

import (
	"go-demos/microservice/users/routes"
	"net/http"
)

func main() {
	r := routes.NewRouter()
	http.ListenAndServe(":8000", r)
}
