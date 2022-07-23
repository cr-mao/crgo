package current

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

const (
	//token
	USER_SID = "cr_session_id"
)

func User(ctx context.Context) {
}

//用户sessionID 就是所谓的token
func SessionId(ctx context.Context) string {
	md := metautils.ExtractIncoming(ctx)
	return md.Get(USER_SID)
}

////用户id
//func UserID(ctx context.Context) int64 {
//	return sessionBiz.Extract(ctx).UserID
//}
//
////用户唯一id
//func Guid(ctx context.Context) string {
//	return sessionBiz.Extract(ctx).Guid
//}
