package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var conf = Conf{
	ClientId:     "",
	ClientSecret: "",
	RedirectUrl:  "http://localhost:8080/oauth/redirect",
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

// Hello 返回欢迎页面
func Hello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	var temp *template.Template
	var err error
	if temp, err = template.ParseFiles("views/hello.html"); err != nil {
		fmt.Println("读取文件失败，错误信息为:", err)
		return
	}

	// 利用给定数据渲染模板(html页面)，并将结果写入w，返回给前端
	if err = temp.Execute(w, conf); err != nil {
		fmt.Println("读取渲染html页面失败，错误信息为:", err)
		return
	}
}

// Oauth 认证并获取用户信息
func Oauth(w http.ResponseWriter, r *http.Request) {
	var err error
	// 获取 code
	var code = r.URL.Query().Get("code")
	// 获取 token
	var tokenAuthUrl = GetTokenAuthUrl(code)
	var token *Token
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}

	// 通过token，获取github用户信息
	var userInfo map[string]interface{}
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return
	}

	// 将用户信息返回前端
	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息(map)转为用户信息([]byte)时发生错误，错误信息为:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println("在将用户信息([]byte)返回前端时发生错误，错误信息为:", err)
		return
	}
}

// GetTokenAuthUrl 通过code获取token认证url
func GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		conf.ClientId, conf.ClientSecret, code,
	)
}

// GetToken 获取 token
func GetToken(url string) (*Token, error) {
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}

	// 将响应体解析为 token，并返回
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(token *Token) (map[string]interface{}, error) {
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/oauth/redirect", Oauth) // 这个和 Authorization callback URL 有关

	log.Println("server listen on 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("监听失败，错误信息为:", err)
		return
	}
}
