package httputils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Client interface {
	RequestGet(ctx context.Context, url string, header map[string]string, data map[string]interface{}) ([]byte, error)
	RequestPOST(ctx context.Context, url string, header map[string]string, data interface{}) ([]byte, error)
	RequestPatch(ctx context.Context, url string, header map[string]string, data interface{}) ([]byte, error)
}

var (
	defaultClient Client
	once          sync.Once
)

type myClient struct {
	cli *http.Client
}

func GetClient(timeout time.Duration) Client {
	once.Do(func() {
		defaultClient = &myClient{
			cli: &http.Client{
				Timeout: timeout,
			},
		}
		return
	},
	)

	return defaultClient
}

func (c *myClient) RequestGet(ctx context.Context, url string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
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

	if resp, err = c.cli.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get failed, status=%d, resp=%+v, url=%v, query=%v", resp.StatusCode, string(respBody), url, data)
	}

	return respBody, nil
}

func (c *myClient) RequestPOST(ctx context.Context, url string, header map[string]string, data interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody)); err != nil {
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
		return nil, err
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (c *myClient) RequestPatch(ctx context.Context, url string, header map[string]string, data interface{}) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var reqBody []byte
	var respBody []byte

	reqBody, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(http.MethodPatch, url, bytes.NewReader(reqBody)); err != nil {
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
		return nil, err
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
