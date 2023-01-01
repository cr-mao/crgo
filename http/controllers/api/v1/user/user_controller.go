package user

import (
	"context"
	"crgo/http/util"
	"crgo/infra/errcode"
	"crgo/infra/response"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	protoUser "crgo/biz/user"
	"crgo/infra/log"
)

type UserController struct {
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
	response.Success(c, errcode.ErrCodes.ErrNo, userResponse)
	return
}
