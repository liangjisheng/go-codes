package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const endpoint = "http://127.0.0.1:8080/"

// 发送一个简单的http GET请求
func httpSimpleGet() {
	resp, err := http.Get(endpoint + "index?aa=AA&bb=BB")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	// 获取响应内容
	resultByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resultByte))
}

// 设置请求头和请求参数的Get请求
func httpGet() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Println(err)
	}
	// 在请求头中添加自定义数据
	request.Header.Add("company", "PG")
	request.Header.Add("appkey", "Test_0001")
	// 添加请求参数
	params := request.URL.Query()
	params.Add("name", "pg")
	params.Add("addr", "chain")
	request.URL.RawQuery = params.Encode()
	// 发送http请求,请求成功,获取响应
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	// 获取所有的响应内容
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// 打印响应的内容
	fmt.Println(string(result))
}

// http POST请求,并发送json数据
func httpPostJSON() {
	// post json 数据应用比较广泛
	// 发送json数据,我们一般是用使用map或者结构体存储数据
	// 然后转换成json数据
	// 然后转换成byte数据,放在发送的body中一起发送
	// 我们模拟一下这个过程
	var std = map[string]string{"work": "programmer", "skills": "golang", "addr": "北京"}
	data, err := json.Marshal(std)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		log.Println(err)
	}
	// 设置请求头
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(result))
}

// 模拟post发送表单数据
func httpPostForm() {
	formData := url.Values{}
	formData.Set("userName", "admin")
	formData.Set("userPwd", "admin123456")
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(formData.Encode())))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
}

//模拟客户端发送DELETE请求
func httpDelete() {
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		log.Println(err)
	}
	// 添加请求参数 与发送get请求类似
	params := req.URL.Query()
	params.Add("user", "pahnaskdjalsdklasd")
	req.URL.RawQuery = params.Encode()
	// 发送http请求,请求成功,获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	// 获取所有的响应内容
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// 打印响应的内容
	fmt.Println(string(result))
}

func main() {
	// httpSimpleGet()
	// httpGet()
	// httpPostJSON()
	// httpPostForm()
	httpDelete()
}
