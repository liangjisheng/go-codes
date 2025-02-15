package etcdv3

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
)

const schema = "grpclb"

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli        *clientv3.Client // etcd client
	cc         resolver.ClientConn
	serverList map[string]resolver.Address // 服务列表
	lock       sync.Mutex
}

// NewServiceDiscovery  新建发现服务
func NewServiceDiscovery(endpoints []string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceDiscovery{
		cli: cli,
	}
}

// Build 为给定目标创建一个新的`resolver`，当调用`grpc.Dial()`时执行
func (s *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	log.Println("Build")
	s.cc = cc
	s.serverList = make(map[string]resolver.Address)
	prefix := "/" + target.Scheme + "/" + target.Endpoint + "/"
	// 根据前缀获取现有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}
	s.cc.UpdateState(resolver.State{Addresses: s.GetServices()})
	// 监视前缀, 修改变更的server
	go s.watcher(prefix)
	return s, nil
}

// ResolveNow 监视目标更新
func (s *ServiceDiscovery) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}

// Scheme return schema
func (s *ServiceDiscovery) Scheme() string {
	return schema
}

// Close 关闭
func (s *ServiceDiscovery) Close() {
	log.Println("Close")
	s.cli.Close()
}

// watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix:%s now...", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: // 修改或者新增
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: // 删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	s.serverList[key] = resolver.Address{Addr: val}
	s.lock.Unlock()

	s.cc.UpdateState(resolver.State{Addresses: s.GetServices()})
	log.Println("put key :", key, "val:", val)
}

// DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	delete(s.serverList, key)
	s.lock.Unlock()

	s.cc.UpdateState(resolver.State{Addresses: s.GetServices()})
	log.Println("del key:", key)
}

// GetServices 获取服务地址
func (s *ServiceDiscovery) GetServices() []resolver.Address {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]resolver.Address, 0, len(s.serverList))
	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

// func main() {
// 	var endpoints = []string{"117.51.148.112:2379"}
// 	ser := NewServiceDiscovery(endpoints)
// 	defer ser.Close()
// 	ser.WatchService("/web/")
// 	ser.WatchService("/gRPC/")
// 	for {
// 		select {
// 		case <-time.Tick(10 * time.Second):
// 			log.Println(ser.GetServices())
// 		}
// 	}
// }
