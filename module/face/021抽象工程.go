package main

import "fmt"

//有个抽象工厂类，  任何的工厂有他自己的工厂类 用来创建资自身对象 ，这样就实现了 松耦合
type User interface {
	GetRole() string
}

type Admin struct {
}

type Member struct {
}

type AbstractFactory interface {
	CreateUser() User
}

type AdminFactory struct {
}

type MemberFactory struct {
}

func (this *Admin) GetRole() string {
	return "后台用户"
}

func (this *Member) GetRole() string {
	return "前台用户"
}

func (this *AdminFactory) CreateUser() User {
	return &Admin{}
}

func (this *MemberFactory) CreateUser() User {
	return &Member{}
}

func main() {

	var admin AbstractFactory = new(AdminFactory)
	fmt.Println(admin.CreateUser().GetRole())

	var member AbstractFactory = new(MemberFactory)
	fmt.Println(member.CreateUser().GetRole())
}
