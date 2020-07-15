package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		helloTo := r.FormValue("helloTo")
		helloStr := fmt.Sprintf("hello, %s", helloTo)
		w.Write([]byte(helloStr))
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
