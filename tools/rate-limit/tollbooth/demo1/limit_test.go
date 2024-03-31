package main_test

import (
	"github.com/didip/tollbooth/v6"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resp, err := http.Get("http://127.0.0.1:8080/")
			if err != nil {
				log.Println("err", err)
				return
			}
			defer resp.Body.Close()
			respBody, _ := ioutil.ReadAll(resp.Body)
			log.Println("req ", i, string(respBody))
		}(i)
	}

	wg.Wait()
}

func TestLimitByKeys(t *testing.T) {
	lmt := tollbooth.NewLimiter(1, nil) // Only 1 request per second is allowed.

	httperror := tollbooth.LimitByKeys(lmt, []string{"127.0.0.1", "/"})
	if httperror != nil {
		t.Errorf("First time count should not return error. Error: %v", httperror.Error())
	}

	httperror = tollbooth.LimitByKeys(lmt, []string{"127.0.0.1", "/"})
	if httperror == nil {
		t.Errorf("Second time count should return error because it exceeds 1 request per second.")
	}

	<-time.After(1 * time.Second)
	httperror = tollbooth.LimitByKeys(lmt, []string{"127.0.0.1", "/"})
	if httperror != nil {
		t.Errorf("Third time count should not return error because the 1 second window has passed.")
	}
}

func TestDefaultBuildKeys(t *testing.T) {
	lmt := tollbooth.NewLimiter(1, nil)

	request, err := http.NewRequest("GET", "/", strings.NewReader("Hello, world!"))
	if err != nil {
		t.Errorf("Unable to create new HTTP request. Error: %v", err)
	}

	request.Header.Set("X-Real-IP", "2601:7:1c82:4097:59a0:a80b:2841:b8c8")

	sliceKeys := tollbooth.BuildKeys(lmt, request)
	if len(sliceKeys) == 0 {
		t.Fatal("Length of sliceKeys should never be empty.")
	}

	for _, keys := range sliceKeys {
		expectedKeys := [][]string{
			{request.Header.Get("X-Real-IP")},
			{request.URL.Path},
		}

		checkKeys(t, keys, expectedKeys)
	}
}

func checkKeys(t *testing.T, keys []string, expectedKeys [][]string) {
	for i, keyChunk := range keys {
		switch {
		case i == 0 && !isInSlice(keyChunk, expectedKeys[0]):
			t.Errorf("The (%v) chunk should be remote IP. KeyChunk: %v", i+1, keyChunk)
		case i == 1 && !isInSlice(keyChunk, expectedKeys[1]):
			t.Errorf("The (%v) chunk should be request path. KeyChunk: %v", i+1, keyChunk)
		}
	}

	for _, ekeys := range expectedKeys {
		found := false
		for _, ekey := range ekeys {
			for _, key := range keys {
				if ekey == key {
					found = true
					break
				}
			}
		}

		if !found {
			t.Fatalf("expectedKeys missing: %v", strings.Join(ekeys, " "))
		}
	}
}

func isInSlice(key string, keys []string) bool {
	for _, sliceKey := range keys {
		if key == sliceKey {
			return true
		}
	}
	return false
}
