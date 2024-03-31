package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "../../ca/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	// 代码通过将tls.Config.ClientAuth赋值为tls.RequireAndVerifyClientCert
	// 来实现Server强制校验client端证书。ClientCAs是用来校验客户端证书的
	// ca certificate
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Println("server start")
	err = s.ListenAndServeTLS("../../ca/server.crt", "../../ca/server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
