package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		time.Sleep(time.Second)
		return errors.New("test")
	})
	group.Go(func() error {
		time.Sleep(time.Second * 2)
		return nil
	})

	group.Go(func() error {
		time.Sleep(time.Second * 3)
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(ctx.Err())


}
