package helper

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

// RealIPAddr 获取真实的Ip地址
func RealIPAddr(ctx *fasthttp.RequestCtx) []byte {
	if v := ctx.Request.Header.Peek("X-Forwarded-For"); len(v) > 0 {
		return v
	} else if v := ctx.Request.Header.Peek("X-Real-Ip"); len(v) > 0 {
		return v
	} else {
		return []byte(ctx.RemoteAddr().String())
	}
}

// 协程共享的
var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableCompression:  true,
		MaxIdleConnsPerHost: 50,
	},
	Timeout: time.Duration(30) * time.Second,
}

// Post ...
func Post(url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if req, err = http.NewRequest("POST", url, bytes.NewReader(reqBody)); err != nil {
		return nil, err
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = httpClient.Do(req); err != nil {
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

	if resp, err = httpClient.Do(req); err != nil {
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
