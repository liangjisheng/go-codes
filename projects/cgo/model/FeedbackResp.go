package model

import "go-demos/cgo/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}
