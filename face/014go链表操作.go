package main

import (
	"container/list"
	"fmt"
)

func main(){

	data :=list.New()
	data.PushBack(8)
	data.PushBack(9)
	data.PushBack(10)

	for e:=data.Front(); e!=nil;e=e.Next(){
		fmt.Println(e.Value)
	}

}