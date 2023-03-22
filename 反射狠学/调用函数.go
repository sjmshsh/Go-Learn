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

func Add(a, b int) int {
	return a + b
}

func main() {
	// 函数也是一种数据类型
	valueFunc := reflect.ValueOf(Add)
	typeFunc := reflect.TypeOf(Add)
	//函数输入参数的个数
	argNum := typeFunc.NumIn()
	// 准备函数的输入参数
	args := make([]reflect.Value, argNum)
	for i := 0; i < argNum; i++ {
		if typeFunc.In(i).Kind() == reflect.Int {
			args[i] = reflect.ValueOf(3)
		}
	}
	//返回[]reflect.Value，因为go语言的函数返回可能是一个列表
	sumValue := valueFunc.Call(args)
	if typeFunc.Out(0).Kind() == reflect.Int {
		sum := sumValue[0].Interface().(int) // 从Value转为原始数据类型
		fmt.Println(sum)
	}
}
