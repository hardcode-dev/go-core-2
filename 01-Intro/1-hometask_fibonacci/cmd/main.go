package main

import (
	"fmt"
	"go-core-2/01-Intro/1-hometask_fibonacci/pkg/fibonacci"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	result, err := fibonacci.FindFibonacci(number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
