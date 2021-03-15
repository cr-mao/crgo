package job

import "context"


//任务
type Task func(context.Context) error

// 中间件， task 的前后的工作
type TaskInterceptor func(Task) Task


// Job 是会被调度执行的进程
// + Once Job 只执行一次 Task，然后退出
// + Daemon Job 按照给定的频率，持续执行 Task
type Job interface {
	Run(context.Context) error
}


