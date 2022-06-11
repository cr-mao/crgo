package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGoTask(t *testing.T) {

	var closeChan = make(chan bool, 5)
	taskChan := make(chan int, 12)
	resChan := make(chan int, 12)

	//生成任务
	go func() {
		for i := 1; i <= 12; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	// 处理任务
	for i := 1; i <= 5; i++ {
		go Task(taskChan, resChan, closeChan)
	}
	// 判断是否执行完成，信息聚合
	go func() {
		for i := 1; i <= 5; i++ {
			<-closeChan
		}
		close(closeChan)
		close(resChan)
	}()

	for id := range resChan {
		fmt.Println("res=" + strconv.Itoa(id))
	}

}

func Task(taskChan chan int, resChan chan int, closeChan chan bool) {
	for taskId := range taskChan {
		resChan <- taskId * 2
	}
	closeChan <- true
}
