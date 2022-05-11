package main

import "fmt"

func main1() {

	func() {
		defer func() {
			fmt.Println("1")
		}()

		defer func() {
			fmt.Println("2")
		}()

		defer func() {
			fmt.Println("3")
		}()

		panic("异常1")
	}()

	panic("异常2")


  //3,2,1 panic：异常1
}



func main() {

	defer func() {
		defer func() {
			fmt.Println("1")
		}()

		defer func() {
			fmt.Println("2")
		}()

		defer func() {
			fmt.Println("3")
		}()

		panic("异常1")
	}()

	panic("异常2")


	//3,2,1 panic：异常2 ,panic: 异常1
}
