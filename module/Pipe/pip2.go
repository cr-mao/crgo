package main

import "fmt"


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
			c <- num * 10
		}
	}()
	return c

}

func Pipe(args []int, c1 CmdFunc, c2 PipeCmd) chan int {
	res := c1(args)
	return c2(res)
}

func main() {
	nums := []int{2, 3, 6, 12, 22, 16, 4, 9, 23, 64, 62}
	ret := Pipe(nums, EvenFunc, M10)

	for r := range ret {
		fmt.Println(r)
	}
}
