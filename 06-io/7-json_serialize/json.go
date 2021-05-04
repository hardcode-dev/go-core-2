package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type T struct {
	I int
	S string
}

func main() {
	t1 := T{I: 10, S: "ABC"}
	// сериализация
	b, err := json.Marshal(t1) // установить точку останова
	if err != nil {
		log.Fatal(err)
	}
	var t2 T
	// десериализация
	err = json.Unmarshal(b, &t2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.DeepEqual(t1, t2))
}
