package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bankAccount float64 = 0.00
	if !(bankAccount > 0) {
		fmt.Println("Transaction could not be completed")
	}
	var username string = "chris"
	var password string = "dsxscg34"
	if username == "chris" && password == "dsxscg34" {
		fmt.Println("This person has the right credentials")
	}
	var electricity bool = false
	var internet bool = false
	if !(electricity && internet) {
		fmt.Println("You must have electricity and internet to run this code")
	} else {
		fmt.Println("Running code beep boop")
	}

	var color string = "Red"
	if color == "Blue" {
		fmt.Println("Blue like the sky")
	} else if color == "Red" {
		fmt.Println("Red like the sun")
	} else if color == "Green" {
		fmt.Println("Green like the trees")
	} else {
		fmt.Println("Please choose a valid color")
	}

	var foodChoice string = "Tacos"
	if foodChoice == "Tacos" {
		fmt.Println("$6.99")
	} else if foodChoice == "Burgers" {
		fmt.Println("$10.99")
	} else if foodChoice == "Spaghetti" {
		fmt.Println("$4.99")
	} else if foodChoice == "Salad" {
		fmt.Println("$1.00")
	} else {
		fmt.Println("Aren't you hungry?")
	}

	var colorTwo string = "Blue"
	switch colorTwo {
	case "Blue":
		fmt.Println("Blue like the sky")
	case "Red":
		fmt.Println("Red like the sun")
	case "Green":
		fmt.Println("Green like the trees")
	default:
		fmt.Println("Please choose a valid color")
	}

	var foodChoiceTwo string = "Spaghetti"
	switch foodChoiceTwo {
	case "Tacos":
		fmt.Println("$6.99")
	case "Burgers":
		fmt.Println("$10.99")
	case "Spaghetti":
		fmt.Println("$4.99")
	case "Salad":
		fmt.Println("$1.00")
	default:
		fmt.Println("Aren't you hungry")
	}

	var power2 int64 = 1
	var a int64 = 1
	for a < 10 {
		fmt.Println("2 to the power of " + strconv.FormatInt(a, 10) + " is equal to " + strconv.FormatInt(power2, 10))
		power2 += power2
		a++
	}

	var message string = "birdbirdbird"
	fmt.Println(message)
	for i := 1; i < len(message); i += 2 {
		fmt.Print(string(message[i]))
	}
	fmt.Println("")
	for i, c := range message {
		if i%2 == 0 {
			fmt.Print(i)
		} else {
			fmt.Print(string(c))
		}
	}
	fmt.Println("")

	b := 0
	for ; ; b += 2 {
		if b > 10 {
			break
		} else {
			fmt.Println(strconv.Itoa(b) + " is an even number")
		}
	}

}
