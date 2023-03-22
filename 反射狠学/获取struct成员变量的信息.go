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
	typeUser := reflect.TypeOf(User{}) // 需要用struct的Type，不能用指针的Type
	fieldNum := typeUser.NumField()    // 成员变量的个数
	for i := 0; i < fieldNum; i++ {
		field := typeUser.Field(i)
		fmt.Printf("%d %s offset %d anonymous %t type %s exported %t json tag %s\n", i,
			field.Name,            // 变量名称
			field.Offset,          // 相对于结构体首地址的内存偏移量，string类型会占据16个字节
			field.Anonymous,       // 是否是匿名变量
			field.Type,            // 数据类型，reflect.Type类型
			field.IsExported(),    // 包外是否可见（是否是以大写字母开头）
			field.Tag.Get("json"), // 获取成员变量后面``里面定义的tag
		)
	}
	fmt.Println()
	// 可以通过FieldByName获取Field
	if nameField, ok := typeUser.FieldByName("Name"); ok {
		fmt.Printf("Name is exported %t\n", nameField.IsExported())
	}
	// 也可以根据FieldByIndex获取Field
	thirdField := typeUser.FieldByIndex([]int{1})
	fmt.Printf("third field name %s\n", thirdField.Name)
}
