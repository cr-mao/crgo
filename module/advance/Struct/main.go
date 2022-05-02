package main

import "fmt"

type User struct {
	Id, Age int
	Name    string
}

type UserFuncAttr func(*User)

type UserFuncAttrs []UserFuncAttr

func (this UserFuncAttrs) Apply(u *User) {
	for _, f := range this {
		f(u)
	}
}

func WithUserId(Id int) UserFuncAttr {
	return func(u *User) {
		u.Id = Id
	}
}

func WithUserName(name string) UserFuncAttr {
	return func(u *User) {
		u.Name = name
	}
}
func NewUser(fs ...UserFuncAttr) *User {
	u := new(User)
	UserFuncAttrs(fs).Apply(u)
	return u
}

//有选择性的对id  赋值
func main() {
	u := NewUser(WithUserId(12), WithUserName("MAOZHONGYU"))
	fmt.Println(u)
}
