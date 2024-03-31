package cgo

import (
	"go-demos/projects/cgo/session"
	_ "go-demos/projects/cgo/session/memory"
)

var globalSession *session.Manager

func init() {
	var err error
	globalSession, err = session.NewManager("memory", "GSESSIONID", 3600)
	if err != nil {
		panic(err)
	}
	globalSession.GC()
}

func GlobalSession() *session.Manager {
	return globalSession
}
