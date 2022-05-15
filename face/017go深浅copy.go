package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3}

	b := a
	a[1] = 4
	fmt.Println(b) //1,4,3    //浅copy

	var c = make([]int, 3)
	copy(c, a)               //深copy
	a[1] = 5
	fmt.Println(c) //1,4,3
}
