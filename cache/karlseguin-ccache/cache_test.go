package karlseguin_ccache

import (
	"testing"
	"time"

	"github.com/karlseguin/ccache/v3"
	"github.com/karlseguin/ccache/v3/assert"
)

//https://github.com/karlseguin/ccache/blob/master/cache_test.go

func Test_CacheDeletesAValue(t *testing.T) {
	cache := ccache.New(ccache.Configure[string]())
	defer cache.Stop()
	assert.Equal(t, cache.ItemCount(), 0)

	cache.Set("spice", "flow", time.Minute)
	cache.Set("worm", "sand", time.Minute)
	assert.Equal(t, cache.ItemCount(), 2)

	cache.Delete("spice")
	assert.Equal(t, cache.Get("spice"), nil)
	assert.Equal(t, cache.Get("worm").Value(), "sand")
	assert.Equal(t, cache.ItemCount(), 1)
}
