package dao

import (
	"go-demos/projects/cgo/cgo"
	"go-demos/projects/cgo/entity"
	"log"
)

type PictureDao struct {
}

func (p *PictureDao) Insert(picture *entity.Picture) int64 {
	result, err := cgo.DB.Exec("INSERT INTO picture(`feedback_id`,`address`,`create_time`) VALUE(?,?,?)",
		picture.FeedbackID, picture.Address, picture.CreateTime)
	if err != nil {
		log.Println(err)
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0
	}
	return id
}
