package auth

import (
	v1 "crgo/http/controllers/api/v1"
	"crgo/http/models/user"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type SignupController struct {
	v1.BaseAPIController
}

// 请求对象
type PhoneExistRequest struct {
	Phone string `json:"phone"`
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := PhoneExistRequest{}

	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
	}

	opts := govalidator.Options{
		Data:  &request,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		data, _ := json.MarshalIndent(e, "", "  ")
		fmt.Println(string(data))
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
