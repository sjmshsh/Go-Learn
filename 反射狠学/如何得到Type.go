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
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI) // int
	fmt.Println(typeS) // string

	typeUser := reflect.TypeOf(&User{})
	fmt.Println(typeUser)               // *main.User
	fmt.Println(typeUser.Kind())        // ptr
	fmt.Println(typeUser.Elem().Kind()) // struct
}
