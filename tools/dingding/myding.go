package myding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type robot struct {
	accessToken string
	secret      string
}
type MsgInfo struct {
	Msgtype  string `json:"msgtype"`
	At       At
	Markdown *Markdown `json:"markdown,omitempty"`
	Text     *Text     `json:"text,omitempty"`
	Link     *Link     `json:"link,omitempty"`
}
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type Resp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewRobot(accessToken, secret string) *robot {
	return &robot{accessToken: accessToken, secret: secret}
}

func (r *robot) SendMsg(msg MsgInfo) error {
	if msg.Text != nil {
		msg.Msgtype = "text"
	} else if msg.Markdown != nil {
		msg.Msgtype = "markdown"
	} else if msg.Link != nil {
		msg.Msgtype = "link"
	}

	timestamp := time.Now().UnixMilli()
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&timestamp=%v&sign=%v", r.accessToken, timestamp, r.Sign(timestamp))
	out, err := OnPostJSON(url, GetJSONStr(msg, false))
	if err != nil {
		return err
	}
	var resp Resp
	_ = json.Unmarshal(out, &resp)
	if resp.Errcode != 0 {
		return fmt.Errorf("ding send err:%v", resp.Errmsg)
	}
	return nil
}

func (r *robot) Sign(timestamp int64) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, r.secret)
	hash := hmac.New(sha256.New, []byte(r.secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	return url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
}

// GetJSONStr obj to json string
func GetJSONStr(obj interface{}, isFormat bool) string {
	var b []byte
	if isFormat {
		b, _ = json.MarshalIndent(obj, "", "     ")
	} else {
		b, _ = json.Marshal(obj)
	}
	return string(b)
}

//OnPostJSON 发送修改密码
func OnPostJSON(url, jsonstr string) ([]byte, error) {
	//解析这个 URL 并确保解析没有出错。
	body := bytes.NewBuffer([]byte(jsonstr))
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body1, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return nil, err
	}

	return body1, nil
}
