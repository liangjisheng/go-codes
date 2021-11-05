package model

import "go-demos/projects/cgo/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}
