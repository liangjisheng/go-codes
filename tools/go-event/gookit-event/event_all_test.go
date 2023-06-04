package event__test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gookit/event"
)

var emptyListener = func(e event.Event) error {
	return nil
}

type testListener struct {
	userData string
}

func (l *testListener) Handle(e event.Event) error {
	if ret := e.Get("result"); ret != nil {
		str := ret.(string) + fmt.Sprintf(" -> %s(%s)", e.Name(), l.userData)
		e.Set("result", str)
	} else {
		e.Set("result", fmt.Sprintf("handled: %s(%s)", e.Name(), l.userData))
	}
	return nil
}

func TestEvent(t *testing.T) {
	e := &event.BasicEvent{}
	e.SetName("n1")
	e.SetData(event.M{
		"arg0": "val0",
	})
	e.SetTarget("tgt")

	e.Add("arg1", "val1")

	assert.False(t, e.IsAborted())
	e.Abort(true)
	assert.True(t, e.IsAborted())

	assert.Equal(t, "n1", e.Name())
	assert.Equal(t, "tgt", e.Target())
	assert.Contains(t, e.Data(), "arg1")
	assert.Equal(t, "val0", e.Get("arg0"))
	assert.Equal(t, nil, e.Get("not-exist"))

	e.Set("arg1", "new val")
	assert.Equal(t, "new val", e.Get("arg1"))

	e1 := &event.BasicEvent{}
	e1.Set("k", "v")
	assert.Equal(t, "v", e1.Get("k"))
}
