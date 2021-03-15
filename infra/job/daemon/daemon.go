package daemon

import (
	"context"
	"crgo/infra/job"
	"crgo/infra/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type daemon struct {
	work             job.Task
	rate             time.Duration
	taskInterceptors []job.TaskInterceptor
}

func NewDaemon(work job.Task, rate time.Duration, taskInterceptors []job.TaskInterceptor) *daemon {
	return &daemon{
		work:             work,
		rate:             rate,
		taskInterceptors: taskInterceptors,
	}
}

func (d *daemon) Run(ctx context.Context) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.Tick(d.rate)
	handler := d.work
	for _, i := range d.taskInterceptors {
		handler = i(handler)
	}

	for {
		select {
		case <-ticker:
			err := handler(ctx)
			// todo
			log.Error(err)
		case <-quit:


		}
	}

	return nil

}
