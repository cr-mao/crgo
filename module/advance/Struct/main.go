package main

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Id   int
	Age  int
	Name string `name:"name"`
}

type Iservice interface {
	Save(data interface{}) Iservice
	List() Iservice
}

type UserService struct {
}

type ProdService struct {
}

func NewProdService() *ProdService {
	return &ProdService{}
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) List() Iservice {
	log.Println("UserService list")
	return this
}

func (this *UserService) Save(data interface{}) Iservice {
	log.Println("UserService Save")
	if user, ok := data.(*User); ok {
		log.Println(user.Name)
	} else {
		log.Fatal("用户参数错误")
	}
	return this
}

func (this *ProdService) List() Iservice {
	log.Println("ProdService  List")

	return this
}

func (this *ProdService) Save(data interface{}) Iservice {
	log.Println("ProdService  Save")

	return this
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
	//fmt.Println(u)
	users := NewUserService()
	users.Save(u)

	t := reflect.TypeOf(u) //类型反射
	if t.Kind() == reflect.Ptr {
		t = t.Elem() //把t 变成了 指针指向的变量 ，不用这个函数只能针对结构体
	}
	// fmt.Println(t.Name())  // User
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, t.Field(i).Type, t.Field(i).Tag)
	}

	t1 := reflect.ValueOf(u)

	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem() //把t 变成了 指针指向的变量 ，不用这个函数只能针对结构体
	}

	for i := 0; i < t1.NumField(); i++ {
		if t1.Field(i).Kind() == reflect.Int {
			fmt.Println(t1.Field(i).Int())
		}
		if t1.Field(i).Kind() == reflect.String {
			fmt.Println(t1.Field(i).String())
		}
		fmt.Println(t1.Field(i).Interface()) //只要值  用这个就行
	}

	//利用反射设置stuct 属性值
	t2 := reflect.ValueOf(u)
	if t2.Kind() == reflect.Ptr {
		t2 = t2.Elem() //把t 变成了 指针指向的变量 ，不用这个函数只能针对结构体
	}
	for i := 0; i < t2.NumField(); i++ {
		if t2.Field(i).Kind() == reflect.Int {
			//通用set 方式
			t2.Field(i).Set(reflect.ValueOf(100))
		}
		if t1.Field(i).Kind() == reflect.String {
			t2.Field(i).SetString("maodada")
		}
	}
	//打印修改属性值 后的 u
	fmt.Println(u) // &{100 100 maodada}
	//通过切片 隐射  改变 结构体的属性值
	s1 := []interface{}{
		1, 2, "mao",
	}
	t3 := reflect.ValueOf(u)
	if t3.Kind() == reflect.Ptr {
		t3 = t3.Elem() //把t 变成了 指针指向的变量 ，不用这个函数只能针对结构体
	}
	for i := 0; i < t3.NumField(); i++ {
		if t3.Field(i).Kind() == reflect.ValueOf(s1[i]).Kind() {
			//通用set 方式
			t3.Field(i).Set(reflect.ValueOf(s1[i]))
		}
	}
	fmt.Println(u)
	//map 转结构体
	u1 := &User{}
	m := map[string]interface{}{
		"Id":         2,
		"age":        "1", //类型也是要一致的
		"name":       "mdd",
		"otherfield": "xxx",
	}

	Map2Struct(m, u1)
	fmt.Println(u1) // &{2 0 mdd}
}

func Map2Struct(m map[string]interface{}, u interface{}) {
	v := reflect.ValueOf(u)
	if v.Kind() != reflect.Ptr {
		panic("u 必须是指针")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		panic("必须是结构体指针")
	}

	findFromMap := func(key string, nameTag string) interface{} {
		for k, v1 := range m {
			//只要属性key  或者  属性的tag 的name 值  和map 的  k相等  就取值。
			if k == key || nameTag == k {
				return v1
			}
		}
		return nil
	}

	for i := 0; i < v.NumField(); i++ {
		//value 可以直接转type
		fmt.Println(v.Type().Field(i).Tag.Get("name"))
		getValue := findFromMap(v.Type().Field(i).Name, v.Type().Field(i).Tag.Get("name"))
		if getValue != nil && reflect.ValueOf(getValue).Kind() == v.Field(i).Kind() {
			v.Field(i).Set(reflect.ValueOf(getValue))
		}
	}
}
