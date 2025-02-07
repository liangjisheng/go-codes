package runtime_demo

import "testing"

func TestGetGoroutineId(t *testing.T) {
	id, err := GetGoroutineId()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id)
}

func TestMemory(t *testing.T) {
	PointMemory()
}
