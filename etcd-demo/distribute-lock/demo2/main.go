package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	prefix := "/lock"

	go func () {
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, prefix)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go1 get mutex failed " + err.Error())
		}
		fmt.Printf("go1 get mutex sucess\n")
		fmt.Printf("mutex %+v\n", m)
		fmt.Printf("header %+v\n", m.Header())
		fmt.Printf("key %+v\n", m.Key())
		time.Sleep(time.Duration(2) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go1 release lock\n")
	}()

	go func() {
		session, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m := concurrency.NewMutex(session, prefix)
		if err := m.Lock(context.TODO()); err != nil {
			log.Fatal("go2 get mutex failed " + err.Error())
		}
		fmt.Printf("go2 get mutex sucess\n")
		fmt.Printf("mutex %+v\n", m)
		fmt.Printf("header %+v\n", m.Header())
		fmt.Printf("key %+v\n", m.Key())
		time.Sleep(time.Duration(2) * time.Second)
		m.Unlock(context.TODO())
		fmt.Printf("go2 release lock\n")
	}()

	<-c
}
