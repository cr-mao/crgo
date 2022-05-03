package main

import (
	"fmt"
	"sort"
)

// map 是无须的
type User map[string]interface{}

func NewUser() User {
	return make(map[string]interface{})
}

func (this User) With(k string, v interface{}) User {
	this[k] = v
	return this
}

//实现 Stringer 接口，改变打印格式
/**
type Stringer interface {
	String() string
}
*/
func (this User) String() string {
	str := ""
	for index, k := range this.Fields() {
		str += fmt.Sprintf("%d->%v->%v\n", index+1, k, this[k])
	}
	return str
}

/**
  map 的字段切片
*/
func (this User) Fields() []string {
	keys := []string{}
	for k, _ := range this {
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	//sort.Sort(sort.StringSlice(keys))
	//倒排
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return keys
}

func main() {
	u := NewUser().With("name", "mao").With("nickname", "cr-mao").With("id", 101)
	// 	fmt.Println(u)
	u1 := NewUser().With("name", "mao").With("nickname", "cr-mao").With("id", 102)
	u2 := NewUser().With("name", "mao").With("nickname", "cr-mao").With("id", 103)
	// // map 是无须的 ,多次打印  顺序会变
	// map 切片排序
	var users = []User{}
	users = append(users, u, u1, u2)
	sort.Slice(users, func(i, j int) bool {
		return users[i]["id"].(int) < users[j]["id"].(int)
	})
	fmt.Println(users)
}
