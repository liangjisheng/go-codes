package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func get() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer cli.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	putResp, err := cli.Put(ctx, "name", "alice")
	defer cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}
	log.Printf("Header %+v\n", putResp.Header)

	ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := cli.Get(ctx, "name")
	defer cancelFunc()
	if err != nil {
		log.Println("cli.Get", err.Error())
		return
	}

	for _, v := range res.Kvs {
		log.Println("CreateRevision", v.CreateRevision)
		log.Println("Version", v.Version)
		log.Println("ModRevision", v.ModRevision)
		log.Println("Lease", v.Lease)
		log.Println("key:", string(v.Key))
		log.Println("value:", string(v.Value))
	}
}
