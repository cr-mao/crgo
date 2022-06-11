package test

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	var p *int
	fmt.Println(p)  // nil
	t.Logf("%p", p) //0x0

	//fmt.Println(*p) // 报错

	p = new(int)
	fmt.Println(*p) //0
}
