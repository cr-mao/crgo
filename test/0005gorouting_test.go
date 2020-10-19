package test

import (
	"sync"
	"testing"
)

func TestCountWaitGroup(t *testing.T) {

	var wg sync.WaitGroup
	var lock sync.Mutex
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				lock.Unlock()
			}()
			lock.Lock()
			count++
			wg.Done()

		}()
	}
	wg.Wait()
	t.Log(count)

}
