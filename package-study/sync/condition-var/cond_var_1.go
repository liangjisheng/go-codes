package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const maxClients = 3

func cond3() {
	runtime.GOMAXPROCS(4)
	testCond3()
}

func testCond3() {
	s := NewServer()
	go s.IOloop()

	time.Sleep(time.Second * 1)
	go func() {
		s.Release()
	}()

	go func() {
		s.Release()
	}()

	time.Sleep(time.Second * 1)
	s.Release()
	time.Sleep(time.Second * 1)
	fmt.Println("[testCond] end.")
}

// Server ...
type Server struct {
	clients uint64
	cond    *sync.Cond
}

// NewServer ...
func NewServer() *Server {
	s := &Server{}
	s.cond = sync.NewCond(&sync.Mutex{})
	return s
}

// IOloop ...
func (s *Server) IOloop() {
	for {
		s.cond.L.Lock()
		for s.clients == maxClients {
			fmt.Println("[IOloop] 等于MAX_CLIENTS了,等待Cond通知.即有触发Release()")
			s.cond.Wait()
		}
		s.cond.L.Unlock()
		s.clients++
		fmt.Println("[IOloop] clients:", s.clients)
	}
}

// Release ...
func (s *Server) Release() {
	s.cond.L.Lock()
	s.clients--
	fmt.Println("[Release] a clients:", s.clients)
	s.cond.Signal()
	fmt.Println("[Release] b clients:", s.clients)
	s.cond.L.Unlock()
}
