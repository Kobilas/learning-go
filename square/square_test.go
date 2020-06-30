package main

import (
	"fmt"
	"testing"
)

// test if the Square function produces the correct output for a particular
// scenario a = 2
// *testing.T is a type passed to Test functions to manage test state and
// support formatted test logs
func TestSquare(t *testing.T) {
	ans := Square(2)
	if ans != 4 {
		// use t to record the testing error
		t.Errorf("Square(2) = %d; Should be 4", ans)
	}
}

// use table-driven testing to test the function in various situations
func TestSquareTableDriven(t *testing.T) {
	var tests = []struct {
		a      int
		expect int
	}{
		// the first input is a = 0 and the expected output is 0
		{0, 0},
		// the second input is a = 1 and the expected output is 1 * 1 = 1
		{1, 1},
		// the third input is a = 2 and the expected output is 2 * 2 = 4
		{2, 4},
		{6, 36},
		{5, 25},
	}
	// iterate through each test table entry
	for _, tt := range tests {
		// display the value to be tested
		testname := fmt.Sprintf("%d", tt.a)
		// use t.Run method to execute the test for the current table entry
		t.Run(testname, func(t *testing.T) {
			// the second argument of Run is a function that simply calls the
			// Square function and compares it against the expected output
			ans := Square(tt.a)
			if ans != tt.expect {
				// use t (type testing.T) to record the error
				t.Errorf("got %d, want %d", ans, tt.expect)
			}
		})
	}
}
