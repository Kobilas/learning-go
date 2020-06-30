package main

import "fmt"

func Square(a int) int {
	return a * a
}

func main() {
	fmt.Println("Main function")
	fmt.Println(Square(2))
}
