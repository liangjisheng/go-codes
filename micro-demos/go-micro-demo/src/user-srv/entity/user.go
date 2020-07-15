package entity

import (
	"go-demos/micro-demos/go-micro-demo/src/share/pb"
)

// User ...
type User struct {
	ID      int32  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"Address" db:"address"`
	Phone   string `json:"Phone" db:"phone"`
}

// ToProtoUser ...
func (u User) ToProtoUser() *pb.User {
	return &pb.User{
		Id:      u.ID,
		Name:    u.Name,
		Address: u.Address,
		Phone:   u.Phone,
	}
}
