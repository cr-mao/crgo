package main

import (
	"fmt"
	"time"
)

func main() {
Label:
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("超时")
			return
		default:
			break Label
		}
	}

}
