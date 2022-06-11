package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeticker(t *testing.T) {
	// 周期性定时器
	t1 := time.NewTicker(time.Second)
	go func() {
		for {
			<-t1.C
			fmt.Println(time.Now())
		}
	}()

	for {

	}

}
