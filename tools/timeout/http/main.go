package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)

	if err != nil {
		fmt.Println("http new request failed, err: ", err)
		return
	}

	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()

	select {
	case <- ctx.Done():
		tr.CancelRequest(req)
		res := <- c
		fmt.Printf("Timeout, res= %v\n", res)
	case res := <- c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s\n", out)
	}
}

func main() {
	process()
}
