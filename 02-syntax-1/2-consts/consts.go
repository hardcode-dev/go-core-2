package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	var i32 int32 = 20
	var i64 int64 = 30

	// константа без указания типа
	const intConst = 100
	fmt.Printf("intConst: %d, тип intConst: %v\n", intConst, reflect.TypeOf(intConst))
	// константа с указанием типа
	const int64Const int64 = 100
	fmt.Printf("int64Const: %d, тип int64Const: %v\n", int64Const, reflect.TypeOf(int64Const))

	// приведени типов константы при присвоении переменным
	i = intConst
	i32 = intConst
	i64 = intConst

	//i = i32
	i = int(i32)
	//i = i64
	i = int(i64)
	i64 = int64(i)
}
