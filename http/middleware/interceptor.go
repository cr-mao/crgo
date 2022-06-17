package middleware

import (
	"context"
	redis "crgo/infra/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		cookie, err := c.Cookie("SOCIAL_SID")
		if cookie == "" || err != nil {
			log.Println(err)
		}
		//session
		res, err := redis.Client("session").Get(context.Background(), "session_info_"+cookie).Result()
		if err == redis.Nil {
			// 错误
			c.JSON(http.StatusUnauthorized, gin.H{
				"err_code": 40101,
				"err_msg":  "未登陆",
				"payload":  "",
			})
			c.Abort()  // 终止当前请求的处理函数调用链
			return

		} else if err != nil {
			panic(err)
		} else {
			if res != "1" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"err_code": 40101,
					"err_msg":  "未登陆",
					"payload":  "",
				})
				c.Abort()  // 终止当前请求的处理函数调用链
				return
			}
			c.Next()
		}
	}

}
