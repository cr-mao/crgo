package current

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
)

//获得用户设备  1 ios， 2: java  3：其他 ， 也可以用其他方式判断， 约定好即可
func GetUserOS(ctx context.Context) int {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return -1
	}
	uaSlice, ok := md["user-agent"]
	if !ok {
		return 3
	}
	//通过  user-agent  是有ios ，或者 grpc-objc 判断是否是 ios
	if strings.Contains(uaSlice[0], "ios") && strings.Contains(uaSlice[0], "grpc-objc") {
		return 1
	} else if strings.Contains(uaSlice[0], "grpc-java-okhttp") {
		return 2
	} else {
		return 3
	}
}
