package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main1() {
	var n int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n++

		}()
	}
	wg.Wait()
	fmt.Println(n) //不确定是多少 ，如何解决
}


//使用互斥锁  sync.Mutex解决
func main2() {
	var n int32
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			n++

		}()
	}
	wg.Wait()
	fmt.Println(n) //不确定是多少 ，如何解决
}

//使用atomic保证原子性
func main() {
	var n int32
	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)

		}()
	}
	wg.Wait()
	fmt.Println(n) //不确定是多少 ，如何解决
}
