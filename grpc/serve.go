package grpc

import (
	"crgo/infra/conf"
	"google.golang.org/grpc"
	"net"
	"strings"

	"crgo/grpc/biz/helloworld"
	"crgo/infra/log"
)

func NewGrpcServe() *grpc.Server {
	var grpcOption = []grpc.ServerOption{
		grpc.UnaryInterceptor(HelloInterceptor),
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
