package cmd

import (
	"context"
	"crgo/infra/conf"
	"crgo/infra/discovery"
	"errors"
	"github.com/google/uuid"
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
		if err := Run(cmd.Context()); err != nil {
			log.Error(err)
		}
	},
}

//http , grpc服务一起启动
func Run(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	var errChan = make(chan error, 3)
	var stopChan = make(chan struct{})
	go func() {
		errChan <- HttpServe(ctx, stopChan)
	}()


	go func() {
		errChan <- GrpcServe(stopChan)
	}()
	//更新黑名单 内存
	//go func() {
	//	blackService :=services.NewBlackService()
	//	blackService.WatchBlacklist()
	//}()
	go func() {
		<-quit
		errChan <- errors.New("手动关闭")
	}()

	var consulClient *discovery.DiscoveryClient
	var instanceId string
	if conf.GetBool("consul_on",false ) {
		consulClient = discovery.NewDiscoveryClient(conf.GetString("consul_addr"), conf.GetInt("consul_port"))
		instanceId = "httpserve:"+ "-" + uuid.New().String()
		err := consulClient.Register(ctx, "httpserve", instanceId, "/health", conf.GetString("http_addr"), conf.GetInt("http_port"), nil, nil)
		if err != nil{
			log.Fatalf("register service err : %s", err)
		}
	}


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
	if conf.GetBool("consul_on",false ) {
		consulClient.Deregister(ctx, instanceId)
	}

	return nil
}

func HttpServe(ctx context.Context, stop <-chan struct{}) error {
	httpServe := http.NewServe()
	go func() {
		<-stop
		//可以给个超时关闭
		ctx, cancel := context.WithTimeout(ctx, time.Duration(2)*time.Second)
		defer cancel()
		httpServe.Shutdown(ctx)
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
