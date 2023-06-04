package bluele_gcache__test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bluele/gcache"
)

func TestSetGet(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}

func TestExpire(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.SetWithExpire("key", "ok", time.Second*2)
	value, _ := gc.Get("key")
	fmt.Println("Get:", value)

	// Wait for value to expire
	time.Sleep(time.Second * 3)

	value, err := gc.Get("key")
	if err != nil && err != gcache.KeyNotFoundError {
		panic(err)
	} else if err == gcache.KeyNotFoundError {
		fmt.Println("key not found")
	} else {
		fmt.Println("Get:", value)
	}
}

func TestLoaderFunc(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return "ok", nil
		}).
		Build()
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}

func TestLoaderExpireFunc(t *testing.T) {
	var evictCounter, loaderCounter, purgeCounter int
	gc := gcache.New(20).
		LRU().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			loaderCounter++
			expire := 1 * time.Second
			return "ok", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			evictCounter++
			fmt.Println("evicted key:", key)
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			purgeCounter++
			fmt.Println("purged key:", key)
		}).
		Build()

	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)

	time.Sleep(1 * time.Second)
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)

	gc.Purge()
	if loaderCounter != evictCounter+purgeCounter {
		panic("bad")
	}
}

func TestARC(t *testing.T) {
	gc := gcache.New(10).
		ARC().
		Build()

	gc.Set("key", "value")
	val, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", val)
}

func TestSimpleCache(t *testing.T) {
	//SimpleCache has no clear priority for evict cache. It depends on key-value map order.
	gc := gcache.New(10).Build()
	gc.Set("key", "value")
	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", v)
}
