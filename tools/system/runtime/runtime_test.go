package runtime

import (
	"testing"
)

func TestGetGoroutineId(t *testing.T) {
	id, err := GetGoroutineId()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id)
}
