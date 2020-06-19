package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getReady(name string, person chan string) {
	var timeToReady int = rand.Intn(31) + 60
	//var timeToReady int = 5
	//var name string = <-person
	fmt.Println(name, "started getting ready")
	//time.Sleep(6 * time.Second)
	time.Sleep(time.Duration(timeToReady) * time.Second)
	person <- fmt.Sprintf("%s spent %d seconds getting ready", name, timeToReady)
}

func armAlarm(alarm chan string) {
	var timeToArm int = 60
	fmt.Println("Arming alarm")
	time.Sleep(time.Duration(int(timeToArm/2)) * time.Second)
	fmt.Println("Alarm is counting down")
	time.Sleep(time.Duration(int(timeToArm/2)) * time.Second)
	alarm <- fmt.Sprintf("Alarm is armed")
}

func putOnShoe(name string, person chan string) {
	var timeToShoe int = rand.Intn(11) + 35
	//var timeToShoe int = 5
	fmt.Println(name, "started putting on shoes")
	time.Sleep(time.Duration(timeToShoe) * time.Second)
	person <- fmt.Sprintf("%s spent %d seconds putting on shoes", name, timeToShoe)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//var bobChan, aliceChan chan string
	var bobChan = make(chan string)
	var aliceChan = make(chan string)
	var alarmChan = make(chan string)
	var bob, alice, alarm string
	/* bobChan <- "Bob"
	aliceChan <- "Alice" */
	fmt.Println("Let's go for a walk!")
	go getReady("Bob", bobChan)
	go getReady("Alice", aliceChan)
	bob = <-bobChan
	alice = <-aliceChan
	fmt.Println(bob)
	fmt.Println(alice)
	go armAlarm(alarmChan)
	go putOnShoe("Bob", bobChan)
	go putOnShoe("Alice", aliceChan)
	bob = <-bobChan
	alice = <-aliceChan
	fmt.Println(bob)
	fmt.Println(alice)
	fmt.Println("Exiting and locking the door")
	alarm = <-alarmChan
	fmt.Println(alarm)
}
