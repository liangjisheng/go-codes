package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func lock1() {
	var (
		client        *clientv3.Client
		err           error
		leaseResp     *clientv3.LeaseGrantResponse
		keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	)

	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		log.Println("client new err", err)
		return
	}

	lease := clientv3.NewLease(client)
	ctx, cancelFunc := context.WithCancel(context.TODO())
	// 做分布式锁相关,执行事务
	// 建立租约、用租约抢key，抢到后执行业务逻辑，抢失败返回。函数退出时要defer吧租约关闭
	if leaseResp, err = lease.Grant(ctx, 10); err != nil {
		log.Println("create lease err", err)
		cancelFunc()
		return
	}

	// defer逻辑可以保证租约被清理，防止长期占用key
	defer client.Revoke(ctx, leaseResp.ID)
	defer cancelFunc()

	keepAliveChan = make(<-chan *clientv3.LeaseKeepAliveResponse)
	if keepAliveChan, err = client.KeepAlive(ctx, leaseResp.ID); err != nil {
		log.Println("lease keepalive err", err)
		return
	}
	go func() {
		for {
			log.Println(<-keepAliveChan)
		}
	}()

	// 打开下面可以看锁已经被抢占的情况
	putResp, err := client.Put(ctx, "/cron/txn/job1", "I GET FIRST", clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Println("put key err:", err)
		return
	}
	log.Println(putResp.Header.String())

	select {}
}
