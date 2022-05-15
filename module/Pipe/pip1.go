package main

import "fmt"

type Cmd func(list []int) (ret []int)

func Evens(list []int) (ret []int) {
	ret = make([]int, 0)
	for _, num := range list {
		ret = append(ret, num)
	}
	return
}

func Multiply(list []int) (ret []int) {
	ret = make([]int, 0)
	for _, num := range list {
		if num%2 == 0 {
			ret = append(ret, num*10)
		}
	}
	return ret
}
func p(args []int, c1 Cmd, c2 Cmd) []int {
	ret := c1(args)
	return c2(ret)
}

func main() {
	nums := []int{2, 3, 6, 12, 22, 16, 4, 9, 23, 64, 62}

	fmt.Println(p(nums, Evens, Multiply))
}
