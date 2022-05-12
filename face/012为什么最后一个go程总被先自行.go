package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 设置一核，
// 循环创建多个协程
// p就绪后，默认先执行最后一个协程
//然后再继续执行其他协成，按顺序执行

func main() {
	runtime.GOMAXPROCS(1)
	//这个程序输出顺序是固定的
	wg := sync.WaitGroup{}
	wg.Add(6)

	for i := 0; i < 5; i++ {
		go func(input int) {
			defer wg.Done()
			fmt.Println(input)
		}(i)
	}
	go func() {
		defer wg.Done()
		fmt.Println("我要开始执行了")
	}()
	wg.Wait()
}
