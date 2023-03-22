package main

import (
	"fmt"
	"sync"
)

func walk1(key, value interface{}) bool {
	fmt.Println("Key =", key, "Value =", value)
	return true
}

func walk2(key, value interface{}) bool {
	fmt.Println("Key =", key, "Value =", value)
	return false
}

func main() {
	var lxy sync.Map
	lxy.Store("Db", "Redis")
	lxy.Store("lxy", "zyq")
	lxy.Range(walk1)
	lxy.Range(walk2)
}
