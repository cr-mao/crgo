package current

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/peer"

	"crgo/infra/conf"
)

// 用户ip
func IP(ctx context.Context) string {
	if conf.IsDev() {
		if p, ok := peer.FromContext(ctx); ok {
			return p.Addr.String()
		}
		return ""
	}

	md := metautils.ExtractIncoming(ctx)
	// TODO: 非开发环境，通过 Nginx 设置 X-Real-IP 来获取用户真实 IP
	return md.Get("X-Real-IP")
}
