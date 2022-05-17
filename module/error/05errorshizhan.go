package main

import (
	"errors"
	"fmt"
	perrors "github.com/pkg/errors"
)

// error 工程实战

var my = errors.New("my")

func main() {
	err := test02()
	fmt.Printf("%+v", err)

}

func test02() error {
	return test01()
}

func test01() error {
	return test0()
}

func test0() error {
	return perrors.Wrapf(my, "test02 failed")

}
