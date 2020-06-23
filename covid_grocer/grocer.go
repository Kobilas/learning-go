package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait(i int, wg *sync.WaitGroup, out chan string) {
	fmt.Println("Customer", i, "has entered the queue")
	shop(i, wg, out)
}

func shop(i int, wg *sync.WaitGroup, out chan string) {
	var timeToShop int = (rand.Intn(6) + 5) * 100
	fmt.Println("Customer", i, "has entered the grocery store")
	time.Sleep(time.Duration(timeToShop) * time.Millisecond)
	out <- fmt.Sprintf("Customer %d has left the grocery store", i)
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var numCustomers int = rand.Intn(51) + 50
	var wg sync.WaitGroup
	var timeToCustomer, customersInStore, arrived int
	output := make(chan string, 5)
	fmt.Println("You think", numCustomers, "people are going to come today")
	for arrived < numCustomers {
		if customersInStore < 5 {
			timeToCustomer = (rand.Intn(6) + 5) * 100
			time.Sleep(time.Duration(timeToCustomer) * time.Millisecond)
			wg.Add(1)
			go wait(arrived, &wg, output)
			customersInStore++
			continue
		}
		fmt.Println(<-output)
		customersInStore--
		arrived++
	}
	wg.Wait()
}
