package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var dollarsStr string
	var dollars, rollsTaken, maxMoneyHeld, rollsMaxMoneyHeld int
	ns := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(ns)
	fmt.Println("How many dollars do you have? ")
	fmt.Scan(&dollarsStr)
	dollars, _ = strconv.Atoi(dollarsStr)
	maxMoneyHeld = dollars
	for {
		num1 := rng.Intn(6) + 1
		num2 := rng.Intn(6) + 1
		if num1+num2 == 7 {
			dollars += 4
		} else {
			dollars--
		}
		rollsTaken++
		if dollars > maxMoneyHeld {
			maxMoneyHeld = dollars
			rollsMaxMoneyHeld = rollsTaken
		}
		if dollars <= 0 {
			break
		}
	}
	fmt.Println("You are broke after", rollsTaken, "rolls.")
	fmt.Println("You should have quit after", rollsMaxMoneyHeld,
		"rolls when you had", maxMoneyHeld, "dollars.")
}
