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

func main() {
	users := make([]*User, 1, 5) // len=1, cap=5
	users[0] = &User{
		Name: "lxy",
		Age:  18,
	}
	// 准备通过Value修改users，所以传递users地址
	sliceValue := reflect.ValueOf(&users)
	if sliceValue.Elem().Len() > 0 {
		// 获得slice的长度
		sliceValue.Elem().Index(0).Elem().FieldByName("Name").
			SetString("lll")
		fmt.Println(users[0].Name)
	}
}
