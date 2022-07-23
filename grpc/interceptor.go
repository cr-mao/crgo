package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"crgo/infra"
)

//拦截器

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")
	return resp, err
}

//注入版本中间件
func injectVersionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		trailer := metadata.Pairs("crgo-version", infra.Version)
		grpc.SetTrailer(ctx, trailer)
		return handler(ctx, req)
	}
}

// auth中间
func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		//if ctx.FromInComming

		return handler(ctx, req)
	}
}
