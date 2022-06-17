package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"crgo/http/middleware"
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
	)
}

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
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
