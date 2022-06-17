package middleware

import (
	"encoding/json"


	"github.com/gin-gonic/gin"
)


type Response struct {
	Code    int         `json:"errno"`
	Msg     string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code int, err error) {
	resp := &Response{Code: code, Msg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}, msg string) {
	resp := &Response{Code: 200, Msg: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
