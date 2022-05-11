package main

import "fmt"

type test struct {
}

func NewTest() *test {
	return &test{}
}

func (this *test) do(a int) *test {
	fmt.Println(a)
	return this
}
func main() {
	t := NewTest()
	defer t.do(1).do(2).do(4) //defer 是对最后一个 函数do(4)会defer t.do(1).do(2)会先执行
	t.do(3)
}

/**
结果是1，2,3，,4
*/
