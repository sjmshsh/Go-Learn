package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Add(a, b int) int {
	return a + b
}

func main() {
	// 获取函数类型
	typeFunc := reflect.TypeOf(Add)
	fmt.Printf("is function type %t\n", typeFunc.Kind() == reflect.Func)
	argInNum := typeFunc.NumIn()   // 输入参数个数
	argOutNum := typeFunc.NumOut() // 输出参数的个数
	for i := 0; i < argInNum; i++ {
		argTyp := typeFunc.In(i)
		fmt.Printf("第%d个输入参数的类型%s\n", i, argTyp)
	}
	for i := 0; i < argOutNum; i++ {
		argTyp := typeFunc.Out(i)
		fmt.Printf("第%d个输出参数的类型%s\n", i, argTyp)
	}
}

// is function type true
// 第0个输入参数的类型int
// 第1个输入参数的类型int
// 第0个输出参数的类型int
