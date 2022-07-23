package grpc

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
	"strings"

	"crgo/grpc/biz/helloworld"
	"crgo/infra/conf"
	"crgo/infra/log"
)

func NewGrpcServe() *grpc.Server {
	var grpcOption = []grpc.ServerOption{
		//grpc 默认不支持 多个函数中间件
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(

			//注入版本 到metadata
			injectVersionUnaryServerInterceptor(),
			//
			AuthUnaryServerInterceptor(),
		)),
	}

	s := grpc.NewServer(grpcOption...)

	helloworld.RegisterGreeterServer(s, &helloworld.Binding{})

	return s
}

func NewListen() net.Listener {
	listenAddress := strings.Join([]string{conf.GetString("grpc_addr"), conf.GetString("grpc_port")}, ":")

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	log.Infof("grpc listening on %s", listenAddress)
	return lis
}
