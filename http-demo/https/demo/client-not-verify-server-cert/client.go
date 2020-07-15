package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get1() {
	resp, err := http.Get("https://127.0.0.1:8080:/")
	if err != nil {
		// Get error: Get https://127.0.0.1:8080/:
		// x509: certificate signed by unknown authority
		// 显然从客户端日志来看，go实现的Client端默认也是要对服务端传过来的数字证书进行校验的
		// 但客户端提示：这个证书是由不知名CA签发的
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func get2() {
	// 客户端忽略校验服务器传过来的证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8080:/")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// get1()
	get2()
}
