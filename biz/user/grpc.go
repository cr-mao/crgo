package user

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

//rpc GetUserList(PageInfo) returns (UserListResonse); //用户列表
//rpc GetUserByMobile(MobileRequest) returns (UserInfoResponse); //通过mobile查询用户
//rpc GetUserById(IdRequest) returns (UserInfoResponse); //通过id查询用户
//rpc CreateUser(CreateUserInfo) returns (UserInfoResponse); //添加用户
//rpc UpdateUser(UpdateUserInfo) returns (google.protobuf.Empty); // 更新用户
//rpc CheckPassWord (PasswordCheckInfo) returns (CheckResponse); //检查密码

func (s *UserService) GetUserList(ctx context.Context, in *PageInfo) (*UserListResonse, error) {
	fmt.Println(in.Pn)
	fmt.Println(in.PSize)
	return &UserListResonse{}, nil
}

func (s *UserService) GetUserByMobile(ctx context.Context, in *MobileRequest) (*UserInfoResponse, error) {
	return &UserInfoResponse{}, nil
}
func (s *UserService) GetUserById(ctx context.Context, in *IdRequest) (*UserInfoResponse, error) {
	return &UserInfoResponse{}, nil
}
func (s *UserService) CreateUser(ctx context.Context, in *CreateUserInfo) (*UserInfoResponse, error) {
	return &UserInfoResponse{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, in *UpdateUserInfo) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) CheckPassWord(ctx context.Context, in *PasswordCheckInfo) (*CheckResponse, error) {
	return &CheckResponse{}, nil
}
