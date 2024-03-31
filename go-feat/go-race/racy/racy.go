package main

import "fmt"

func main() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"

	go func() {
		m["name"] = "data race"
		done <- true
	}()

	fmt.Println("Hello,", m["name"])
	<-done
}

//Then run it with the race detector enabled:
//$ go run -race racy.go
