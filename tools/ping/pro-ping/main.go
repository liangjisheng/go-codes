package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func ping1() {
	//Here is a very simple example that sends and receives three packets:

	//pinger, err := probing.NewPinger("www.baidu.com")
	//pinger, err := probing.NewPinger("www.google.com")
	pinger, err := probing.NewPinger("127.0.0.1")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	//如果回包大于等于1则判为ping通
	if stats.PacketsRecv >= 1 {
		fmt.Println("ping not ok")
	} else {
		fmt.Println("ping ok")
	}
}

func ping2() {
	//Here is an example that emulates the traditional UNIX ping command:

	//pinger, err := probing.NewPinger("www.google.com")
	pinger, err := probing.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}

	// Listen for Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}

	pinger.OnFinish = func(stats *probing.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	//ping1()
	ping2()
}
