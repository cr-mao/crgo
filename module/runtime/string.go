package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "毛大1"
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
}
