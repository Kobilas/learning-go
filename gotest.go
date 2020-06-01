// Matthew Kobilas
// 06/01/20
/*
	Hello world program which prints hello world,
	followed by my name, hometown, and favorite food.
*/
package main

import (
	"fmt"
	"strconv"
)

var firstNameTwo, streetAddress, birthYear = "Matthew", "25 Cedar Avenue", 1998

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("My name is Matthew Kobilas.")
	fmt.Println("I live in Middletown, New Jersey.")
	fmt.Println("My favorite food is kielbasa.")

	var output string = "Hello world!"
	var yr2020 string = "The year is 2020."
	fmt.Println(output, " ", yr2020)

	/*
		customerName
		accountType
		accountNumber
		currentBalance
		transactionDescription -> txDesc
		transactionAmount -> txAmt
	*/
	fmt.Print("Enter your first name: ")
	var firstName string
	fmt.Scanln(&firstName)
	fmt.Print("Enter your last name: ")
	var lastName string
	fmt.Scanln(&lastName)
	fmt.Print("Enter your house number: ")
	var houseNumber string
	fmt.Scanln(&houseNumber)
	fmt.Print("Enter your street name: ")
	var streetName string
	fmt.Scanln(&streetName)
	fmt.Print("Enter your city: ")
	var city string
	fmt.Scanln(&city)
	fmt.Print("Enter your state abbreviated (2 letters): ")
	var stateAbbr string
	fmt.Scanln(&stateAbbr)
	fmt.Print("Enter your zip code: ")
	var zipCode string
	fmt.Scanln(&zipCode)
	fmt.Println(firstName + " " + lastName)
	fmt.Println(houseNumber + " " + streetName)
	fmt.Println(city + " " + stateAbbr + " " + zipCode)

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
