package utils

import (
	"fmt"
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// NewToken ...
func NewToken(uid int32) (string, error) {
	u, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("%d-%s", uid, u.String())
	return token, nil
}

// RemoveUserSessions 清除一个用户的所有Session
func RemoveUserSessions(uid int32, pool *redis.Pool) error {
	var tokenkeys []interface{}
	keyOfSet := KeyOfSet(uid)
	conn := pool.Get()
	tokens, err := redis.Strings(conn.Do("SMEMBERS", keyOfSet))
	conn.Close()
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil
	}
	for _, token := range tokens {
		tokenkeys = append(tokenkeys, KeyOfToken(token))
	}
	log.Println("退出或重登时清除用户的所有旧token", zap.Strings("tokens", tokens))
	conn = pool.Get()
	if _, err = conn.Do("DEL", tokenkeys...); err != nil {
		log.Println("删除token出错", zap.Error(err), zap.Any("tokenkeys", tokenkeys))
	}

	if _, err = conn.Do("DEL", keyOfSet); err != nil {
		log.Println("删除keyofset出错", zap.Error(err), zap.String("key", keyOfSet))
	}
	conn.Close()
	return nil
}

// GetUIDByToken ...
func GetUIDByToken(token string, pool *redis.Pool) (int64, error) {
	conn := pool.Get()
	str, err := redis.String(conn.Do("GET", KeyOfToken(token)))
	conn.Close()
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		}
		return 0, err
	}
	uid := int64(0)
	if len(str) == 0 {
		return 0, nil
	}
	uid, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// GetSession ...
func GetSession(uid int32, pool *redis.Pool) (*pb.Session, error) {
	conn := pool.Get()
	m, err := redis.StringMap(conn.Do("HGETALL", KeyOfSession(uid)))
	conn.Close()
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
	}
	//判断是否有值
	if len(m) == 0 {
		log.Println("通过uid获取session的值为空", zap.String("参数", KeyOfSession(uid)))
		return nil, nil
	}
	return Map2Session(m), nil
}
