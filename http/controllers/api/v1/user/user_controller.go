package user

import (
	"context"
	"crgo/http/global"
	"crgo/http/middleware"
	"crgo/http/util"
	"crgo/infra/errcode"
	"crgo/infra/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"time"

	protoUser "crgo/biz/user"
	"crgo/infra/log"
)

type UserController struct {
}
type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	//Birthday string `json:"birthday"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}

type UserListReponse struct {
	Total int32
	Data  []UserResponse
}

//假的登录
func (cc *UserController) Login(c *gin.Context) {
	//假设账号密码 输入正确 用户id是1 ，昵称是crmao
	//生成token
	id := uint(1)
	nickname := "crmao"
	j := middleware.NewJWT()
	claims := global.CustomClaims{
		ID:          id,
		NickName:    nickname,
		AuthorityId: uint(0),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "crmao",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer, "生成token失败")
		return
	}
	response.Success(c, errcode.ErrCodes.ErrNo, map[string]interface{}{
		"id":         id,
		"nick_name":  nickname,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

//获得用户列表
func (cc *UserController) GetUserList(c *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Errorf("client conn err :%v", err)
	}
	defer conn.Close()

	client := protoUser.NewUserClient(conn)
	userResponse, err := client.GetUserList(context.Background(), &protoUser.PageInfo{
		Pn:    0,
		PSize: 20,
	})
	if err != nil {
		log.Errorf("get userlist  grpc error:%v", err)
		util.HandleGrpcErrorToHttp(err, c)
		return
	}

	var resp = make([]UserResponse, 0)
	for _, value := range userResponse.Data {
		resp = append(resp, UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		})
	}

	var respData UserListReponse
	respData.Total = userResponse.Total
	respData.Data = resp

	response.Success(c, errcode.ErrCodes.ErrNo, respData)
	return
}
