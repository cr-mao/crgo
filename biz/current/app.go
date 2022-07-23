package current

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

const (
	CrgoVersion  = "crgo-version"
	CrgoBundleID = "crgo-bid"
)

func AppVersion(ctx context.Context) string {
	md := metautils.ExtractIncoming(ctx)
	return md.Get(CrgoVersion)
}

// ios 每个app 都有一个BundleId ,不然就是非法的 app
func AppBundleID(ctx context.Context) string {
	md := metautils.ExtractIncoming(ctx)
	return md.Get(CrgoBundleID)
}
