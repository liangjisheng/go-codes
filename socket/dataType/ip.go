package main

import (
	"fmt"
	"net"
)

func main() {
	// ipOp()
	// ipOp1()
	tcpAddrParse()
}

func ipOp() {
	ip := "127.0.0.1"
	// ParseIP(s string) IP函数会把一个IPv4或者IPv6的地址转化成IP类型
	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}

	ip = "2002:c0e8:82e7:0:0:0:c0e8:82e7"
	addr = net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}
}

func ipOp1() {
	myIP := "192.168.100.100"
	fmt.Println("myIP:")
	typeof(myIP)
	sizeof(myIP)
	fmt.Println("    len is:", len(myIP))

	addr := net.ParseIP(myIP)
	fmt.Println("addr:")
	typeof(addr)
	sizeof(addr)
	fmt.Println("    len is:", len(addr))

	myErr := "1.1.1.1.1.1"
	errAddr := net.ParseIP(myErr)
	fmt.Println("erraddr is:", errAddr)
	if errAddr == nil {
		fmt.Println("no data")
	} else {
		typeof(errAddr)
		sizeof(errAddr)
	}

	myIP6 := "1:1:1:1:1:1:1:1"
	addrV6 := net.ParseIP(myIP6)
	fmt.Println("addrv6:")
	if addrV6 == nil {
		fmt.Println("no data")
	} else {
		fmt.Println(addrV6)
		typeof(addrV6)
		sizeof(addrV6)
		fmt.Println("    len is:", len(addrV6))
	}

	var myStr string
	myStr = "999999999kkkkkkkkkkkkkkkkkkkkkkkkk"
	fmt.Println("mystr")
	typeof(myStr)
	sizeof(myStr)
	fmt.Println("    len is:", len(myStr))
}

func tcpAddrParse() {
	addr := "www.baidu.com:80"
	tcpAddr, err := net.ResolveTCPAddr("", addr)
	checkError(err)

	fmt.Println("tcpAddr is:", tcpAddr)
	fmt.Println("IP is:", tcpAddr.IP.String(), "Port is", tcpAddr.Port)
	typeof(addr)
	typeof(tcpAddr)
	sizeof(addr)
	sizeof(tcpAddr)
	fmt.Println("addr len is:", len(addr))
	fmt.Println("tcpaddr len is:", len(tcpAddr.String()))
}
