package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Event struct{}


func (e *Event) List(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("event list")
	ctx.JSON(status, resp)
}
func (e *Event) Info(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("event info")
	ctx.JSON(status, resp)
}
func (e *Event) Subscribe(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("event subscribe")
	ctx.JSON(status, resp)
}