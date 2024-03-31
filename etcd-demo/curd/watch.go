package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func watch1() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer cli.Close()
	fmt.Println("conn success")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, "name", "alice")
	cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}

	// watch
	for {
		rch := cli.Watch(context.Background(), "name")
		for resp := range rch {
			for k, v := range resp.Events {
				fmt.Println(k, v.Type, string(v.Kv.Key), string(v.Kv.Value))
			}
		}
	}
}

func watch2() {
	var (
		client *clientv3.Client
		err    error
	)

	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		log.Println(err.Error())
		return
	}

	ctxWithTimeout, cancelFunc := context.WithCancel(context.TODO())
	wch := client.Watch(ctxWithTimeout, "/cron/watch/job1")

	// 20 秒后调用取消函数关闭通道退出监控键 /cron/watch/job1
	tt := time.After(20 * time.Second)
	go func() {
		select {
		case <-tt:
			cancelFunc()
		}
	}()

	for resp := range wch {
		for _, res := range resp.Events {
			log.Println(res.Type, string(res.Kv.Key), string(res.Kv.Value))
		}
	}
}
