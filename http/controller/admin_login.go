package controller

import (
	"crgo/http/dto"
	"crgo/http/middleware"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct{}

func AdminRegisterRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)

}

// @Summary 后台登录
// @Product json
// @Param user_name body string true "用户名" minlength(3) maxlength(20)
// @Param password body string true "密码"  minlength(3)
// @Success 200 {object} Response "成功"
// @Failure 400 {object} Response
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {

	params := &dto.AdminLoginInput{}
	if err := c.ShouldBind(params); err != nil {
		middleware.ResponseError(c,400,err)
		return
	}
	middleware.ResponseSuccess(c,nil,"success")

}
