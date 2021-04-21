package fibonacci

import "fmt"

const upperLimit = 20

// FindFibonacci calculates Nth Fibonacci number iteratively using O(1) memory
func FindFibonacci(number int) (int, error) {

	// validate that the passed parameter is not above the upperLimit
	err := checkUpperLimit(number)
	if err != nil {
		return 0, err
	}

	// a few more checks
	if number == 1 {
		return number, nil
	}
	if number < 1 {
		return 1, nil
	}

	// now let's make iterations and get the answer
	var n2, n1 int = 1, 0

	for i := 2; i <= number; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

func checkUpperLimit(number int) error {
	if number > upperLimit {
		return fmt.Errorf("the Fibonacci number to find is above the upper limit of %d", upperLimit)
	}
	return nil
}
