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
	typeUser := reflect.TypeOf(&User{})
	typeUser2 := reflect.TypeOf(User{})
	fmt.Println(typeUser.Elem()) // main.User
	fmt.Println(typeUser2) // main.User
}
