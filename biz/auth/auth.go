package auth

import (
	"crgo/biz/current"
	session2 "crgo/biz/session"
	"crgo/infra/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"context"
)

// grpc 代表需要登录请求的接口
type MessageWithGuid interface {
	GetGuid() string
}

// 无需看 认证的接口实现
type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (context.Context, error)
}

type AuthFunc func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (newCtx context.Context, err error)

//认证处理函数
func DefaultAuthFuc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (newCtx context.Context, err error) {
	//ios 客户端 是否有效验证
	//bundleID := current.AppBundleID(ctx)
	//if current.GetUserOS(ctx) == 1 && (bundleID == "" || !xxxValid(ctx, bundleID)) {
	//	return nil, status.Errorf(codes.PermissionDenied, "无效的客户端")
	//}
	if msg, ok := req.(MessageWithGuid); ok {
		if msg.GetGuid() == "" || current.SessionId(ctx) == "" {
			log.Warnf("auth failed. method: %s, request: `%s`, metadata: `%s`",
				info.FullMethod, req, metautils.ExtractIncoming(ctx))
			return nil, status.Errorf(codes.Unauthenticated, "认证失败，请重新登录")
		}
		// 验证 guid 与 sessionId 是否匹配
		sessionSvc := session2.NewService()
		sessionObj := sessionSvc.Get(ctx, current.SessionId(ctx))
		if sessionObj == nil || sessionObj.Anonymous || sessionObj.Guid != msg.GetGuid() {
			log.Warnf("auth failed. method: %s, request: `%s`, metadata: `%s`, session: %s",
				info.FullMethod, req, metautils.ExtractIncoming(ctx), sessionObj)
			return nil, status.Errorf(codes.Unauthenticated, "认证失败，请重新登录")
		}
		newCtx := session2.Load(ctx, sessionObj)
		// info.FullMethod  可以进行 特殊的方法处理
		var funcMap = map[string]int64{
			"/crgo.gateway.Greeter/SayHello": 1,
		}
		if _, ok := funcMap[info.FullMethod]; ok {
			//一些特殊请求特出处理逻辑
		}
		return newCtx, nil
	}
	return AnonymousAuthFunction(ctx, req, info)
}

// AnonymousAuthFunction 允许未登录用户请求 gRPC 接口
func AnonymousAuthFunction(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (context.Context, error) {
	// 仍然尝试找到对应的 session，不检查 guid 是否匹配
	sessionSvc := &session2.Service{}
	sessionObj := sessionSvc.Get(ctx, current.SessionId(ctx))
	if sessionObj != nil {
		return session2.Load(ctx, sessionObj), nil
	}
	return session2.Load(ctx, session2.Empty()), nil
}

// auth中间
func AuthUnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var newCtx context.Context
		var err error
		// 不走
		if srv, ok := info.Server.(ServiceAuthFuncOverride); ok {
			newCtx, err = srv.AuthFuncOverride(ctx, req, info)
		} else {
			newCtx, err = authFunc(ctx, req, info)
		}
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}
