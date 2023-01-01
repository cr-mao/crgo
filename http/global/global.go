package global

import (
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

var GrpcConnect *grpc.ClientConn
