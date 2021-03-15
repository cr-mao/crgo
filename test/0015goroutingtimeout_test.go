package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type result struct {
	msg string
	err error
}

func search(term string) (string, error) {
	time.Sleep(time.Second * 3 )
	return "somevalue", nil
}

func Process(term string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	var resChan = make(chan result)
	defer cancel()
	go func() {
		res, err := search(term)
		resChan <- result{
			msg: res,
			err: err,
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
		return errors.New("timeout")
	case res := <-resChan:
		if res.err != nil {
			return res.err
		}
		fmt.Println(res.msg)
		return nil
	}
}

func TestSearch(t *testing.T) {
	Process("test")
}
