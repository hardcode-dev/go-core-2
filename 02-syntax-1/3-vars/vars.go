package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s1 string
	var s2 string = "ABC"
	s3 := "def"
	fmt.Printf("s1: %s, %s\ns2: %s\ns3: %s\n", s1, reflect.TypeOf(s1), s2, s3)
	i1 := 10
	var i2 int = 20
	var i3 int64 = 30
	fmt.Printf("i1: %d, %s\ni2: %d, %s\ni3: %d, %s\n",
		i1, reflect.TypeOf(i1), i2, reflect.TypeOf(i2), i3, reflect.TypeOf(i3))
}
