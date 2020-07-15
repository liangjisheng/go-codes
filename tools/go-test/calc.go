package calc

import "net/http"

// Add ...
func Add(a int, b int) int {
	return a + b
}

// Mul ...
func Mul(a int, b int) int {
	return a * b
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
