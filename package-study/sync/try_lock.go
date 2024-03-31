package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标识位置（从 1 开始）
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识 waiter 的起始 bit 位置
)

type Mutex struct {
	sync.Mutex
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {

	// 取锁的地址，从中可获取 state 的值。
	addr := (*int32)(unsafe.Pointer(&m.Mutex))

	// fast path：如果能成功抢到锁，返回 true。
	if atomic.CompareAndSwapInt32(addr, 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，不参与竞争返回 false。
	old := atomic.LoadInt32(addr)
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32(addr, old, new)
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
