package groupcache__test

import (
	"context"
	"github.com/golang/groupcache"
	"sync"
	"testing"
)

var (
	once                    sync.Once
	stringGroup, protoGroup groupcache.Getter

	stringc = make(chan string)

	dummyCtx = context.TODO()

	// cacheFills is the number of times stringGroup or
	// protoGroup's Getter have been called. Read using the
	// cacheFills function.
	cacheFills groupcache.AtomicInt
)

const (
	stringGroupName = "string-group"
	protoGroupName  = "proto-group"
	testMessageType = "google3/net/groupcache/go/test_proto.TestMessage"
	fromChan        = "from-chan"
	cacheSize       = 1 << 20
)

func testSetup() {
	stringGroup = groupcache.NewGroup(stringGroupName, cacheSize, groupcache.GetterFunc(func(_ context.Context, key string, dest groupcache.Sink) error {
		if key == fromChan {
			key = <-stringc
		}
		cacheFills.Add(1)
		return dest.SetString("ECHO:" + key)
	}))
}

func countFills(f func()) int64 {
	fills0 := cacheFills.Get()
	f()
	return cacheFills.Get() - fills0
}

func TestCaching(t *testing.T) {
	once.Do(testSetup)
	fills := countFills(func() {
		for i := 0; i < 10; i++ {
			var s string
			if err := stringGroup.Get(dummyCtx, "TestCaching-key", groupcache.StringSink(&s)); err != nil {
				t.Fatal(err)
			}
		}
	})
	if fills != 1 {
		t.Errorf("expected 1 cache fill; got %d", fills)
	}
}
