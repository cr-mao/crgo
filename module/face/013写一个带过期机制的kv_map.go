package main

import (
	"fmt"
	"sync"
	"time"
)

var kv sync.Map

func Set(key string, value interface{}, expire time.Duration) {
	kv.Store(key, value)
	time.AfterFunc(expire, func() {
		kv.Delete(key)
	})
}

func main() {
	Set("id", 101, time.Second*5)

	for {
		fmt.Println(kv.Load("id"))
		time.Sleep(time.Second)
	}
}
