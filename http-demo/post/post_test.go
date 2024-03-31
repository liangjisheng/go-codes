package post_test

import (
	"encoding/json"
	"fmt"
	"httpdemo/post"
	"testing"
	"time"
)

func TestPOSTCreatePost(t *testing.T) {
	url := "http://127.0.0.1:8899/api/post/create"
	head := make(map[string]string)
	head["Content-Type"] = "application/json"
	head["zk-uuid"] = "e5e2bf2d-1f55-4f42-ac75-8ff8335f893f"

	data := make(map[string]interface{})
	data["title"] = "title8"
	data["summary"] = "summary8"
	data["author"] = "alice8"
	data["time"] = time.Now().UTC().Unix()
	data["thumb"] = "thumb8"
	data["show"] = 1
	data["weight"] = 1
	data["category"] = 2
	data["operator"] = "alice8"
	data["drop"] = 0
	data["content"] = "content8"
	data["language"] = "en"

	reqBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Printf(string(reqBody))

	bytes, err := post.RequestPost(url, head, data)
	if err != nil {
		t.Error("POSTCreatePost error ", err)
		return
	}

	t.Log(string(bytes))
}

func TestGETPost(t *testing.T) {
	url := "http://127.0.0.1:8899/api/post/1"
	head := make(map[string]string)
	head["zk-uuid"] = "e5e2bf2d-1f55-4f42-ac75-8ff8335f893f"

	bytes, err := post.RequestGet(url, head, nil)
	if err != nil {
		t.Error("GETPost error ", err)
		return
	}

	t.Log(string(bytes))
}
