package main

import "fmt"

type User interface {
	GetRole() string
}

type Admin struct {
}

type Member struct {
}

func (this *Admin) GetRole() string {
	return "后台用户"
}

func (this *Member) GetRole() string {
	return "前台用户"
}

const (
	AdminUser = iota
	MemberUser
)

//创建工厂 ，如果有新的类型要创建 就是要改 这个工程方法
func CreateUser(t int) User {
	switch t {
	case AdminUser:
		return new(Admin)
	case MemberUser:
		return new(Member)
	default:
		return new(Member)
	}
}

func main() {
	fmt.Println(CreateUser(MemberUser).GetRole())
}
