package middleware

import (
	"bytes"
	"crgo/infra/log"
	"crgo/infra/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 response 内容
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		// 设置开始时间
		start := time.Now()
		c.Next()
		// 开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := c.Writer.Status()
		logField := map[string]interface{}{
			"status":     responStatus,
			"request":    c.Request.Method + " " + c.Request.URL.String(),
			"query":      c.Request.URL.RawQuery,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"errors":     c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"time":       util.MicrosecondsStr(cost),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求的内容
			logField["RequestBody"] = string(requestBody)
			// 响应的内容
			logField["ResponseBody"] = w.body.String()
		}

		if responStatus > 400 && responStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			log.Warn("HTTP Warning "+cast.ToString(responStatus), logField)
		} else if responStatus >= 500 && responStatus <= 599 {
			// 除了内部错误，记录 error
			log.Error("HTTP Error "+cast.ToString(responStatus), logField)
		} else {
			log.Debug("HTTP Access Log", logField)
		}
	}
}
