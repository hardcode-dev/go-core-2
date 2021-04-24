package main

import "fmt"

func main() {
	var array = [...]int{1, 2, 3, 4, 5} // массив
	var slice []int = []int{1}          // слайс
	_ = slice[len(slice)-1]
	fmt.Println("array", array, len(array), cap(array))
	fmt.Println("slice", slice, len(slice), cap(slice))
	c := cap(slice)
	for i := 0; i < 1_000_000; i++ {
		slice = append(slice, i)
		if cap(slice) != c {
			fmt.Println("Длина slice:", len(slice),
				"Новая ёмкость slice:", cap(slice),
				"Коэффициент:", float64(cap(slice))/float64(c))
			c = cap(slice)
		}
	}
	slice = make([]int, 10, 20)
	fmt.Println("\nslice после инициализации:", slice, len(slice), cap(slice))

	// ассоциативный массив
	var m map[int]string
	//m[3] = "ABC" // panic: assignment to entry in nil map
	m = make(map[int]string)
	m = map[int]string{}
	m[3] = "ABC"
	m[6] = "DEF"
	for k, v := range m {
		fmt.Printf("ключ %d, значение %s\n", k, v)
	}
}
