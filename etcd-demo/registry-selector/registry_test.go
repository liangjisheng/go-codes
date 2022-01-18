package registryselector

import (
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func TestNewRegistry(t *testing.T) {
	var op = Options{
		name: "svc.info",
		ttl:  10,
		config: clientv3.Config{
			Endpoints:   []string{"http://117.51.148.112:2379/"},
			DialTimeout: 5 * time.Second},
	}

	for i := 1; i <= 3; i++ {
		r, err := NewRegistry(op)
		if err != nil {
			t.Error(err)
			return
		}
		err = r.RegistryNode(PutNode{Addr: fmt.Sprintf("117.51.148.112:%d%d%d%d", i, i, i, i)})
		if err != nil {
			t.Error(err)
			return
		}
		if i == 3 {
			go func() {
				time.Sleep(time.Second * 3)
				r.UnRegistry()
			}()
		}
	}
	time.Sleep(time.Hour)
}

func TestNewSelector(t *testing.T) {
	var op = SelectorOptions{
		name: "svc.info",
		config: clientv3.Config{
			Endpoints:   []string{"http://117.51.148.112:2379/"},
			DialTimeout: 5 * time.Second},
	}
	s, err := NewSelector(op)
	if err != nil {
		t.Error(err)
		return
	}
	for {
		val, err := s.Next()
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}
}
