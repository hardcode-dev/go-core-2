package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var i64 int64
	var s string
	fmt.Printf("i: %d, %d\ni64: %d, %d\ns: %s, %d\n", i, unsafe.Sizeof(i), i64, unsafe.Sizeof(i64), s, unsafe.Sizeof(s))
	var array [3]int
	var slice []int
	fmt.Printf("array: %v, size: %d, length: %d\nslice: %v, size: %d, length: %d\n",
		array, unsafe.Sizeof(array), len(array), slice, unsafe.Sizeof(slice), len(slice))
	//i = i64 // ошибка
	//s = i // ошибка
	fmt.Println("пустой слайс == nil?", slice == nil) // истина
	//fmt.Println(array == nil) // ошибка
	var myMap map[int]string
	var iface interface{}
	var f func(int) string
	_, _, _ = myMap, iface, f
	b := []byte{10, 11, 12, 13}
	s1 := string(b)
	b2 := []byte(s1)
	_ = b2
}
