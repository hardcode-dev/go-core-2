package fibo

import (
	"fmt"
)

const max = 20

// Num calculates Nth Fibonacci n iteratively using O(1) memory
func Num(n int) (int, error) {

	// validate that the passed parameter is not above the max (upper limit)
	err := checkMax(n)
	if err != nil {
		return 0, err
	}

	// a few more checks
	if n == 1 {
		return n, nil
	}
	if n < 1 {
		return 1, nil
	}

	// now let's make iterations and get the answer
	var n2, n1 int = 1, 0

	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

func checkMax(n int) error {
	if n > max {
		return fmt.Errorf("the Fibonacci number to find is above the upper limit of %d", max)
	}
	return nil
}
