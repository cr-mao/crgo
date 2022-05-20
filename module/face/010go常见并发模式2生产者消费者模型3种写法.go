package main

import (
	"fmt"
	"time"
)

func producer1(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i * 2
		time.Sleep(1 * time.Second)
	}
}

func consumer1(c chan int) {

	for item := range c {
		fmt.Println(item)
	}
}

func consumer2(c chan int, ret chan struct{}) {
	for item := range c {
		fmt.Println(item)
	}

	ret <- struct{}{}
}

func consumer3(c chan int, ret chan struct{}) chan struct{} {

	go func() {
		defer func() {
			defer close(ret)
			ret <- struct{}{}
		}()
		for item := range c {
			fmt.Println(item)
		}
	}()
	return ret
}

func main1() {
	var c = make(chan int)
	go producer1(c)
	consumer1(c)
}

func main2() {
	var c = make(chan int)
	var ret = make(chan struct{})
	go producer1(c)
	consumer3(c, ret)
	<-ret
}

func main() {
	var c = make(chan int)
	var ret = make(chan struct{})
	go producer1(c)
	go consumer2(c, ret)
	<-ret
}
