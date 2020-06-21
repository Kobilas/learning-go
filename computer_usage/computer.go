package main

import (
	"fmt"
	"math/rand"
	"time"
)

func useComputer(tourist *int, touristIdx int, output chan string) {
	fmt.Println("Tourist", touristIdx, "is online.")
	var computingTime int = rand.Intn(106) + 15
	time.Sleep(time.Duration(computingTime/10) * time.Second)
	*tourist = computingTime
	output <- fmt.Sprintf("Tourist %d is done, having spent %d minutes online.", touristIdx, *tourist)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tourists := make([]int, 25)
	output := make(chan string)
	var sent int = 0
	var availableComputers = 8
	var first bool = true
	for sent < len(tourists)-1 {
		if availableComputers > 0 {
			for j := range tourists {
				if availableComputers > 0 {
					if tourists[j] == 0 {
						availableComputers--
						tourists[j] = -1
						go useComputer(&tourists[j], j+1, output)
					}
				}
			}
		}
		if first {
			for j, val := range tourists {
				if val == 0 {
					fmt.Println("Tourist", j+1, "is waiting for turn.")
				}
			}
			first = false
		}
		for _, val := range tourists {
			if val == -1 {
				fmt.Println(<-output)
				sent++
				if availableComputers < 8 {
					availableComputers++
				}
				break
			}
		}
	}
	fmt.Println("The place is empty, let's close up and go to the beach!")
}
