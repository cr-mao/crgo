package test

//  函数前 调用特定同类型函数  ，
import (
	"testing"
)

type task func() error

type interceptors func(task) task

// b,c,a
func TestRunInterceptor(t *testing.T) {
	var a task = func() error {
		t.Log("a")
		return nil
	}
	var b interceptors = func(a task) task {
		t.Log("b")
		return a
	}
	var c interceptors = func(a task) task {
		t.Log("c")
		return a
	}
	interceptor := []interceptors{
		b, c,
	}
	var d task
	for _, i := range interceptor {
		d = i(a)
	}
	d()

}
