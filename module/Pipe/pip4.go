package main

import (
	"fmt"
	"sync"
	"time"
)

//管道只要 写 和 读 在不同协成执行即可

type CmdFunc func([]int) chan int

type PipeCmd func(in chan int) chan int

func EvenFunc(list []int) chan int {
	c1 := make(chan int)
	go func() {
		defer close(c1)
		for _, num := range list {
			if num%2 == 0 {
				c1 <- num
			}

		}
	}()
	return c1
}

func M10(in chan int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for num := range in {
			time.Sleep(time.Second)
			c <- num * 10
		}
	}()
	return c

}

func M2(in chan int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for num := range in {
			c <- num * 2
		}
	}()
	return c

}

//多路复用
func Pipe2(args []int, c1 CmdFunc, c2 ...PipeCmd) chan int {
	res := c1(args)
	if len(c2) == 0 {
		return res
	}
	var wg sync.WaitGroup
	var out = make(chan int)
	for _, item := range c2 {

		//其实 就是  只有一个 res 的 通道 (个数固定），  抢占式消费 ，
		getChan := item(res)
		wg.Add(1)
		go func(input chan int) {
			defer wg.Done()
			for it := range input {
				out <- it
			}

		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()
	}()
	return out
}

func main() {
	nums := []int{2, 3, 6, 12, 22, 16, 4, 9, 23, 64, 62}
	ret := Pipe2(nums, EvenFunc, M10, M10, M10, M10)

	for r := range ret {
		fmt.Println(r)
	}
}
