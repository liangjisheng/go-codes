package handler

import (
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
	tlog "go-demos/micro-demos/go-micro-demo/src/share/utils/log"
	"go-demos/micro-demos/go-micro-demo/src/user-srv/db"
	"go-demos/micro-demos/go-micro-demo/src/user-srv/entity"
	"log"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// UserHandler ...
type UserHandler struct {
	logger *zap.Logger
}

// NewUserHandler new一个UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		logger: tlog.Instance().Named("UserHandler"),
	}
}

// InsertUser ...
func (c *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserReq, rsp *pb.InsertUserRep) error {
	log.Println("InsertUser ...")
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}

	insertID, err := db.InsertUser(user)
	if err != nil {
		log.Fatal("user.db.InsertUser has a error")
		return err
	}
	rsp.Id = int32(insertID)
	return nil
}

// DeletetUser ...
func (c *UserHandler) DeletetUser(ctx context.Context, req *pb.DeletetUserReq, rsp *pb.DeletetUserRep) error {
	log.Println("DeletetUser ...")
	err := db.DeleteUser(req.GetId())
	if err != nil {
		log.Fatal("user.db.DeleteUser has a error")
		return err
	}
	return nil
}

// SelectUser ...
func (c *UserHandler) SelectUser(ctx context.Context, req *pb.SelectUserReq, rsp *pb.SelectUserRep) error {
	c.logger.Info("SelectUser ... ")
	user, err := db.SelectUserByID(req.GetId())
	if err != nil {
		log.Fatal("user.db.SelectUserByUId has a error")
		return err
	}
	if user != nil {
		rsp.Users = user.ToProtoUser()
	}
	return nil
}

// ModifyUser ...
func (c *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserReq, rsp *pb.ModifyUserRep) error {
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		ID:      req.Id,
	}

	log.Println("ModifyUser ...", req.GetId())
	err := db.ModifyUser(user)
	if err != nil {
		log.Fatal("user.db.SelectUserByUId通过id获取用户信息出错")
		return err
	}
	return nil
}
