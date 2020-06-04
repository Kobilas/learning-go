// Matthew Kobilas 06/04/20
package main

import (
	"fmt"
	"strings"
)

func main() {
	var size, coffeeType, flavorYN, flavoring string
	var price float64
	fmt.Print("Do you want small, medium, or large? ")
	fmt.Scan(&size)
	size = strings.ToLower(size)
	fmt.Print("Do you want brewed, espresso, or iced? ")
	fmt.Scan(&coffeeType)
	coffeeType = strings.ToLower(coffeeType)
	fmt.Print("Do you want a flavored syrup? (Yes or No) ")
	fmt.Scan(&flavorYN)
	flavorYN = strings.ToLower(flavorYN)
	if flavorYN == "yes" {
		fmt.Print("Do you want hazelnut, vanilla, or caramel? ")
		fmt.Scan(&flavoring)
		flavoring = strings.ToLower(flavoring)
	}
SIZE_SWITCH:
	switch size {
	case "small":
		price += 2.0
	case "medium":
		price += 3.0
	case "large":
		price += 4.0
	default:
		fmt.Print("Do you want small, medium, or large? ")
		fmt.Scan(&size)
		size = strings.ToLower(size)
		goto SIZE_SWITCH
	}
TYPE_IF:
	if coffeeType == "espresso" {
		price += 0.5
	} else if coffeeType == "iced" {
		price += 1.0
	} else if coffeeType != "brewed" {
		fmt.Print("Do you want brewed, espresso, or iced? ")
		fmt.Scan(&coffeeType)
		coffeeType = strings.ToLower(coffeeType)
		goto TYPE_IF
	}
CHOICE_IF:
	if flavorYN == "yes" {
	FLAVOR_IF:
		if flavoring == "hazelnut" || flavoring == "vanilla" || flavoring == "caramel" {
			price += 0.5
		} else {
			fmt.Print("Do you want hazelnut, vanilla, or caramel? ")
			fmt.Scan(&flavoring)
			flavoring = strings.ToLower(flavoring)
			goto FLAVOR_IF
		}
	} else if flavorYN != "no" {
		fmt.Print("Do you want a flavored syrup? (Yes or No) ")
		fmt.Scan(&flavorYN)
		flavorYN = strings.ToLower(flavorYN)
		goto CHOICE_IF
	}
	fmt.Println("You asked for a", size, "cup of", coffeeType, "coffee with", flavoring, "syrup.")
	fmt.Printf("Your cup of coffee costs %.2f\n", price)
	fmt.Printf("The price with a tip is %.2f\n", price*1.15)
}
