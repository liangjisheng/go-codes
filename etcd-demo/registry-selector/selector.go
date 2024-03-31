package registryselector

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Selector ...
type Selector interface {
	Next() (Node, error)
}

type selectorServer struct {
	cli     *clientv3.Client
	nodes   []Node
	options SelectorOptions
}

// SelectorOptions ...
type SelectorOptions struct {
	name   string
	config clientv3.Config
}

// NewSelector ...
func NewSelector(options SelectorOptions) (Selector, error) {
	cli, err := clientv3.New(options.config)
	if err != nil {
		return nil, err
	}
	var s = &selectorServer{
		options: options,
		cli:     cli,
	}
	go s.Watch()
	return s, nil
}

func (s *selectorServer) Next() (Node, error) {
	if len(s.nodes) == 0 {
		return Node{}, fmt.Errorf("no node found on the %s", s.options.name)
	}
	i := rand.Int() % len(s.nodes)
	return s.nodes[i], nil
}

func (s *selectorServer) Watch() {
	res, err := s.cli.Get(context.TODO(), s.GetKey(), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		log.Printf("[Watch] err : %s", err.Error())
		return
	}
	for _, kv := range res.Kvs {
		node, err := s.GetVal(kv.Value)
		if err != nil {
			log.Printf("[GetVal] err : %s", err.Error())
			continue
		}
		s.nodes = append(s.nodes, node)
	}
	ch := s.cli.Watch(context.TODO(), prefix, clientv3.WithPrefix())
	for {
		select {
		case c := <-ch:
			for _, e := range c.Events {
				switch e.Type {
				case clientv3.EventTypePut:
					node, err := s.GetVal(e.Kv.Value)
					if err != nil {
						log.Printf("[EventTypePut] err : %s", err.Error())
						continue
					}
					s.AddNode(node)
				case clientv3.EventTypeDelete:
					keyArray := strings.Split(string(e.Kv.Key), "/")
					if len(keyArray) <= 0 {
						log.Printf("[EventTypeDelete] key Split err : %s", err.Error())
						return
					}
					nodeID, err := strconv.Atoi(keyArray[len(keyArray)-1])
					if err != nil {
						log.Printf("[EventTypePut] key Atoi : %s", err.Error())
						continue
					}
					s.DelNode(uint32(nodeID))
				}
			}
		}
	}
}

func (s *selectorServer) DelNode(id uint32) {
	var nodes []Node
	for _, v := range s.nodes {
		if v.ID != id {
			nodes = append(nodes, v)
		}
	}
	s.nodes = nodes
}

// AddNode ...
func (s *selectorServer) AddNode(node Node) {
	var exist bool
	for _, v := range s.nodes {
		if v.ID == node.ID {
			exist = true
		}
	}
	if !exist {
		s.nodes = append(s.nodes, node)
	}
}

// GetKey ...
func (s *selectorServer) GetKey() string {
	return fmt.Sprintf("%s%s", prefix, s.options.name)
}

// GetVal ...
func (s *selectorServer) GetVal(val []byte) (Node, error) {
	var node Node
	err := json.Unmarshal(val, &node)
	if err != nil {
		return node, err
	}
	return node, nil
}
