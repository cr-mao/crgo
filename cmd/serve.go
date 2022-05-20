package cmd

import (
	"context"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"crgo/grpc"
	"crgo/http"
	"crgo/infra/log"
	"crgo/infra/util"
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
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	httpServe := http.NewServe()
	grpcServe := grpc.NewGrpcServe()
	grpcListener := grpc.NewListen()
	go func() {
		if err := httpServe.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
			log.Errorf("http ListenAndServe error %v", err)
			quit <- syscall.SIGINT
		}
	}()
	go func() {
		if err := grpcServe.Serve(grpcListener); err != nil {
			log.Errorf("grpc Serve err: %v", err)
			quit <- syscall.SIGINT
		}
	}()
	//更新黑名单 内存
	go func() {
		util.WatchBlacklist()
	}()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServe.Shutdown(ctx); err != nil {
		log.Warnf("http server forced to shutdown:%v", err)
	}
	log.Info("HTTP Server exit.")
	// shutdown grpc server
	grpcServe.GracefulStop()
	log.Info("gRPC Server exit.")
	return nil
}
