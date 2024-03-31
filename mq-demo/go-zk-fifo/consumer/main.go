package main

import (
	"fmt"

	"github.com/nladuo/go-zk-fifo/fifo"
)

var (
	hosts    = []string{"117.51.148.112:2181"}
	basePath = "/fifo"
	fifoData = []byte("the fifo data")
	prefix   = "seq-"
)

func consume(f *fifo.DistributedFIFO) {
	for {
		data := f.Poll()
		if len(data) != 0 {
			fmt.Println("Poll : ", string(data))
		}
	}
}

func main() {
	fifo.EstablishZkConn(hosts)
	myfifo := fifo.NewFifo(basePath, fifoData, prefix)
	for i := 0; i < 5; i++ {
		go consume(myfifo)
	}

	ch := make(chan int)
	<-ch
	fifo.CloseZkConn()
}
