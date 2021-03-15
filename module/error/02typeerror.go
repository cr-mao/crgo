package main

import "fmt"

// before go 1.13


type NotFoundError struct {
	Name string
}


func (e *NotFoundError) Error() string {
	return e.Name + ": not found"
}


func SetName(name string ) error {
	return &NotFoundError{Name:name}
}



func main(){
	err :=SetName("crmao")
	if e,ok :=err.(*NotFoundError);ok {
		fmt.Println(e)  // crmo not found
	}
}


