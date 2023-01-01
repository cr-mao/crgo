package util

import (
	"crgo/infra/errcode"
	"crgo/infra/response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//grcp 错误转http
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				response.ErrorAbort(c, errcode.ErrCodes.ErrNotFound)
			case codes.Internal:
				response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer, "内部错误")
			case codes.InvalidArgument:
				response.ErrorAbort(c, errcode.ErrCodes.ErrParams)
			case codes.Unavailable:
				response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer, "用户服务不可用")
			default:
				response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer, e.Err().Error())
			}
			return
		}
	}
}
