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
	var i interface{} // 接口没有指向任何具体的值
	v := reflect.ValueOf(i)
	fmt.Printf("v持有值 %t, type of v is Invalid %t\n", v.IsValid(), v.Kind() == reflect.Invalid)

	var user *User = nil
	v = reflect.ValueOf(user) // Value指向一个nil
	if v.IsValid() {
		// 调用IsNil()前先确保IsValid()，否则会panic
		fmt.Printf("v持有的值是nil %t\n", v.IsNil())
	}

	var u User // 只声明，里面的值都是0
	v = reflect.ValueOf(u)
	if v.IsValid() {
		//调用IsZero()前先确保IsValid()，否则会panic
		fmt.Println(v.IsZero())
	}
}

// v持有值 false, type of v is Invalid true
// v持有的值是nil true
// true               

