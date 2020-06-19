package main

import (
	"fmt"
	"math/rand"
	"time"
)

type philosopher struct {
	timeThinking int
	timeEating   int
	index        int
}

func genPhilosopher(index int) philosopher {
	return philosopher{0, 0, index}
}

func think(philo *philosopher) {
	var timeToThink int = rand.Intn(1) + 1
	fmt.Println("Philosopher", philo.index, "is thinking")
	time.Sleep(time.Duration(timeToThink) * time.Second)
	philo.timeThinking += timeToThink
}

func eat(philo *philosopher, forkOne, forkTwo *int) {
	var forkChanOne = make(chan string)
	var forkChanTwo = make(chan string)
	var forkOneStr, forkTwoStr string
	var timeToEat int = rand.Intn(1) + 1
	fmt.Println("Philosopher is going to eat")
	if *forkOne == -1 {
		fmt.Println("Philosopher tries to pick up fork one")
		grabFork(philo, forkOne, forkChanOne)
		forkOneStr = <-forkChanOne
		fmt.Println(forkOneStr)
	}
	if *forkTwo == -1 {
		fmt.Println("Philosopher tries to pick up fork two")
		grabFork(philo, forkTwo, forkChanTwo)
		forkTwoStr = <-forkChanTwo
		fmt.Println(forkTwoStr)
	}
	fmt.Println("Philosopher", philo.index, "is eating with forks", *forkOne, "and", *forkTwo)
	time.Sleep(time.Duration(timeToEat) * time.Second)
	philo.timeEating += timeToEat
	*forkOne = -1
	*forkTwo = -1
}

func grabFork(philo *philosopher, forkPtr *int, fork chan string) {
	fmt.Println("Grab fork")
	fork <- fmt.Sprintf("%d fork taken", forkPtr)
	*forkPtr = philo.index
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var philo = [5]philosopher{
		genPhilosopher(0),
		genPhilosopher(1),
		genPhilosopher(2),
		genPhilosopher(3),
		genPhilosopher(4),
	}
	var fork = [5]int{-1, -1, -1, -1, -1}
	for i := 0; i != 50000; i++ {
		go eat(&philo[(i+1)%5], &fork[i%5], &fork[(i+1)%5])
		go think(&philo[(i+1)%5])

		time.Sleep(50 * time.Millisecond)
	}
}
