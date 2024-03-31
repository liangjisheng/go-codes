package main

// etcd的 租约模式:客户端申请 一个租约 并设置 过期时间，每隔一段时间就要请求 etcd 申请续租
// 客户端可以通过租约存key 如果不续租过期了 etcd 会删除这个租约上的 所有key-value 类似于心跳模式
// 一般相同的服务存的 key 的前缀是一样的 比如 "server/001" => "127.0.0.1:1212" 和
// "server/002"=>"127.0.0.1:1313" 这种模式然后客户端就直接匹配 "server/" 这个key

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

// ServiceReg 租约注册服务
type ServiceReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	canclefunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

// NewServiceReg ...
func NewServiceReg(addr []string, timeNum int64) (*ServiceReg, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	var client *clientv3.Client

	if clientTem, err := clientv3.New(conf); err == nil {
		client = clientTem
	} else {
		return nil, err
	}

	ser := &ServiceReg{
		client: client,
	}
	if err := ser.setLease(timeNum); err != nil {
		return nil, err
	}
	go ser.ListenLeaseRespChan()
	return ser, nil
}

// setLease 设置租约
func (s *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(s.client)

	// 设置租约时间
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}

	s.lease = lease
	s.leaseResp = leaseResp
	s.canclefunc = cancelFunc
	s.keepAliveChan = leaseRespChan
	return nil
}

// ListenLeaseRespChan 监听 续租情况
func (s *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-s.keepAliveChan:
			if leaseKeepResp == nil {
				fmt.Printf("已经关闭续租功能\n")
				return
			}
			fmt.Printf("续租成功\n")
		}
	}
}

// PutService 通过租约 注册服务
func (s *ServiceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(s.client)
	_, err := kv.Put(context.TODO(), key, val, clientv3.WithLease(s.leaseResp.ID))
	return err
}

// RevokeLease 撤销租约
func (s *ServiceReg) RevokeLease() error {
	s.canclefunc()
	time.Sleep(2 * time.Second)
	_, err := s.lease.Revoke(context.TODO(), s.leaseResp.ID)
	return err
}

func main() {
	ser, _ := NewServiceReg([]string{"117.51.148.112:2379"}, 5)
	ser.PutService("/node/111", "heiheihei")
	ser.PutService("/node/222", "hahaha")
	select {}
}
