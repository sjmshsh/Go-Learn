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
	user := User{
		Name: "lxy",
		Age:  20,
	}
	// 必须传指针，因为BMI()定义的时候它是指针的方法
	valueUser := reflect.ValueOf(&user)
	//MethodByName()通过Name返回类的成员变量
	bmiMethod := valueUser.MethodByName("BMI")
	// 无参数的时候传入一个空的切片
	resultValue := bmiMethod.Call([]reflect.Value{})
	result := resultValue[0].Interface().(float32)
	fmt.Printf("%.2f\n", result)

	//Think()在定义的时候用的不是指针，valueUser可以用指针也可以不用指针
	thinkMethod := valueUser.MethodByName("Think")
	thinkMethod.Call([]reflect.Value{})

	valueUser2 := reflect.ValueOf(user)
	thinkMethod = valueUser2.MethodByName("Think")
	thinkMethod.Call([]reflect.Value{})
}
