package auth

import (
	"crgo/http/models/user"
	"crgo/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignupController struct {
}

// 请求对象
type PhoneExistRequest struct {
	Phone string `json:"phone"`
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
