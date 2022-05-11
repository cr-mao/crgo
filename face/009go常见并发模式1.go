package main

import (
	"fmt"
	"sync"
)

//利用waitGroup 控制
func main1() {

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			fmt.Println(index * 2)
		}(i)
	}

	wg.Wait()
}
//利用chan 并发控制
func main() {
	var c = make(chan int, 5)
	for i := 0; i < 5; i++ {
		go func(index int) {
			c <- index * 2
		}(i)
	}

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
}
