package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func curd1() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := context.TODO()

	if putResp, err = client.Put(ctx, "test", "haha"); err != nil {
		return
	}
	fmt.Println("put test ", putResp.Header)

	kv := clientv3.NewKV(client)
	// 用kv设置key
	if putResp, err = kv.Put(ctx, "/cron/jobs/job2", "hello"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("put job2 ", putResp.Header.String())

	// 用kv获取Key
	if getResp, err = kv.Get(ctx, "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResp.Kvs)

	// 用kv删除key
	if delResp, err = kv.Delete(ctx, "/cron/jobs/job2"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(delResp.Deleted)
}
