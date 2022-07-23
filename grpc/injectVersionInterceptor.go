package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"crgo/infra"
)

//注入版本中间件
func injectVersionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		trailer := metadata.Pairs("crgo-version", infra.Version)
		grpc.SetTrailer(ctx, trailer)
		return handler(ctx, req)
	}
}
