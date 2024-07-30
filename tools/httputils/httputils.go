package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	gClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        5000,
			MaxIdleConnsPerHost: 1000,
			MaxConnsPerHost:     1000,
		},
		Timeout: time.Second * 30,
	}
)

func RequestGet(url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
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
			query.Set(k, fmt.Sprint(v))
		}

		req.URL.RawQuery = query.Encode()
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get failed, status=%d, resp=%+v, url=%v, query=%v", resp.StatusCode, string(respBody), url, data)
	}

	return respBody, nil
}

func RequestGetWithBody(url string, header map[string]string, queryData map[string]interface{}, body interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}

	if req, err = http.NewRequest("GET", url, bytes.NewReader(reqBody)); err != nil {
		return nil, err
	}

	if len(queryData) > 0 {
		query := req.URL.Query()
		for k, v := range queryData {
			query.Set(k, fmt.Sprint(v))
		}

		req.URL.RawQuery = query.Encode()
	}

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get failed, status=%d, resp=%+v, url=%v, query=%v",
			resp.StatusCode, string(respBody), url, queryData)
	}

	return respBody, nil
}

func RequestPOST(url string, header map[string]string, data interface{}) ([]byte, error) {
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

	req.Header.Add("Content-Type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func RequestPatch(url string, header map[string]string, data interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if req, err = http.NewRequest("PATCH", url, bytes.NewReader(reqBody)); err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func RequestPOSTWithBodyBytes(url string, header map[string]string, body []byte) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	if req, err = http.NewRequest("POST", url, bytes.NewReader(body)); err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func RequestPOSTForm(urlStr string, header map[string]string, data map[string]string) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	d := url.Values{}
	for k, v := range data {
		d[k] = []string{v}
	}
	form := d.Encode()

	if req, err = http.NewRequest("POST", urlStr, bytes.NewReader([]byte(form))); err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func RequestPOSTFile(urlStr string, header map[string]string, contentType string, body io.Reader) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	req, err = http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}

	// add headers
	req.Header.Add("Content-Type", contentType)
	for key, value := range header {
		req.Header.Add(key, value)
	}

	if resp, err = gClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
