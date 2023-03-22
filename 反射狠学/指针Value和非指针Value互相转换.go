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
	userPtrValue := reflect.ValueOf(&User{
		Age:  7,
		Name: "杰克逊",
	})
	userValue := userPtrValue.Elem()                    // Elem()指针Value转为非指针Value
	fmt.Println(userValue.Kind(), userPtrValue.Kind())  // struct ptr
	userPtrValue3 := userValue.Addr()                   // Addr()非指针Value转为指针Value
	fmt.Println(userValue.Kind(), userPtrValue3.Kind()) // struct ptr
}
