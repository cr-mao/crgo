package routers

import (
	"crgo/http/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()
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

	//
	//{
	//	authRouter :=router.Group("/").Use(Auth())
	//	authRouter.GET("/", func(context *gin.Context) {
	//		context.String(http.StatusOK, "hello gin")
	//	})
	//}

	{
		userRoute := router
		userRoute.POST("/users", Create)
		userRoute.GET("/users/:name", Get)
	}

	return router
}

var users []*dto.User

func Create(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": 10001})
		return
	}

	for _, u := range users {
		if u.Name == user.Name {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s already exist", user.Name), "code": 10001})
			return
		}
	}

	users = append(users, &user)
	c.JSON(http.StatusOK, user)
}

func Get(c *gin.Context) {
	username := c.Param("name")
	for _, u := range users {
		if u.Name == username {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("user %s not exist", username), "code": 10002})
}
