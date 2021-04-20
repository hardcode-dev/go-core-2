package fib

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		Description string
		Got         int
		Want        int
	}{
		{"Fibonacci`s number on index 0 is 0", 0, 0},
		{"Fibonacci`s number on index 1 is 0", 1, 1},
		{"Fibonacci`s number on index 2 is 0", 2, 1},
		{"Fibonacci`s number on index 3 is 0", 3, 2},
		{"Fibonacci`s number on index 4 is 0", 4, 3},
		{"Fibonacci`s number on index 5 is 0", 5, 5},
		{"Fibonacci`s number on index 6 is 0", 6, 8},
		{"Fibonacci`s number on index 7 is 0", 7, 13},
		{"Fibonacci`s number on index 8 is 0", 8, 21},
		{"Fibonacci`s number on index 9 is 0", 9, 34},
		{"Fibonacci`s number on index 10 is 0", 10, 55},
		{"Fibonacci`s number on index 11 is 0", 11, 89},
		{"Fibonacci`s number on index 12 is 0", 12, 144},
		{"Fibonacci`s number on index 13 is 0", 13, 233},
		{"Fibonacci`s number on index 14 is 0", 14, 377},
		{"Fibonacci`s number on index 15 is 0", 15, 610},
		{"Fibonacci`s number on index 16 is 0", 16, 987},
		{"Fibonacci`s number on index 17 is 0", 17, 1597},
		{"Fibonacci`s number on index 18 is 0", 18, 2584},
		{"Fibonacci`s number on index 19 is 0", 19, 4181},
		{"Fibonacci`s number on index 20 is 0", 20, 6765},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Count(test.Got)
			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
