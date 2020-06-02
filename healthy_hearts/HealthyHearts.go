package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Healthy Hearts")
	fmt.Print("What is your age? ")
	var age string
	fmt.Scan(&age)
	var ageNum int
	ageNum, _ = strconv.Atoi(age)
	fmt.Println("Your maximum heart rate should be", 120-ageNum)
	fmt.Println("Your target HR Zone is", math.Trunc(0.5*float64(120-ageNum)),
		"-", math.Trunc(0.85*float64(120-ageNum)))
}
