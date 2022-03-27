package lib1

import (
	"bou.ke/monkey"
	"runtime"
	"testing"
	"time"
)

func no() bool  { return false }
func yes() bool { return true }

func TestTimePatch(t *testing.T) {
	defer monkey.UnpatchAll()
	before := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	})
	during := time.Now()
	assert(t, monkey.Unpatch(time.Now))
	after := time.Now()

	assert(t, time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC) == during)
	assert(t, before != during)
	assert(t, during != after)
}

func TestGC(t *testing.T) {
	value := true
	monkey.Patch(no, func() bool {
		return value
	})
	defer monkey.UnpatchAll()
	runtime.GC()
	assert(t, no())
}

func TestSimple(t *testing.T) {
	defer monkey.UnpatchAll()
	monkey.Patch(no, func() bool {
		return true
	})
	assert(t, no())
	assert(t, monkey.Unpatch(no))
	assert(t, !no())
	assert(t, !monkey.Unpatch(no))
}

func TestUnpatchAll(t *testing.T) {
	assert(t, !no())
	monkey.Patch(no, yes)
	assert(t, no())
	monkey.UnpatchAll()
	assert(t, !no())
}

func assert(t *testing.T, b bool, args ...interface{}) {
	t.Helper()
	if !b {
		t.Fatal(append([]interface{}{"assertion failed"}, args...))
	}
}

func panics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if v := recover(); v == nil {
			t.Fatal("expected panic")
		}
	}()
	f()
}
