package main

import "fmt"

const num = 20
/*
Функцию Фибоначи подсмотрел в интернете, помню что сумма двух предыдущих, но алгоритм увы...
*/
func main() {
	for f,s := 0,1; s< num; f,s = s+f,f {
		fmt.Println(f)
	}
}