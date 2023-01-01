package user

import (
	"context"
	"crgo/http/util"
	"crgo/infra/errcode"
	"crgo/infra/response"
	"fmt"
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
