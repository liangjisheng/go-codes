package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
)

// Result ...
type Result struct {
	Num, Ans int
}

var client *rpc.Client

func init() {
	config := &tls.Config{
		// 如果客户端不需要对服务端鉴权，那么可以设置
		// InsecureSkipVerify:true，即可跳过对服务端的鉴权
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8080", config)
	if err != nil {
		log.Println(`tls.Dial("tcp", "127.0.0.1:8080", config) err:`, err)
		os.Exit(-1)
	}
	// defer conn.Close()
	client = rpc.NewClient(conn)
}

func syncCall() {
	// client.Call 是同步调用的方式，会阻塞当前的程序，直到结果返回
	// 如果有异步调用的需求，可以考虑使用client.Go
	var res Result
	if err := client.Call("Cal.Square", 12, &res); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d", res.Num, res.Ans)
}

func asyncCall() {
	var res Result
	aCall := client.Go("Cal.Square", 12, &res, nil)
	log.Printf("%d^2 = %d", res.Num, res.Ans)

	<-aCall.Done
	log.Printf("%d^2 = %d", res.Num, res.Ans)
}

// 如果需要对服务器端鉴权，那么需要将服务端的证书添加到信任证书池中
func trustSyncCall() {
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("../server/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.crt")
	}
	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		RootCAs: certPool,
	}

	// 这个得用localhost,不能用127.0.0.1,因为证书里写的是localhost
	conn, err := tls.Dial("tcp", "localhost:8080", config)
	if err != nil {
		log.Println(`tls.Dial("tcp", "localhost:8080", config) err:`, err)
		os.Exit(-1)
	}
	defer conn.Close()

	trustClient := rpc.NewClient(conn)

	var res Result
	if err := trustClient.Call("Cal.Square", 12, &res); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d", res.Num, res.Ans)
}

func main() {
	// syncCall()
	// asyncCall()

	trustSyncCall()
}
