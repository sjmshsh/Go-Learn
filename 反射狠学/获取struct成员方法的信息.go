package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) GET2() {

}

func (u *User) GET1() {

}

func (u *User) GET() {

}

func main() {
	typeUser := reflect.TypeOf(User{})
	methodNum := typeUser.NumMethod() // 成员方法的个数，接收者为指针的方法不包含在内
	for i := 0; i < methodNum; i++ {
		method := typeUser.Method(i)
		fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}
	fmt.Println()

	typeUser2 := reflect.TypeOf(&User{})
	// 成员方法的个数。接收者为指针或值的方法【都】包含在内，也就是说值实现的方法指针也实现了（反之不成立）
	methodNum1 := typeUser2.NumMethod()
	for i := 0; i < methodNum1; i++ {
		method := typeUser2.Method(i)
		fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}
}

// method name:GET2 ,type:func(main.User), exported:true

// method name:GET ,type:func(*main.User), exported:true
// method name:GET1 ,type:func(*main.User), exported:true
// method name:GET2 ,type:func(*main.User), exported:true
