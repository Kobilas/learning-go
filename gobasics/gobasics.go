// Matthew Kobilas
// 06/01/20
/*
	Hello world program which prints hello world,
	followed by my name, hometown, and favorite food.
*/
package main

import (
	"fmt"
)

var firstNameTwo, streetAddress, birthYear = "Matthew", "25 Cedar Avenue", 1998

func main() {
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
}
