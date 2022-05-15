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

func Pipe(args []int, c1 CmdFunc, c2 ...PipeCmd) chan int {
	res := c1(args)
	if len(c2) == 0 {
		return res
	}
	for _, item := range c2 {
		res = item(res)
	}
	return res


	//ret:=c1(args)
	//if len(c2)==0{
	//	return ret
	//}
	//也可以利用切片    最后一个元素  就是返回的通道
	//retlist:=make([]chan int,0)
	//for index,c:=range c2{
	//	if index==0{
	//		retlist=append(retlist,c(ret))
	//	}else{
	//		getChan:=retlist[len(retlist)-1]
	//		retlist=append(retlist,c(getChan))
	//	}
	//}
	//return retlist[len(retlist)-1]
}

func main() {
	nums := []int{2, 3, 6, 12, 22, 16, 4, 9, 23, 64, 62}
	ret := Pipe(nums, EvenFunc, M10, M10, M2)

	for r := range ret {
		fmt.Println(r)
	}
}
