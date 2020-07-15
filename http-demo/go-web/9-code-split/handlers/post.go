package handlers

import (
	"io"
	"net/http"
)

// GetPosts ...
func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}
