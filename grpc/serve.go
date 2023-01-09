package grpc

import (
	"crgo/biz/auth"
	"crgo/biz/goods"
	"crgo/biz/inventory"
	"crgo/biz/session"
	"crgo/biz/user"
	"crgo/grpc/biz/bootstrap"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	//注册服务健康检查
	// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s
}

func registerService(s *grpc.Server) {
	helloworld.RegisterGreeterServer(s, &helloworld.Binding{})
	bootstrap.RegisterBootstrapServer(s, bootstrap.NewService(session.NewService()))

	// 暂时不区分 服务，都统一在一个服务里面， 知道怎么玩就行
	// 用户服务，
	user.RegisterUserServer(s, user.NewUserService())
	//商品服务
	goods.RegisterGoodsServer(s, goods.NewGoodsService())
	//库存服务
	inventory.RegisterInventoryServer(s, inventory.NewInventoryService())
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
