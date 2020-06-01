package main

import (
	"fmt"
	"strconv"
)

var firstNameTwo, streetAddress, birthYear = "Matthew", "25 Cedar Avenue", 1998

func main() {
	var firstNameTwo, streetAddress, birthYear = "Matthew", "25 Cedar Avenue", 1998
	fmt.Print(&firstNameTwo)
	fmt.Println(": " + firstNameTwo)
	fmt.Print(&streetAddress)
	fmt.Println(": " + streetAddress)
	fmt.Print(&birthYear)
	fmt.Printf(": %d\n", birthYear)

	var firstNum, secondNum string
	var firstInt, secondInt int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&firstNum)
	fmt.Print("Enter the second integer: ")
	fmt.Scan(&secondNum)
	firstInt, _ = strconv.Atoi(firstNum)
	secondInt, _ = strconv.Atoi(secondNum)
	fmt.Println(firstNum + secondNum)
	fmt.Println(firstInt + secondInt)

	var thirdNum, fourthNum string
	var firstFlt, secondFlt float64
	fmt.Print("Enter first float: ")
	fmt.Scan(&thirdNum)
	fmt.Print("Enter second float: ")
	fmt.Scan(&fourthNum)
	firstFlt, _ = strconv.ParseFloat(thirdNum, 64)
	secondFlt, _ = strconv.ParseFloat(fourthNum, 64)
	fmt.Println(thirdNum + fourthNum)
	fmt.Println(firstFlt + secondFlt)
}
