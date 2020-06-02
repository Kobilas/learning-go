package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var name, difficulty, strGuess string
	var answer, guess, numGuesses int
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Select difficulty (easy/medium/hard): ")
	fmt.Scan(&difficulty)
	ns := rand.NewSource(time.Now().UnixNano())
	rndGen := rand.New(ns)
	switch difficulty {
	case "easy":
		answer = rndGen.Intn(5)
	case "medium":
		answer = rndGen.Intn(20)
	case "hard":
		answer = rndGen.Intn(50)
	}
	for guess != answer {
		fmt.Print("Enter a guess: ")
		fmt.Scan(&strGuess)
		guess, _ = strconv.Atoi(strGuess)
		switch difficulty {
		case "easy":
			if guess < 1 || guess > 5 {
				fmt.Println("Error: " + name + ", bounds in easy mode are 1 and 5")
				continue
			}
		case "medium":
			if guess < 1 || guess > 20 {
				fmt.Println("Error: " + name + ", bounds in medium mode are 1 and 20")
				continue
			}
		case "hard":
			if guess < 1 || guess > 50 {
				fmt.Println("Error: " + name + ", bounds in hard mode are 1 and 50")
				continue
			}
		}
		numGuesses++
		if guess == answer {
			if numGuesses == 1 {
				fmt.Println("WOW FIRST TRY! Congratulations " + name + ", you win!")
			} else {
				fmt.Println("Congrats " + name + ", you win!")
			}
		} else if guess > answer {
			fmt.Println("Too high " + name)
		} else {
			fmt.Println("Too low " + name)
		}
	}
}
