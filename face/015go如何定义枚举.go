package main

import "fmt"

type UserType int

const (
	Student UserType = iota
	Teacher UserType = iota
)

func (this UserType) String() string {
	switch this {
	case 0:
		return "学生"
	case 1:
		return "老师"
	default:
		return ""
	}

}

func main() {
	fmt.Println(Student, Teacher)
}
