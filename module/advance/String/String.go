package main

import "fmt"

type String string

func (this String) Len() int {
	return len(this)
}

//遍历字符串
func (this String) Each(f func(item string)) {
	//for i := 0; i < len(this); i++ {
	//	f(fmt.Sprintf("%c", this[i]))
	//}
	for _, v := range this {
		f(fmt.Sprintf("%c", v))
	}
}

func From(str string) String {
	return String(str)
}
func FromInt(a int) String {
	return String(fmt.Sprintf("%d", a))
}
func main() {
	s := From("我abc")
	//fmt.Println(s)
	//a := FromInt(123)
	//
	//fmt.Println(a)
	//fmt.Println(a.Len())

	s.Each(func(item string) {
		fmt.Println(item)
	})
}
