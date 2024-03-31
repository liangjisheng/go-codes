package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
)

func TestJustError(t *testing.T) {
	g := new(errgroup.Group)
	var urls = []string{
		"https://www.golang.org/",
		"https://www.google.com/",
		"https://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
