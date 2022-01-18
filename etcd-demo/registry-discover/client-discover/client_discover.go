package main

import (
	"context"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"
)

// ClientDis ...
type ClientDis struct {
	client     *clientv3.Client
	serverList map[string]string
	lock       sync.Mutex
}

// NewClientDis ...
func NewClientDis(addr []string) (*ClientDis, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	var client *clientv3.Client
	var err error
	if client, err = clientv3.New(conf); err == nil {
		return &ClientDis{
			client:     client,
			serverList: make(map[string]string),
		}, nil
	}
	return nil, err
}

// GetService ...
func (s *ClientDis) GetService(prefix string) ([]string, error) {
	resp, err := s.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := s.extractAddrs(resp)

	go s.watcher(prefix)
	return addrs, nil
}

func (s *ClientDis) watcher(prefix string) {
	ch := s.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for {
		select {
		case c := <-ch:
			for _, e := range c.Events {
				switch e.Type {
				case clientv3.EventTypePut:
					s.SetServiceList(string(e.Kv.Key), string(e.Kv.Value))
				case clientv3.EventTypeDelete:
					s.DelServiceList(string(e.Kv.Key))
				}
			}
		}
	}
}

func (s *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			log.Println("key:", string(resp.Kvs[i].Key), "value:", string(resp.Kvs[i].Value))
			s.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

// SetServiceList ...
func (s *ClientDis) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = string(val)
	log.Println("set data key :", key, "val:", val)
}

// DelServiceList ...
func (s *ClientDis) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("del data key:", key)
}

// SerList2Array ...
func (s *ClientDis) SerList2Array() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

func main() {
	cli, _ := NewClientDis([]string{"117.51.148.112:2379"})
	cli.GetService("/node")
	select {}
}
