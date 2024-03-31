package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postJSON() {
	url := "http://baidu.com"
	usrID := "alice"
	pwd := "123456"
	//json序列化
	post := "{\"UserId\":\"" + usrID + "\",\"Password\":\"" + pwd + "\"}"
	fmt.Println(url, "post", post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(post)))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func postString() {
	url := "http://localhost:9090/page"
	payload := strings.NewReader("username=alice&usertext=teacher")
	// post string
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func postBytes() {
	requestURL := "http://localhost:9090/page"
	// 要post的参数
	form := url.Values{
		"username": {"alice"},
		"usertext": {"teacher"},
	}
	// func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	// 对form进行编码
	fmt.Println(form.Encode())
	body := bytes.NewBufferString(form.Encode())
	// post bytes数据给URL
	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
}

func postForm() {
	requestURL := "http://localhost:9090/page"
	resp, err := http.PostForm(requestURL, url.Values{"username": {"alice"}, "usertext": {"teacher"}})
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
	// postJSON()
	// postString()
	// postBytes()
	//postForm()
}

func RequestPost(url string, header map[string]string, data interface{}) ([]byte, error) {
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

func RequestPut(url string, header map[string]string, data interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if req, err = http.NewRequest("PUT", url, bytes.NewReader(reqBody)); err != nil {
		return nil, err
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

func RequestDelete(url string, header map[string]string, data interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if req, err = http.NewRequest("DELETE", url, bytes.NewReader(reqBody)); err != nil {
		return nil, err
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
