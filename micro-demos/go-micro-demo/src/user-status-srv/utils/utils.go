package utils

import (
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	"log"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

var (
	sessionPool = sync.Pool{New: func() interface{} { return new(pb.Session) }}
)

// SessionFree ...
func SessionFree(s *pb.Session) {
	if s != nil {
		sessionPool.Put(s)
	}
}

// GetSessionFromSessionPool ...
func GetSessionFromSessionPool() *pb.Session {
	session := sessionPool.Get().(*pb.Session)
	session.Reset()
	return session
}

// Map2Session ...
func Map2Session(m map[string]string) *pb.Session {
	var err error
	session := GetSessionFromSessionPool()
	for field, val := range m {
		switch field {
		case "Uid":
			if len(val) > 0 {
				var id int64
				id, err = strconv.ParseInt(val, 10, 64)
				session.Id = int32(id)
				if err != nil {
					log.Println("字符串转数出错", zap.String("field", field), zap.String("val", val))
				}
			}
		case "Name":
			session.Name = val
		case "Address":
			session.Address = val
		case "Phone":
			session.Phone = val
		}
	}
	return session
}
