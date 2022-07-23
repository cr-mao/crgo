package once

import (
	"context"

	"crgo/infra/job"
)

type Once struct {
	task             job.Task
	taskInterceptors []job.TaskInterceptor
}

func (once *Once) Run(ctx context.Context) error {
	task := once.task
	for _, interceptor := range once.taskInterceptors {
		task = interceptor(task)
	}
	return task(ctx)
}

func NewOnce(t job.Task, i []job.TaskInterceptor) *Once {
	return &Once{
		task:             t,
		taskInterceptors: i,
	}
}












