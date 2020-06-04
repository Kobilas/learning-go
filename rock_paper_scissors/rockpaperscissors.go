package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var ansStr, rndsStr, rpsStr string
	var ans, rnds, uWin, uLoss, uTie, rps int
	winMatrix := [][]int{
		[]int{2, 0, 1},
		[]int{1, 2, 0},
		[]int{0, 1, 2},
	}
GAME_START:
	fmt.Print("Enter the number of rounds to play: ")
	fmt.Scan(&rndsStr)
	rnds, _ = strconv.Atoi(rndsStr)
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)
	for i := 0; i < rnds; i++ {
		rps = gen.Intn(3)
		switch rps {
		case 0:
			rpsStr = "rock"
		case 1:
			rpsStr = "paper"
		case 2:
			rpsStr = "scissors"
		}
		fmt.Print("Rock, paper, or scissors? ")
		fmt.Scan(&ansStr)
		switch ansStr {
		case "rock":
			ans = 0
		case "paper":
			ans = 1
		case "scissors":
			ans = 2
		}
		fmt.Println(ansStr, " - ", rpsStr)
		switch winMatrix[ans][rps] {
		case 0:
			uLoss++
		case 1:
			uWin++
		case 2:
			uTie++
		}
	}
	if uWin > uLoss {
		fmt.Println("You win!")
	} else if uLoss > uWin {
		fmt.Println("You lose!")
	} else {
		fmt.Println("Its a tie!")
	}
	fmt.Print("Play again? ")
	fmt.Scan(&ansStr)
	if ansStr == "yes" {
		goto GAME_START
	} else {
		fmt.Println("Thanks for playing!")
	}
}
