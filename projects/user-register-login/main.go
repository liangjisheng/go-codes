package main

import (
	"fmt"
	"net/http"

	webservice "go-demos/projects/user-register-login/webService"
)

func main() {
	fmt.Println("----> go start <----")
	mux := &webservice.CustomMux{}
	http.ListenAndServe("localhost:8080", mux)
}
