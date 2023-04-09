package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("jonson ReflectCallFunc")
}

func main() {
	user := User{
		Id:   1,
		Name: "jonson",
		Age:  25,
	}
	getType := reflect.TypeOf(user)
	getValue := reflect.ValueOf(user)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Println(field.Name)
		fmt.Println(field.Tag)
		fmt.Println(value)
	}
}
