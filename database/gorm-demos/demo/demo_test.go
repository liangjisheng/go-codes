package mysql

import (
	"testing"
	"time"
)

func TestAddDemo(t *testing.T) {
	demo := &Demo{
		CreateTime:  time.Now(),
		CreateTime1: time.Now(),
		UpdateTime:  time.Now(),
	}

	err := Instance().AddDemo(demo)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("AddDemo ok.")
}
