package test

import "testing"

func TestMap(t *testing.T) {
	m := map[int]bool{}
	t.Log(m[1])     //false
	t.Log(m[1000])  //false
	delete(m, 1000) //不报错

	m1 := map[string]func(a int) int{}
	m1["test"] = func(a int) int { return a }
	t.Log(m1["test"](2))
}
