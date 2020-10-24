package main

import (
	v24 "github.com/google/go-github/v24/github"
	v25 "github.com/google/go-github/v25/github"
	"golang.org/x/text/width"
)

var (
	_ = v24.Tag{}
	_ = v25.Tag{}
	_ = width.EastAsianAmbiguous
)

func main() {
	return
}
