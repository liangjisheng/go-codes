package main

import (
	"net/http"

	"go-demos/projects/restful/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":9090", r)
}
