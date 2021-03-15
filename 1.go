package main

import (
	"fmt"
	"time"
)

func main() {
	var t1  = time.NewTicker(time.Second*2)
	for {
		select {
		case <- t1.C:
			time.Sleep(time.Second * 10)
			fmt.Println(time.Now().Format("01-02 15:04:05"))
		}
	}
}
