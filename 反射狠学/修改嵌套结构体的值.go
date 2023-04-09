package main

import "reflect"

type Children struct {
	Age int
}

type Nested struct {
	X     int
	Child Children
}

func main() {
	vs := reflect.ValueOf(&Nested{}).Elem()
	vz := vs.Field(1)
	vz.Set(reflect.ValueOf(Children{Age: 19}))
}
