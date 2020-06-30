package main

import "fmt"

// FormatAmount formats a float64 into a string that says USD and the number to 2 decimal points
func FormatAmount(a float64) string {
	if a < 0 {
		return "Impossible operation"
	}
	// use %.2f for precision 2 which is enough to represent dollar amounts for now
	return fmt.Sprintf("USD %.2f", a)
}

// SubtractFormatAmount subtracts the two parameters then formats the result as a dollar amount
// as in function above
func SubtractFormatAmount(a, b float64) string {
	// if a or b is negative
	if a < 0 || b < 0 {
		return "Impossible operation"
	}
	if a >= b {
		return fmt.Sprintf("USD %.2f", a-b)
	}
	// if 0 < a < b
	return "Impossible operation"
}

// AddFormatAmount adds two parameters together then formats result as dollar amount
func AddFormatAmount(a, b float64) string {
	if a < 0 || b < 0 {
		return "Impossible operation"
	}
	return fmt.Sprintf("USD %.2f", a+b)
}

func main() {
	fmt.Println("main function")
	fmt.Println(FormatAmount(0.00))
	fmt.Println(FormatAmount(-10.00))
	fmt.Println(SubtractFormatAmount(0.03, 0.42))
	fmt.Println(SubtractFormatAmount(-0.03, -0.42))
}
