package middleware

import (
	"crgo/infra/conf"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// 基于 alibaba/sentinel 请求限流
func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e, b := sentinel.Entry(conf.GetString("sentinel_flow_resource", "request"), sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			// Blocked. We could get the block reason from the BlockError.
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
			ctx.AbortWithStatusJSON(500, gin.H{
				"code":    501,
				"msg":     b.Error(),
				"payload": "",
			})
			return
		} else {
			//停止，重新开始计算
			e.Exit()
			ctx.Next()
		}
	}
}
