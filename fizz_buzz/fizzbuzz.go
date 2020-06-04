// Matthew Kobilas 06/04/20
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var n, fizzbuzz int
	fmt.Print("How many fizzing and buzzing units do you need in your life? ")
	fmt.Scan(&s)
	n, _ = strconv.Atoi(s)
	for i := 0; fizzbuzz < n; i++ {
		if i%3 == 0 {
			fmt.Print("fizz ")
			fizzbuzz++
			if i%5 == 0 {
				fmt.Print("buzz")
			}
			fmt.Println()
			continue
		}
		if i%5 == 0 {
			fmt.Println("buzz")
			fizzbuzz++
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("TRADITION!!")
}
