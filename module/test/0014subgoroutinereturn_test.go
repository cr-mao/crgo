package test

import (
	"context"
	"fmt"
	"testing"
)

// 父协成退出，子协成退出 todo
func TestSubGoroutineReturn(t *testing.T) {

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func(ctx context.Context) {

			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("do task .....")

				}
			}

		}(ctx)
	}()

}
