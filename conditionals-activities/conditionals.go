// Activity 1
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var worthStr string
	var worth int
	fmt.Println("Enter amount of money in wallet: ")
	fmt.Scan(&worthStr)
	worth, _ = strconv.Atoi(worthStr)
	if worth >= 20 {
		fmt.Println("You're rich!")
	} else {
		fmt.Println("You're broke!")
	}
} */

// Activity 2
/* package main
import "fmt"
func main() {
	var dogs, cats string
	fmt.Print("Do you have any dogs (y/n)? ")
	fmt.Scan(&dogs)
	fmt.Print("Do you have any cats (y/n)? ")
	fmt.Scan(&cats)
	if dogs == "y" && cats == "y" {
		fmt.Println("You must really love pets!")
	} else {
		fmt.Println("Maybe you need more pets.")
	}
	//
	//	if dogs == "y" && cats == "y" {
	//		fmt.Println("You must really love pets!")
	//		return;
	//	}
	//	fmt.Println("Maybe you need more pets.")
	//
} */
// Activity 3
// Activity 4
/* package main
import "fmt"
func main() {
	var season string
	fmt.Print("What season is it? ")
	fmt.Scan(&season)
	if season == "Fall" || season == "Autumn" {
		fmt.Println("I bet the leaves are pretty there!")
	} else if season == "Winter" {
		fmt.Println("I hope you're ready for snow!")
	} else if season == "Spring" {
		fmt.Println("I can smell the flowers!")
	} else if season == "Summer" {
		fmt.Println("Make sure your AC is working!")
	} else {
		fmt.Println("I don't recognize that season.")
	}
} */

// Activity 5
/* package main

import "fmt"

func main() {
	var season string
	fmt.Print("What season is it? ")
	fmt.Scan(&season)
	switch season {
	case "Autumn":
		fallthrough
	case "Fall":
		fmt.Println("I bet the leaves are pretty there!")
	case "Winter":
		fmt.Println("I hope you're ready for snow!")
	case "Spring":
		fmt.Println("I can smell the flowers!")
	case "Summer":
		fmt.Println("Make sure your AC is working!")
	default:
		fmt.Println("I don't recognize that season.")
	}
} */
