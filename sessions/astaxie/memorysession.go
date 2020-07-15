package main

import (
	"container/list"
	"sync"
	"time"
)

var provider = &ProviderStruct{list: list.New()}

// SessionStore ...
type SessionStore struct {
	sid          string                      // session id
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} // session里面存储的值
}

// Set ...
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	provider.SessionUpdate(st.sid)
	return nil
}

// Get ...
func (st *SessionStore) Get(key interface{}) interface{} {
	provider.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete ...
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	provider.SessionUpdate(st.sid)
	return nil
}

// SessionID ...
func (st *SessionStore) SessionID() string {
	return st.sid
}

// ProviderStruct ...
type ProviderStruct struct {
	lock     sync.Mutex
	sessions map[string]*list.Element // 用来存储在内存
	list     *list.List               // 用来做gc
}

// SessionInit ...
func (pder *ProviderStruct) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// SessionRead ...
func (pder *ProviderStruct) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	}

	sess, err := pder.SessionInit(sid)
	return sess, err
}

// SessionDestroy ...
func (pder *ProviderStruct) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC ...
func (pder *ProviderStruct) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// SessionUpdate ...
func (pder *ProviderStruct) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}
