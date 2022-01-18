package registryselector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

var prefix = "/registry/server/"

// Registry ...
type Registry interface {
	RegistryNode(node PutNode) error
	UnRegistry()
}

type registryServer struct {
	cli        *clientv3.Client
	stop       chan bool
	isRegistry bool
	options    Options
	leaseID    clientv3.LeaseID
}

// PutNode ...
type PutNode struct {
	Addr string `json:"addr"`
}

// Node ...
type Node struct {
	ID   uint32 `json:"id"`
	Addr string `json:"addr"`
}

// Options ...
type Options struct {
	name   string
	ttl    int64
	config clientv3.Config
}

// NewRegistry ...
func NewRegistry(options Options) (Registry, error) {
	cli, err := clientv3.New(options.config)
	if err != nil {
		return nil, err
	}
	return &registryServer{
		stop:       make(chan bool),
		options:    options,
		isRegistry: false,
		cli:        cli,
	}, nil
}

// RegistryNode ...
func (s *registryServer) RegistryNode(put PutNode) error {
	if s.isRegistry {
		return errors.New("only one node can be registered")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.options.ttl)*time.Second)
	defer cancel()
	grant, err := s.cli.Grant(context.Background(), s.options.ttl)
	if err != nil {
		return err
	}
	var node = Node{
		ID:   s.HashKey(put.Addr),
		Addr: put.Addr,
	}
	nodeVal, err := s.GetVal(node)
	if err != nil {
		return err
	}
	log.Printf("registry node key: %s, value: %s\n", s.GetKey(node), nodeVal)
	_, err = s.cli.Put(ctx, s.GetKey(node), nodeVal, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	s.leaseID = grant.ID
	s.isRegistry = true
	go s.KeepAlive()
	log.Printf("registry server: %+v\n", *s)
	return nil
}

// UnRegistry ...
func (s *registryServer) UnRegistry() {
	s.stop <- true
}

func (s *registryServer) Revoke() error {
	_, err := s.cli.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[Revoke] err : %s", err.Error())
	}
	s.isRegistry = false
	return err
}

func (s *registryServer) KeepAlive() error {
	keepAliveCh, err := s.cli.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[KeepAlive] err : %s", err.Error())
		return err
	}
	for {
		select {
		case <-s.stop:
			_ = s.Revoke()
			return nil
		case _, ok := <-keepAliveCh:
			if !ok {
				_ = s.Revoke()
				return nil
			}
		}
	}
}

// GetKey ...
func (s *registryServer) GetKey(node Node) string {
	return fmt.Sprintf("%s%s/%d", prefix, s.options.name, s.HashKey(node.Addr))
}

// GetVal ...
func (s *registryServer) GetVal(node Node) (string, error) {
	data, err := json.Marshal(&node)
	return string(data), err
}

// HashKey ...
func (s *registryServer) HashKey(addr string) uint32 {
	return crc32.ChecksumIEEE([]byte(addr))
}
