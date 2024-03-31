package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHTTPDo() {
	url := "http://localhost:9090/"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpGet() {
	url := "http://localhost:9090/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//var httpClient = &http.Client{
//	Transport: &http.Transport{
//		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
//		DisableCompression:  true,
//		MaxIdleConnsPerHost: 50,
//	},
//	Timeout: time.Duration(30) * time.Second,
//}

// Get ...
func Get(url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}

	if len(data) > 0 {
		query := req.URL.Query()
		for k, v := range data {
			query.Add(k, fmt.Sprint(v))
		}

		req.URL.RawQuery = query.Encode()
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = http.DefaultClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ErrCode:%d, with %s => %#v", resp.StatusCode, url, data)
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func main() {
	// getHTTPDo()
	httpGet()
}
