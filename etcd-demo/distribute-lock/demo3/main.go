package main

import (
	"github.com/zieckey/etcdsync"
	"log"
	"time"
)

func main() {
	m, err := etcdsync.New("/mylock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	log.Printf("start get lock")
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
	} else {
		log.Printf("etcdsync.Lock OK")
	}

	log.Printf("Get the lock. Do something here.")
	time.Sleep(5 * time.Second)

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
