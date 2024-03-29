package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var mycounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ginapi_requests_total",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(mycounter)
}

//统计 方法的 请求数 到监控 prometheus
func MetheusPathCount() gin.HandlerFunc {
	return func(context *gin.Context) {
		mycounter.With(prometheus.Labels{
			"method": context.Request.Method,
			"path":   context.Request.RequestURI,
		}).Inc()
		context.Next()
	}
}
