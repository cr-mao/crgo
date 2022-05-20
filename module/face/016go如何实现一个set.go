package main

import (
	"bytes"
	"fmt"
)

type Empty struct{}

type Set map[interface{}]Empty

func (this Set) Add(key interface{}) Set {
	this[key] = Empty{}
	return this
}

func NewSet() Set {
	return make(map[interface{}]Empty)
}
func (this Set) String() string {
	var buf bytes.Buffer
	for k, _ := range this {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", k))
	}
	return buf.String()
}

func main() {
	fmt.Println(NewSet().Add("a").Add("b"))
}
