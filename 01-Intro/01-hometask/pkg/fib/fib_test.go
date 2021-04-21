package fib

import (
	"testing"
)

func TestCalculation(t *testing.T) {
	cases := []struct {
		Description string
		Input       int
		Want        int
	}{
		{"Fibonacci`s number on index 0 is 0", 0, 0},
		{"Fibonacci`s number on index 1 is 1", 1, 1},
		{"Fibonacci`s number on index 2 is 1", 2, 1},
		{"Fibonacci`s number on index 3 is 2", 3, 2},
		{"Fibonacci`s number on index 4 is 3", 4, 3},
		{"Fibonacci`s number on index 5 is 5", 5, 5},
		{"Fibonacci`s number on index 6 is 8", 6, 8},
		{"Fibonacci`s number on index 7 is 13", 7, 13},
		{"Fibonacci`s number on index 8 is 21", 8, 21},
		{"Fibonacci`s number on index 9 is 34", 9, 34},
		{"Fibonacci`s number on index 10 is 55", 10, 55},
		{"Fibonacci`s number on index 11 is 89", 11, 89},
		{"Fibonacci`s number on index 12 is 144", 12, 144},
		{"Fibonacci`s number on index 13 is 233", 13, 233},
		{"Fibonacci`s number on index 14 is 377", 14, 377},
		{"Fibonacci`s number on index 15 is 610", 15, 610},
		{"Fibonacci`s number on index 16 is 987", 16, 987},
		{"Fibonacci`s number on index 17 is 1597", 17, 1597},
		{"Fibonacci`s number on index 18 is 2584", 18, 2584},
		{"Fibonacci`s number on index 19 is 4181", 19, 4181},
		{"Fibonacci`s number on index 20 is 6765", 20, 6765},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := count(test.Input)
			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	cases := []struct {
		Description string
		Input       int
		Want        string
	}{
		{"Should respond with error on less than 0 inputs", -1, "input cannot be less than zero"},
		{"Should respond with error on more than 20 inputs", 23, "input cannot be more than 20"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			_, got := Count(test.Input)
			if got.Error() != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
