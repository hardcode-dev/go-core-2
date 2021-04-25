package fibo

// Num возвращает число Фибоначчи
// по индексу.
func Num(n int) int {
	x1, x2 := 0, 1
	for i := 1; i <= n; i++ {
		x1, x2 = x2, x1+x2
	}
	return x2
}
