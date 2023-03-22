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

// Value.SetMapIndex()函数：往map里添加一个key-value对

// Value.MapIndex()函数： 根据Key取出对应的map

func main() {
	u1 := &User{
		Name: "j",
		Age:  18,
	}
	u2 := &User{
		Name: "a",
		Age:  23,
	}
	userMap := make(map[int]*User, 5)
	userMap[u1.Age] = u1
	//准备通过Value修改userMap，所以传userMap的地址
	mapValue := reflect.ValueOf(&userMap)
	//SetMapIndex 往map里添加一个key-value对
	mapValue.Elem().SetMapIndex(reflect.ValueOf(u2.Age), reflect.ValueOf(u2))
	//MapIndex 根据Key取出对应的map
	mapValue.Elem().MapIndex(reflect.ValueOf(u1.Age)).Elem().FieldByName("Name").SetString("x")
	for k, user := range userMap {
		fmt.Printf("key %d name %s\n", k, user.Name)
	}
}
