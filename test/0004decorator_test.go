package test

import (
	"fmt"
	"testing"
	"time"
)

//装饰器模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		t1 := time.Now()
		res := inner(op)
		fmt.Println(time.Since(t1).Seconds())
		return res
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestDecorator(t *testing.T) {
	fmt.Println(timeSpent(slowFunc)(3))
}
