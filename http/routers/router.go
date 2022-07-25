package routers

import (
	"crgo/infra/conf"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"net/http"
	"strings"

	"crgo/http/middleware"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/gin-gonic/gin"
)

// 404处理
func setup404Handler(r *gin.Engine) {
	// 添加 Get 请求路路由
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			c.String(http.StatusNotFound, "404 页面不存在")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 40401,
				"msg":  "请求路由不存在",
				"data": "",
			})
		}
	})
}

//全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.ErrorHandler(),
		middleware.RateLimit(),
		middleware.MetheusPathCount(),
	)
}

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}

func NewRouter() *gin.Engine {
	// sentinel 配置文件
	err := sentinel.InitWithConfigFile("sentinel.yaml")
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	// 请求频率直接拒绝
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               conf.GetString("sentinel_flow_resource", "request"),
			Threshold:              conf.GetFloat64("sentinel_flow_Threshold", 5000),
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       uint32(conf.GetUint("sentinel_stat_interval_in_ms", 1000)),
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	router := gin.New()
	gin.SetMode("debug")
	registerGlobalMiddleWare(router)
	setup404Handler(router)
	RegisterAPIRoutes(router)

	//router.LoadHTMLGlob("http/view/*")
	////v1 版本的分组
	//v1 :=router.Group("/v1")
	//{
	//	v1.GET("/upload", func(c *gin.Context) {
	//		c.HTML(http.StatusOK, "upload.html", gin.H{
	//			"title": "Main website",
	//		})
	//	})
	//
	//	adminLoginRoute :=v1.Group("/admin_login")
	//	controller.AdminRegisterRegister(adminLoginRoute)
	//}

	//
	//{
	//	authRouter :=router.Group("/").Use(Auth())
	//	authRouter.GET("/", func(context *gin.Context) {
	//		context.String(http.StatusOK, "hello gin")
	//	})
	//}

	//router.GET("/gorm_test", func(context *gin.Context) {
	//})

	return router

}
