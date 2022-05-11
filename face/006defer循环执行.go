package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	//3,3,3
}
