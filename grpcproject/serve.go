package grpcproject

import (
	"google.golang.org/grpc"
	"net"

	"crgo/grpcproject/biz/helloworld"
	"crgo/infra/log"
)

func Run() error {
	var grpcOption = []grpc.ServerOption{
		grpc.UnaryInterceptor(HelloInterceptor),
	}
	s := grpc.NewServer(grpcOption...)
	helloworld.RegisterGreeterServer(s, &helloworld.Binding{})
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return err
	}
	err = s.Serve(lis)
	//listenAddress := strings.Join([]string{conf.Net.GRPC_ADDR, conf.Net.GRPC_PORT}, ":")
	//log.Infof("listening on %s", listenAddress)

	if err != nil {
		log.Errorf("failed to serve: %v", err)
		return err
	}
	return nil
}
