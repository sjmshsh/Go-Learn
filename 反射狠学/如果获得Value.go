package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("hello")
	userPtrValue := reflect.ValueOf(&User{
		Name: "lxy",
		Age:  18,
	})
	fmt.Println(iValue)
	fmt.Println(sValue)
	fmt.Println(userPtrValue)
}
