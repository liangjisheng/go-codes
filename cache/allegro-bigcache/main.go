package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

func defaultConfig() {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(2*time.Second))

	go func() {
		for i := 0; i < 5; i++ {
			cache.Set("my-unique-key", []byte("value"))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		if entry, err := cache.Get("my-unique-key"); err == nil {
			fmt.Println(string(entry))
		}
		time.Sleep(time.Second)
	}
}

func customConfig() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 2 * time.Second,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 2 * time.Second,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	}

	cache1, initErr := bigcache.New(context.Background(), config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	cache1.Set("my-unique-key", []byte("value"))

	for i := 0; i < 6; i++ {
		if entry, err := cache1.Get("my-unique-key"); err == nil {
			fmt.Println(string(entry))
		}
		time.Sleep(time.Second)
	}
}

func concurrent() {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Second))

	go func() {
		for {
			cache.Set("key", []byte("value"))
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for {
			cache.Set("key", []byte("value1"))
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(10 * time.Millisecond)

	go func() {
		for {
			if entry, err := cache.Get("key"); err == nil {
				fmt.Println("first", string(entry))
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		if entry, err := cache.Get("key"); err == nil {
			fmt.Println("second", string(entry))
		}
		time.Sleep(time.Second)
	}
}

func main() {
	//defaultConfig()
	//customConfig()
	concurrent()
}
