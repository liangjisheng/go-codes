package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	WebHookURL = ""
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
	} `json:"data"`
}

func Text(content string) error {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	data := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": content,
		},
	}

	res, err := RequestPOST(WebHookURL, header, data)
	if err != nil {
		return fmt.Errorf("lark bot push message err %+v\n", err)
	}

	var resp Response
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return fmt.Errorf("unmarshal json err %+v\n", err)
	}

	if resp.Code != 0 || resp.Msg != "success" {
		return fmt.Errorf("lark bot push message response fail %+v\n", resp)
	}
	return nil
}

func RichText(title string, content [][]map[string]interface{}) error {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	data := map[string]interface{}{
		"msg_type": "post",
		"content": map[string]interface{}{
			"post": map[string]interface{}{
				"zh_cn": map[string]interface{}{
					"title":   title,
					"content": content,
				},
			},
		},
	}

	res, err := RequestPOST(WebHookURL, header, data)
	if err != nil {
		return fmt.Errorf("lark bot push message err %+v\n", err)
	}

	var resp Response
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return fmt.Errorf("unmarshal json err %+v\n", err)
	}

	if resp.Code != 0 || resp.Msg != "success" {
		return fmt.Errorf("lark bot push message response fail %+v\n", resp)
	}
	return nil
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

	if resp, err = http.DefaultClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is %d", resp.StatusCode)
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
