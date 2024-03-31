package main

import (
	"fmt"
	"net"
)

func mac() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for _, inter := range interfaces {
		macAddr := inter.HardwareAddr.String()
		// loop back ip (127.0.0.1) no macAddr
		if len(macAddr) == 0 {
			continue
		}
		fmt.Println(inter.Name)
		fmt.Println("MAC = ", macAddr)
	}
	// Output
	// lo
	// MAC =
	// enp4s0f2
	// MAC =  74:d0:2b:1a:94:e8
	// wlp3s0
	// MAC =  24:0a:64:1d:a1:13
	// br-7cc998ce71f4
	// MAC =  02:42:cf:0b:fe:f7
	// br-896e3e45603f
	// MAC =  02:42:8a:1e:a3:4f
	// docker0
	// MAC =  02:42:6d:00:17:27
	// br-e5451bbef790
	// MAC =  02:42:34:fc:6b:e1
	// br-62f5f627a4ec
	// MAC =  02:42:26:ef:3e:02
	// br-772f35b0686f
	// MAC =  02:42:fd:0e:4e:83
	// br-79cfadfb6c44
	// MAC =  02:42:ed:68:19:7c
}
