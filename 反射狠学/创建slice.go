package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	addr string `json:"addr"`
}

func (u *User) BMI() float32 {
	fmt.Println("调用了BMI方法")
	return 5.5
}

func (u *User) Think() {
	fmt.Println("调用了Think方法")
}

func main() {
	var slice []User
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.MakeSlice(sliceType, 1, 3)
	sliceValue.Index(0).Set(reflect.ValueOf(User{
		Name: "lxy",
		Age:  18,
	}))
	users := sliceValue.Interface().([]User)
	fmt.Println(users[0].Name)
}
