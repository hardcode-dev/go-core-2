package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	x := 10
	f := func() {
		log.Println("X доступна из замыкания:", x)
	}
	f()
	printInts(6, 2, 3, 4)
	res, err := division(10, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("division(10, 5) =", res)
}

func printInts(ints ...int) {
	for i, v := range ints {
		fmt.Print(i, ":", v, " ")
	}
	fmt.Println()
}

func division(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("ошибка: деление на ноль")
	}
	return x / y, nil
}

// Output: 	2020/09/27 14:26:40 10 делить на 3 будет 3.333333
//			0 1 2 3
//			2020/09/27 14:26:40 X доступна из замыкания: 10
