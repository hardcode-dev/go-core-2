package fib

func Count(n int) int {
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
