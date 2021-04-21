package fib

import "errors"

func Count(n int) (int, error) {
	switch {
	case n < 0:
		return 0, errors.New("input cannot be less than zero")
	case n > 20:
		return 0, errors.New("input cannot be more than 20")
	default:
		return count(n), nil
	}

}

func count(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		switch i {
		case 0:
			f = 0
		case 1, 2:
			f = 1
		default:
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}
