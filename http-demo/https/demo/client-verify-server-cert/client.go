package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 由于client端需要验证server端的数字证书，因此client端需要预先加载ca.crt
	// 以用于服务端数字证书的校验
	pool := x509.NewCertPool()
	caCertPath := "../ca/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	// 这里需要写和生成server.csr是的ip一样,这里不能替换为127.0.0.1
	resp, err := client.Get("https://localhost:8080:/")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
