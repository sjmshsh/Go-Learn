package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type People interface {
	GetPos()
}

func main() {
	// 通过reflect.TypeOf((*<interface>)(nil)).Elem()获得接口类型。因为People是个接口不能创建实例，所以把nil强制转为*common.People类型
	typeOfPeople := reflect.TypeOf((*People)(nil)).Elem()
	fmt.Printf("typeOfPeople kind is interface %t\n", typeOfPeople.Kind() == reflect.Interface)
	t1 := reflect.TypeOf(User{})
	// t2 := reflect.TypeOf(&User{})
	// 如果值类型实现了接口，则指针类型也实现了接口，反之不成立
	fmt.Printf("t1 implements People interface %t\n", t1.Implements(typeOfPeople))
}

// typeOfPeople kind is interface true
// t1 implements People interface false
