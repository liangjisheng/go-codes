package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func get1() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.51.148.112:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer cli.Close()
	fmt.Println("conn success")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, "name", "ljs")
	cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}

	ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := cli.Get(ctx, "name")
	cancelFunc()
	if err != nil {
		log.Println("cli.Get", err.Error())
		return
	}

	for k, v := range res.Kvs {
		fmt.Println("res", k, string(v.Key), string(v.Value))
	}
}
