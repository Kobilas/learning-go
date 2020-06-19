package blah

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
	var timeToEat int = rand.Intn(1) + 1
	*forkOne = philo.index
	*forkTwo = philo.index + 1
	fmt.Println("Philosopher", philo.index, "is eating with forks", *forkOne, "and", *forkTwo)
	time.Sleep(time.Duration(timeToEat) * time.Second)
	philo.timeEating += timeToEat
	*forkOne = -1
	*forkTwo = -1
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
		if fork[i%5] == -1 && fork[(i+1)%5] == -1 {
			go eat(&philo[(i+1)%5], &fork[i%5], &fork[(i+1)%5])
		} else {
			go think(&philo[(i+1)%5])
		}
		time.Sleep(50 * time.Millisecond)
	}
}
