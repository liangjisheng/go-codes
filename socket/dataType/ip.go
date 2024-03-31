package main

import (
	"fmt"
	"net"
	"os"
	"unsafe"
)

func main() {
	// ipOp()
	// ipOp1()
	tcpaddrParse()
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
	myip := "192.168.100.100"
	fmt.Println("myip:")
	typeof(myip)
	sizeof(myip)
	fmt.Println("    len is:", len(myip))

	addr := net.ParseIP(myip)
	fmt.Println("addr:")
	typeof(addr)
	sizeof(addr)
	fmt.Println("    len is:", len(addr))

	myerr := "1.1.1.1.1.1"
	erraddr := net.ParseIP(myerr)
	fmt.Println("erraddr is:", erraddr)
	if erraddr == nil {
		fmt.Println("no data")
	} else {
		typeof(erraddr)
		sizeof(erraddr)
	}

	myip6 := "1:1:1:1:1:1:1:1"
	addrv6 := net.ParseIP(myip6)
	fmt.Println("addrv6:")
	if addrv6 == nil {
		fmt.Println("no data")
	} else {
		fmt.Println(addrv6)
		typeof(addrv6)
		sizeof(addrv6)
		fmt.Println("    len is:", len(addrv6))
	}

	var mystr string
	mystr = "999999999kkkkkkkkkkkkkkkkkkkkkkkkk"
	fmt.Println("mystr")
	typeof(mystr)
	sizeof(mystr)
	fmt.Println("    len is:", len(mystr))
}

func tcpaddrParse() {
	addr := "www.baidu.com:80"
	tcpaddr, err := net.ResolveTCPAddr("", addr)
	checkError(err)

	fmt.Println("tcpaddr is:", tcpaddr)
	fmt.Println("IP is:", tcpaddr.IP.String(), "Port is", tcpaddr.Port)
	typeof(addr)
	typeof(tcpaddr)
	sizeof(addr)
	sizeof(tcpaddr)
	fmt.Println("addr len is:", len(addr))
	fmt.Println("tcpaddr len is:", len(tcpaddr.String()))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error !", err.Error())
		os.Exit(1)
	}
}

func typeof(v interface{}) {
	fmt.Printf("type is:%T\n", v)
}

func sizeof(v interface{}) {
	fmt.Println("sizeof is: ", unsafe.Sizeof(v))
}
