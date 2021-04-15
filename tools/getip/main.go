package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

var getIP = flag.String("get_ip", "", "external|internal")

func main() {
	fmt.Println("Usage of ./main --get_ip=(external|internal)")
	flag.Parse()
	if *getIP == "external" {
		getExternal()
	}

	if *getIP == "internal" {
		getInternal()
	}
}

func getExternal() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}

func getInternal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
}
