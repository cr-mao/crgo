package routers

import (
	"crgo/http/controllers/api/v1/auth"
	"crgo/http/controllers/api/v1/user"
	"log"
	"net/http"
	"strings"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"crgo/http/middleware"
	"crgo/infra/conf"
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
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}

//全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Logger(),           //自定义请求响应中间件
		middleware.Recovery(),         //panic   错误 拦截处理
		middleware.RateLimit(),        //请求限流
		middleware.MetheusPathCount(), //请求方法 统计基数 监控
	)
}

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//consul 监控检测用
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// prom 用
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		signController := new(auth.SignupController)
		authGroup.POST("/signup/phone/exist", signController.IsPhoneExist)

		//用户api
		userController := &user.UserController{}
		v1.GET("/user_list", userController.GetUserList)
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
	//gin.SetMode(gin.ReleaseMode)
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

	return router

}
