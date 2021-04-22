package main

import (
	"fmt"
	"go-core-2/01-Intro/1-hometask_fibonacci/pkg/fibo"
)

func main() {
	fmt.Println("Please type in your number and press Enter")
	var n int
	fmt.Scanf("%d", &n)

	res, err := fibo.Num(n)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Printf("The calculated result is %d\n", res)
}
