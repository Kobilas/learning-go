package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Tourist struct {
	id         int
	onlineTime int
}

var tourists = make(chan Tourist, 25)

func useComputer(wg *sync.WaitGroup) {
	//go through the buffered channel, and do this for each user in it
	for tourist := range tourists {
		t1 := Tourist(tourist)
		fmt.Println("Tourist", t1.id, "is using a computer")
		time.Sleep(time.Millisecond * time.Duration(t1.onlineTime*100))
		fmt.Println("After", t1.onlineTime, "minutes user", t1.id, "is done with his computer")
	}
	wg.Done()
}

//creates workers, and runs the go routines using those workers
func createWorkerPool(computers int) {
	var wg sync.WaitGroup
	for i := 0; i < computers; i++ {
		wg.Add(1)
		go useComputer(&wg)
		fmt.Println("Loop")
	}
	wg.Wait()
}

// creates tourists, and adds them to the tourists channel
func CreateTourists(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		ot := rand.Intn(106) + 15
		tourist := Tourist{i + 1, ot}
		tourists <- tourist
		fmt.Println("Tourist number", tourist.id, "is in line.")
		time.Sleep(time.Millisecond * 100)
	}
	close(tourists)
}

func main() {
	//create our buffered channel of 25 tourists
	noOfTourists := 25
	go CreateTourists(noOfTourists)
	//create our workers, and run them
	computers := 8
	createWorkerPool(computers)
	//wait after the last user has ended, and exit the program
	time.Sleep(time.Second * 3)
	fmt.Println("Everyones done. Let's go to the beach")
}
