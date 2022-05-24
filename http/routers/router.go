package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router :=gin.New()
	gin.SetMode("debug")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("http/view/*")
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title": "Main website",
		})
	})
	// 添加 Get 请求路路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"err_code": 40401,
			"err_msg":  "请求路由不存在",
			"payload":  "",
		})
	})


	{
		authRouter :=router.Group("/").Use(Auth())
		authRouter.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "hello gin")
		})
	}



	return router
}