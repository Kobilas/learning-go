package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addOne(m int) int {
	return m + 1
}

func modulo(m, n int) int {
	return m % n
}

// DisplayUpper prints string entered in uppercase
func DisplayUpper(x string) {
	fmt.Println("Original text: ", x)
	fmt.Println("Revised text: ", strings.ToUpper(x))
}

func sumN(numbers ...int) {
	sum := 0
	for i, num := range numbers {
		fmt.Println("Current element:", num, "; Current index:", i)
		sum += num
	}
	fmt.Println("Sum of all vallues:", sum)
}

func recursiveFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	}
}

// doesn't work, but i know how closures work already zzz
func fibonacciSequencer(n int) func() string {
	ret := ""
	return func() string {
		var prev int = 0
		for i := 0; i <= n; i++ {
			prev = prev + i
			ret += strconv.Itoa(prev) + " "
		}
		return ret
	}
}

func main() {
	// regular functions with returns
	var mS, nS string
	var m, n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&mS)
	m, _ = strconv.Atoi(mS)
	fmt.Print("Enter another number: ")
	fmt.Scan(&nS)
	n, _ = strconv.Atoi(nS)
	fmt.Println("First number you entered plus one:", addOne(m))
	fmt.Println("First number you entered modulo second number:", modulo(m, n))

	DisplayUpper("elizabeth")

	supermutation := func(a, b, c int) (x, y, z int) {
		x = a + b
		y = a + c
		z = b + c
		return
	}
	fmt.Println(supermutation(4, 6, 5))

	// variadic function with variable number of parameters
	sumN(4, 6, 5)
	sumN(4, 6, 5, 6, 7, 8)

	// recursive function
	fmt.Println("Fibonacci up to 7:", recursiveFibonacci(7))

	// closure that returns fibonacci sequence
	fibSeq := fibonacciSequencer(17)
	fmt.Println("Fibonacci up to 17")
	fmt.Println(fibSeq())
}
