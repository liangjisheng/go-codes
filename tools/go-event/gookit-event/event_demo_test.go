package event__test

import (
	"fmt"
	"github.com/gookit/event"
	"testing"
)

func TestDemo(t *testing.T) {
	// Register event listener
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event normal: %s\n", e.Name())
		fmt.Printf("event data: %+v\n", e.Data())
		return nil
	}), event.Normal)

	// Register multiple listeners
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event high: %s\n", e.Name())
		fmt.Printf("event data: %+v\n", e.Data())
		return nil
	}), event.High)

	// ... ...

	// Trigger event
	// Note: The second listener has a higher priority, so it will be executed first.
	event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})

	dbListener1 := event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		fmt.Printf("event data: %+v\n", e.Data())
		return nil
	})

	event.On("app.db.*", dbListener1, event.Normal)
	doCreate()
	doUpdate()
}

func doCreate() {
	// do something ...
	// Trigger event
	event.MustFire("app.db.create", event.M{"arg0": "val0", "arg1": "val1"})
}

func doUpdate() {
	// do something ...
	// Trigger event
	event.MustFire("app.db.update", event.M{"arg0": "val0"})
}

type MyEvent struct {
	event.BasicEvent
	customData string
}

func (e *MyEvent) CustomData() string {
	return e.customData
}

func TestDemo1(t *testing.T) {
	e := &MyEvent{customData: "hello"}
	e.SetName("e1")
	event.AddEvent(e)

	// add listener
	event.On("e1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("custom Data: %s\n", e.(*MyEvent).CustomData())
		return nil
	}))

	// trigger
	event.Fire("e1", nil)
	// OR
	// event.FireEvent(e)
}
