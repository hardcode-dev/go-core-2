package main

import (
	"fmt"
	"go-core-2/homeworks/01-fibo/pkg/fibo"
)

func main() {
	nums := []int{1, 5, 8, 22}
	for _, n := range nums {
		if n > 20 {
			fmt.Printf("максимум - 20, переданный номер %d\n", n)
			continue
		}
		fmt.Printf("%d: %d\n", n, fibo.Num(n))
	}
}
