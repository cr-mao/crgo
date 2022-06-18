package cmd

import (
	"context"
	"crgo/http/services"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"crgo/grpc"
	"crgo/http"
	"crgo/infra/log"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	Long:  desc,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := Run(); err != nil {
			log.Error(err)
		}
	},
}

//http , grpc服务一起启动
func Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	var errChan = make(chan error, 3)
	var stopChan = make(chan struct{})
	go func() {
		errChan <- HttpServe(stopChan)
	}()

	go func() {
		errChan <- GrpcServe(stopChan)
	}()

	//更新黑名单 内存
	go func() {
		blackService :=services.NewBlackService()
		blackService.WatchBlacklist()
	}()
	go func() {
		<-quit
		errChan <- errors.New("手动关闭")
	}()
	var stopped bool
	for i := 0; i < cap(errChan); i++ {
		if err := <-errChan; err != nil {
			log.Infof("shutdown error:%v", err)
		}
		if !stopped {
			stopped = true
			close(stopChan)
		}
	}

	time.Sleep(time.Second * 2)

	return nil
}

func HttpServe(stop <-chan struct{}) error {
	httpServe := http.NewServe()
	go func() {
		<-stop
		//可以给个超时关闭
		httpServe.Shutdown(context.Background())
	}()
	return httpServe.ListenAndServe()
}

func GrpcServe(stop <-chan struct{}) error {
	grpcServe := grpc.NewGrpcServe()
	grpcListener := grpc.NewListen()
	go func() {
		<-stop
		grpcServe.GracefulStop()
	}()

	return grpcServe.Serve(grpcListener)
}
