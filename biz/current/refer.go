package current

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

func RequestReferer(ctx context.Context) string {
	md := metautils.ExtractIncoming(ctx)
	return md.Get("referer")
}
