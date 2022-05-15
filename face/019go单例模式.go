package main

import (
	"fmt"
	"sync"
)

type WebConfig struct {
	Port string
}

var config *WebConfig
var once sync.Once

func GetConfig() *WebConfig {
	once.Do(func() {
		config = &WebConfig{
			Port: "90",
		}
	})
	return config
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	fmt.Println(c1 == c2) //true
}
