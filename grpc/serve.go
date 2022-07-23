package grpc

import (
	"crgo/biz/auth"
	"crgo/biz/session"
	"crgo/grpc/biz/bootstrap"
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
			//用户认证处理
			auth.AuthUnaryServerInterceptor(auth.DefaultAuthFuc),
		)),
	}

	s := grpc.NewServer(grpcOption...)
	registerService(s)
	return s
}

func registerService(s *grpc.Server) {
	helloworld.RegisterGreeterServer(s, &helloworld.Binding{})
	bootstrap.RegisterBootstrapServer(s, bootstrap.NewService(session.NewService()))
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
