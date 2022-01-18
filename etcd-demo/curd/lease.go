package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func lease1() {
	var (
		config        clientv3.Config
		client        *clientv3.Client
		err           error
		putResp       *clientv3.PutResponse
		getResp       *clientv3.GetResponse
		leaseResp     *clientv3.LeaseGrantResponse
		keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		log.Println("client new err", err)
		return
	}

	ctx := context.TODO()
	lease := clientv3.NewLease(client)
	// 获取一个租约 存活时间为10秒
	if leaseResp, err = lease.Grant(ctx, 10); err != nil {
		log.Println("create lease err", err)
		return
	}

	// 用client也可以设置key，kv是client的一个结构，因此可以使用其方法
	if putResp, err = client.Put(ctx, "/cron/lock/job1", "ok", clientv3.WithLease(leaseResp.ID)); err != nil {
		log.Println("put key err", err)
		return
	}
	log.Println("put response:", putResp.Header.Revision)

	// 由协程来帮自动续租，每秒一次
	keepAliveChan = make(<-chan *clientv3.LeaseKeepAliveResponse)
	if keepAliveChan, err = client.KeepAlive(ctx, leaseResp.ID); err != nil {
		log.Println("keepalive err", err)
		return
	}
	go func() {
		for {
			select {
			case resp := <-keepAliveChan:
				if resp == nil {
					log.Println("续租失败")
					return
				}
				// log.Println(resp)
				log.Println("续租成功")
			}
		}
	}()

	k := 20
	for k != 0 {
		if getResp, err = client.Get(ctx, "/cron/lock/job1"); err != nil {
			log.Println("get key err", err)
			return
		}
		log.Println(getResp.Kvs)
		time.Sleep(1 * time.Second)
		k--
	}
}
