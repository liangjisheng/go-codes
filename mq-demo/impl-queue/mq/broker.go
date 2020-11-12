package mq

import (
	"errors"
	"sync"
	"time"
)

// Broker ...
type Broker interface {
	publish(topic string, msg interface{}) error
	subscribe(topic string) (<-chan interface{}, error)
	unsubscribe(topic string, sub <-chan interface{}) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{})
	setConditions(capacity int)
}

// BrokerImpl ...
type BrokerImpl struct {
	exit     chan bool // 关闭消息队列
	capacity int
	// 一个订阅者对应着一个通道, 一个topic可以有多个订阅者
	topics       map[string][]chan interface{} // key:topic  value:queue
	sync.RWMutex                               // 同步锁
}

// NewBroker ...
func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}

func (b *BrokerImpl) publish(topic string, pub interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		// 没有这个 topic
		return nil
	}

	b.broadcast(pub, subscribers)
	return nil
}

func (b *BrokerImpl) broadcast(msg interface{}, subscribers []chan interface{}) {
	count := len(subscribers)
	concurrency := 1

	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}

	// 采用Timer 而不是使用time.After 原因：time.After会产生内存泄漏 在计时器触发之前，垃圾回收器不会回收Timer
	idleDuration := 5 * time.Millisecond
	idleTimeout := time.NewTimer(idleDuration)
	defer idleTimeout.Stop()

	pub := func(start int) {
		for j := start; j < count; j += concurrency {
			// 使用time.Reset() 方法再次激活定时器
			idleTimeout.Reset(idleDuration)
			select {
			case subscribers[j] <- msg:
			case <-idleTimeout.C:
			case <-b.exit:
				return
			}
		}
	}

	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}

func (b *BrokerImpl) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-b.exit:
		return nil, errors.New("broker closed")
	default:
	}

	ch := make(chan interface{}, b.capacity)
	b.Lock()
	b.topics[topic] = append(b.topics[topic], ch)
	b.Unlock()
	return ch, nil
}

func (b *BrokerImpl) unsubscribe(topic string, sub <-chan interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()

	if !ok {
		return nil
	}

	// delete subscriber
	var newSubs []chan interface{}
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}

	b.Lock()
	b.topics[topic] = newSubs
	b.Unlock()
	return nil
}

func (b *BrokerImpl) close() {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		// 这里主要是为了保证下一次使用该消息队列不发生冲突
		b.topics = make(map[string][]chan interface{})
		b.Unlock()
	}
}

func (b *BrokerImpl) setConditions(capacity int) {
	b.capacity = capacity
}

var _ Broker = &BrokerImpl{}
