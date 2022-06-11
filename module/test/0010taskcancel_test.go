package test

import (
	"fmt"
	"testing"
	"time"
)

// 消息发送 取消
func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//广播
func cancel2(chanChan chan struct{}) {
	close(chanChan)
}

func doTask(i int, cancelChan chan struct{}) {

	for {
		if isCanceled(cancelChan) {
			break
		}
		time.Sleep(time.Millisecond * 100)

	}
	fmt.Println(i, "canceled")
}

func TestCancel(t *testing.T) {
	var cancelChan = make(chan struct{})
	for i := 0; i < 10; i++ {
		go doTask(i, cancelChan)
	}
	cancel2(cancelChan)
	time.Sleep(time.Second * 2)
}
