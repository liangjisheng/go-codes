package sync

import (
	"sync"
	"testing"
)

func TestNewConcurrentIntArray(t *testing.T) {
	length := 10
	array := NewConcurrentIntArray(length)

	wg := sync.WaitGroup{}

	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for count := 0; count < 500; count++ {
				for j := 0; j < length; j++ {
					_, err := array.Set(j, j)
					if err != nil {
						t.Error(err)
					}
				}
			}
		}()
	}

	wg.Wait()
}
