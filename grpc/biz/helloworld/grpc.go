package helloworld

import (
	"context"
	"crgo/infra/code"
	"crgo/infra/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Binding struct {
	UnimplementedGreeterServer
}

func (s *Binding) SayHello(ctx context.Context, r *HelloRequest) (*HelloResponse, error) {
	e := errors.WithCode(code.ErrUserNotFound, "user not found")
	// 内部错误 转为grpc 错误
	return nil, errors.ToGrpcError(e)
	// grpc 错误
	return nil, status.Error(codes.NotFound, " HELLO NOT FOUND")

	return &HelloResponse{
		Message: "hello " + r.Name,
	}, nil
}
