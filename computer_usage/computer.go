package main

import (
	"fmt"
	"math/rand"
	"time"
)

func useComputer(tourist *int, touristIdx int, output chan string) {
	fmt.Println("Tourist", touristIdx, "is online.")
	var computingTime int = rand.Intn(10) + 1
	time.Sleep(time.Duration(computingTime) * time.Second)
	*tourist = computingTime
	output <- fmt.Sprintf("Tourist %d is done, having spent %d minutes online.", touristIdx, *tourist)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tourists := make([]int, 25)
	output := make(chan string)
	var sent int = 0
	var availableComputers = 8
	var first = true
	for i := 0; sent < 24; i = (i + 1) % 25 {
		/* if availableComputers > 0 {
			if tourists[i] == 0 {
				availableComputers--
				tourists[i] = -1
				go useComputer(&tourists[i], i+1, output)
				continue
			}
		} */
		select {
		case x, ok := <-output:
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("Channel closed")
			}
		default:
			//fmt.Println("No value ready")
			if availableComputers > 0 {
				if tourists[i] == 0 {
					availableComputers--
					tourists[i] = -1
					go useComputer(&tourists[i], i+1, output)
					continue
				}
			}
		}
		if first {
			for i, val := range tourists {
				if val == 0 {
					fmt.Println("Tourist", i+1, "is waiting for turn.")
				} /* else if val == -1 {
					fmt.Println(<-output, sent)
					sent++
					if availableComputers < 8 {
						availableComputers++
					}
				} */
			}
			first = false
		}
		if sent == len(tourists) {
			break
		}
	}
	fmt.Println("The place is empty, let's close up and go to the beach!")
}
