package main

import (
	"fmt"
)

func main() {

	a := 1

	defer fmt.Println(a)
	a++

	b := 1
	defer func() {
		fmt.Printf("b:%v\n", b) //2
	}()
	b++

	c := 1
	defer func(p *int) {
		fmt.Printf("c:%v\n", *p) //2
	}(&c)

	c++

}
