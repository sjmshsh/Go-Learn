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
		Age:  7,
		Name: "杰克逊",
	})
	userValue := userPtrValue.Elem() // Elem()指针Value转为非指针Value
	fmt.Printf("origin value iValue is %d %d\n", iValue.Interface().(int), iValue.Int())
	fmt.Printf("origin value sValue is %s %s\n", sValue.Interface().(string), sValue.String())
	user := userValue.Interface().(User)
	fmt.Printf("%d %s", user.Age, user.Name)
	user2 := userPtrValue.Interface().(*User)
	fmt.Printf("%d %s", user2.Age, user2.Name)
}

// origin value iValue is 1 1
// origin value sValue is hello hello
// 7 杰克逊7 杰克逊       
