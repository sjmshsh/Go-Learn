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
	var i int = 10
	var s string = "hello"
	user := User{
		Name: "lxy",
		Age:  18,
	}
	// 由于go语言所有函数传的都是值，所以要想修改原来的值就需要传指针
	valueI := reflect.ValueOf(&i)
	valueS := reflect.ValueOf(&s)
	valueUser := reflect.ValueOf(&user)
	valueI.Elem().SetInt(8) // 由于valueI对应的原始对象是指针，通过Elem()返回指针指向的对象
	valueS.Elem().SetString("golang")
	// FieldByName()通过Name返回类的成员变量
	valueUser.Elem().FieldByName("Name").SetString("ZYQ")
	fmt.Println(user.Name)
	fmt.Println(i)
	fmt.Println(s)
}

// ZYQ
// 8     
// golang
