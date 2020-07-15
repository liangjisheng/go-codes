package handler

import (
	"bytes"
	"context"

	proto "demo/proto/srv.user"
)

// LoginServer ...
type LoginServer struct{}

// LoginByUserName ...
func (s *LoginServer) LoginByUserName(ctx context.Context, req *proto.LoginByUserNameReq, rsp *proto.LoginByUserNameResp) error {
	if !bytes.Equal(req.Username, []byte("liangjisheng")) {
		rsp.ErrNo = proto.LoginByUserNameResp_ERROR_USER
		rsp.ErrMsg = "user not found"
		return nil
	}

	if !bytes.Equal(req.Password, []byte("ljs")) {
		rsp.ErrNo = proto.LoginByUserNameResp_ERROR_PWD
		rsp.ErrMsg = "invalid password"
		return nil
	}

	return nil
}
