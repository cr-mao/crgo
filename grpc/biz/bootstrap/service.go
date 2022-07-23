package bootstrap

import (
	"context"

	"google.golang.org/grpc"

	"crgo/biz/auth"
	"crgo/biz/current"
	"crgo/biz/session"
	"crgo/grpc/biz/appsetting"
)

type Service struct {
	SessionSvc *session.Service
}

func NewService(sessionSvc *session.Service) *Service {
	return &Service{SessionSvc: sessionSvc}
}

func (s *Service) AuthFuncOverride(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo) (context.Context, error) {
	if info.FullMethod == "/crgo.gateway.Bootstrap/Bootstrap" {
		return auth.AnonymousAuthFunction(ctx, req, info)
	}
	return auth.DefaultAuthFuc(ctx, req, info)
}

func (s *Service) Bootstrap(ctx context.Context, in *BootstrapRequest) (*BootstrapResponse, error) {
	sessionId := in.GetSessionId()
	version := current.AppVersion(ctx)
	bundleID := current.AppBundleID(ctx)
	setting := appsetting.GetAppSetting(ctx, bundleID)
	reviewMode := setting.ReviewMode == 1 && version == setting.ReviewVersion
	if sessionId == "" || s.SessionSvc.IsInvalid(ctx, sessionId) {
		newSessionId := s.SessionSvc.GenerateAnonymous(ctx)
		return &BootstrapResponse{
			SessionId:  newSessionId,
			IsNew:      true,
			ReviewMode: reviewMode,
		}, nil
	}
	return &BootstrapResponse{
		SessionId:  sessionId,
		IsNew:      false,
		ReviewMode: reviewMode,
	}, nil
}
