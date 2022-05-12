package main

import (
	"fmt"
	"math/rand"
	"time"
)

func job() int {
	//随机数种子
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(result)) //模式任务延迟
	return result
}

func main() {
	var c = make(chan int, 5) //最大5个任务同时执行， 获取一个返回 就结束
	for i := 0; i < 5; i++ {
		go func() {
			c <- job()
		}()
	}
	fmt.Printf("最快的任务结果耗时:%d", <-c)
}
