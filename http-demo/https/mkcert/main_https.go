package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	key := "127.0.0.1+1-key.pem"
	cert := "127.0.0.1+1.pem"

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world\n")
	})
	log.Fatal(http.ListenAndServeTLS(":8000", cert, key, nil))
}
