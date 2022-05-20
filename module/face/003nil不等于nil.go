package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a func()
	var b *struct{}
	// interface 是有 值和类型 组成的 必须 值是nil ，类型也是nil 才会等于nil

	//可以通过接口断言  和 反射 来搞定
	c := []interface{}{
		a,
		b,
	}

	for _, item := range c {

		fmt.Printf("%T %v\n", item, item)

		if c, ok := item.(func()); ok && c == nil {
			fmt.Println(c)
		}

		if c, ok := item.(*struct{}); ok && c == nil {
			fmt.Println(c)
		}

		if reflect.ValueOf(item).IsNil() {
			fmt.Println(item)
		}

	}

}
