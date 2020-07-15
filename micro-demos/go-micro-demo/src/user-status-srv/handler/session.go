package handler

import (
	"context"
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	"go-demos/micro-demos/go-micro-demo/src/user-status-srv/utils"

	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

//GetUserIDByToken GetUIDByToken
func (s *UserStatusHandler) GetUserIDByToken(ctx context.Context, req *pb.GetUserIDByTokenReq, rsp *pb.GetUserIDByTokenRep) error {
	uid, err := utils.GetUIDByToken(req.Token, s.pool)
	if err != nil {
		s.logger.Info("根据token获取uid失败", zap.Error(err))
	}
	if uid <= 0 {
		s.logger.Info("根据token获取的uid小于等于0", zap.Int64("uid", uid))
	}
	rsp.Uid = uid
	return nil
}

//GetSessionByUID GetSessionByUID
func (s *UserStatusHandler) GetSessionByUID(ctx context.Context, req *pb.GetSessionByUIDReq, rsp *pb.GetSessionByUIDRep) error {
	session, err := utils.GetSession(int32(req.Uid), s.pool)
	if err != nil {
		s.logger.Info("根据uid获取session失败", zap.Error(err))
		utils.SessionFree(session)
		return err
	}
	if session == nil {
		s.logger.Info("根据uid获取到空session")
		return err
	}
	sessionTemp := *session
	rsp.Session = &sessionTemp
	utils.SessionFree(session)
	return nil
}

//GetSessionByToken GetSessionByToken
func (s *UserStatusHandler) GetSessionByToken(ctx context.Context, req *pb.GetSessionByTokenReq, rsp *pb.GetSessionByTokenRep) error {
	uid, err := utils.GetUIDByToken(req.Token, s.pool)
	if err != nil {
		s.logger.Info("根据token获取uid失败", zap.Error(err), zap.String("token", req.Token))
		return err
	}
	if uid <= 0 {
		s.logger.Info("根据token找到的uid小于等于0", zap.Int64("uid", uid))
		return err
	}
	session, err := utils.GetSession(int32(uid), s.pool)
	if err != nil {
		s.logger.Info("根据uid获取session失败", zap.Error(err), zap.Int64("uid", uid))
		utils.SessionFree(session)
		return err
	}
	if session == nil {
		s.logger.Info("根据uid取到空session", zap.Int64("uid", uid))
		return err
	}
	sessionTemp := *session
	rsp.Session = &sessionTemp
	utils.SessionFree(session)
	return nil
}

//NewSession NewSession
func (s *UserStatusHandler) NewSession(ctx context.Context, req *pb.NewSessionReq, rsp *pb.NewSessionRep) error {
	var oldSession *pb.Session
	defer func() {
		utils.SessionFree(oldSession)
	}()
	fieldMap := make(map[string]interface{}, 0)
	fieldMap["Uid"] = req.Id
	fieldMap["Address"] = req.Address
	fieldMap["Phone"] = req.Phone
	fieldMap["Name"] = req.Name
	//生成Token
	token, err := utils.NewToken(req.Id)
	if err != nil {
		s.logger.Info("生成token失败", zap.Error(err), zap.Int32("uid", req.Id))
		return err
	}

	//删除所有旧token
	if err = utils.RemoveUserSessions(req.Id, s.pool); err != nil {
		s.logger.Info("删除所有旧token失败", zap.Error(err), zap.Int32("uid", req.Id))
		return err
	}
	conn := s.pool.Get()
	//会话数据写入redis，格式：t:id => map的哈希值
	if _, err := conn.Do("HMSET", redis.Args{}.Add(utils.KeyOfSession(req.Id)).AddFlat(fieldMap)...); err != nil {
		conn.Close()
		s.logger.Info("会话数据写入redis失败", zap.Error(err), zap.String("key", utils.KeyOfSession(req.Id)), zap.Any("参数", fieldMap))
		return err
	}
	//设置t:id的过期时间
	if _, err := conn.Do("EXPIRE", utils.KeyOfSession(req.Id), s.sessionExpire); err != nil {
		conn.Close()
		s.logger.Error("设置session过期时间失败", zap.Error(err), zap.String("key", utils.KeyOfSession(req.Id)))
		return err
	}

	//用户token写入set里边，格式：t:uid:set:id => token
	keyOfSet := utils.KeyOfSet(req.Id)
	if _, err = conn.Do("SADD", keyOfSet, token); err != nil {
		conn.Close()
		s.logger.Info("token写入用户集合失败", zap.Error(err), zap.String("key", keyOfSet), zap.String("参数", token))
		return err
	}
	//设置t:uid:set:id的过期时间
	if _, err = conn.Do("EXPIRE", keyOfSet, s.sessionExpire); err != nil {
		conn.Close()
		s.logger.Info("设置用户token集合过期时间失败", zap.Error(err), zap.String("key", keyOfSet))
		return err
	}

	//将token和id对应，格式：token => id
	if _, err = conn.Do("SETEX", utils.KeyOfToken(token), s.tokenExpire, req.Id); err != nil {
		conn.Close()
		s.logger.Info("token写入redis失败", zap.Error(err), zap.String("key", utils.KeyOfToken(token)), zap.Int32("参数", req.Id))
		return err
	}

	rsp.Token = token
	return nil
}

//RemoveSession RemoveSession
func (s *UserStatusHandler) RemoveSession(ctx context.Context, req *pb.RemoveSessionReq, rsp *pb.RemoveSessionRep) error {
	uid, err := utils.GetUIDByToken(req.Token, s.pool)
	if err != nil {
		s.logger.Info("获取token所对应的uid失败", zap.Error(err), zap.String("token", req.Token))
		return err
	}
	if uid <= 0 {
		s.logger.Info("获取token获取不到uid", zap.String("token", req.Token))
		return err
	}
	if err = utils.RemoveUserSessions(int32(uid), s.pool); err != nil {
		s.logger.Info("移除用户会话数据失败", zap.Error(err), zap.Int64("uid", uid))
		return err
	}
	return nil
}

//RefreshSession RefreshSession
func (s *UserStatusHandler) RefreshSession(ctx context.Context, req *pb.RefreshSessionReq, rsp *pb.RefreshSessionRep) error {
	id, err := utils.GetUIDByToken(req.Token, s.pool)
	uid := int32(id)
	if err != nil {
		s.logger.Info("根据token获取uid失败", zap.Error(err))
		return err
	}
	if uid <= 0 {
		s.logger.Info("根据token获取不到uid")
		return err
	}
	//获取session
	oldSession, err := utils.GetSession(int32(uid), s.pool)
	if err != nil {
		s.logger.Info("获取uid的会话数据失败", zap.Error(err))
		utils.SessionFree(oldSession)
		return err
	}
	if oldSession == nil {
		s.logger.Info("根据uid获取不到session")
		return err
	}
	keyOfSet := utils.KeyOfSet(oldSession.Id)
	conn := s.pool.Get()
	if _, err := conn.Do("EXPIRE", utils.KeyOfSession(int32(uid)), s.sessionExpire); err != nil {
		conn.Close()
		s.logger.Info("更新uid过期时间失败:", zap.Error(err), zap.String("key", utils.KeyOfSession(uid)))
		utils.SessionFree(oldSession)
		return err
	}
	if _, err := conn.Do("EXPIRE", keyOfSet, s.sessionExpire); err != nil {
		conn.Close()
		s.logger.Info("更新uid过期时间失败:"+keyOfSet, zap.Error(err), zap.String("key", keyOfSet))
		utils.SessionFree(oldSession)
		return err
	}
	if _, err := conn.Do("EXPIRE", req.Token, s.tokenExpire); err != nil {
		conn.Close()
		s.logger.Info("更新token过期时间失败:"+req.Token, zap.Error(err), zap.Int32("uid", uid))
		utils.SessionFree(oldSession)
		return err
	}
	conn.Close()
	rsp.Token = req.Token //token
	sessionTemp := *oldSession
	rsp.Session = &sessionTemp
	utils.SessionFree(oldSession)
	return nil
}

//RemoveSessionByUID RemoveSession
func (s *UserStatusHandler) RemoveSessionByUID(ctx context.Context, req *pb.RemoveSessionByUIDReq, rsp *pb.RemoveSessionByUIDRep) error {
	if err := utils.RemoveUserSessions(int32(req.Uid), s.pool); err != nil {
		s.logger.Info("移除用户会话失败", zap.Error(err), zap.Int64("uid", req.Uid))
		return err
	}
	return nil
}
