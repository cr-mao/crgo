package main

import (
	"fmt"
	"sync"
	"time"
)

// 任务
func job(index int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("执行完毕,序号:%d\n", index)
}
func main() {
	wg := sync.WaitGroup{}
	//利用 缓冲区chan  控制速度
	myPool := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)

		myPool <- struct{}{}
		go func(index int) {
			defer wg.Done()
			defer func() {
				<-myPool
			}()
			job(index)

		}(i)
	}
	wg.Wait()

}
