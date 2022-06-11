package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {}

func  BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})
	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			v.Store(cfg)
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				cfg := v.Load().(*Config)
				cfg.T()
				//fmt.Printf("%v\n", cfg)
			}

			wg.Done()
		}()
	}
	wg.Wait()
}


func  BenchmarkMux(b *testing.B) {
	var lock sync.RWMutex
	var config *Config
	go func() {
		i := 0
		for {
			lock.Lock()
			i++
			config = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			lock.Unlock()
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				lock.RLock()
				config.T()
					//fmt.Println(config)
				lock.RUnlock()
			}

			wg.Done()
		}()
	}
	wg.Wait()
}

