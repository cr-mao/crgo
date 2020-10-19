package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// 1秒的1次性定时器
	t1 := time.NewTimer(time.Second)
	go func() {
		for {
			<-t1.C
			fmt.Println("子go程，定时完毕") //只打印一次
		}
	}()

	for {

	}

}



