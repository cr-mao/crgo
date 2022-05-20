package main

import (
	"fmt"
	"time"
)

//套路： 函数 协成化， 结果用 chan 获得
func job() chan string {
	c := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c <- "success"
	}()

	return c
}

func run() (interface{}, error) {
	select {
	case <-time.After(time.Second * 3):
		return nil, fmt.Errorf("time out")
	case r := <-job():
		return r, nil
	}

}
func main() {
	fmt.Println(run())
}
