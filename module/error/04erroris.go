package main

import (
	"errors"
	"fmt"
)

// 哨兵错误处理
var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrNotFoundUser = errors.New("not found user")
)

func main() {
	err1 := fmt.Errorf("wrap err1: %w\n", ErrInvalidUser)
	err2 := fmt.Errorf("wrap err2: %w\n", err1)
	// golang 1.13 新增 Is() 函数
	if errors.Is(err2, ErrInvalidUser) {
		fmt.Println(ErrInvalidUser)  // invalid user
		return
	}
	fmt.Println("success")
}
