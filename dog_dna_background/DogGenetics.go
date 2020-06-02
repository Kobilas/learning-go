package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ns := rand.NewSource(time.Now().UnixNano())
	rndGen := rand.New(ns)
	// does not actually capture entire scope of random values
	huskyPct := rndGen.Intn(20)
	maltesePct := rndGen.Intn(20)
	poodlePct := rndGen.Intn(20)
	chihuahuaPct := rndGen.Intn(20)
	dachshundPct := 100 - huskyPct - maltesePct - poodlePct - chihuahuaPct
	fmt.Println("What is your dog's name? ")
	var dogName string
	fmt.Scan(&dogName)
	fmt.Println("Well then, I have this highly reliable report on", dogName+"'s",
		"prestigious background right here.")
	fmt.Println(dogName, "is:")
	fmt.Println(huskyPct, "% Husky")
	fmt.Println(maltesePct, "% Maltese")
	fmt.Println(poodlePct, "% Poodle")
	fmt.Println(chihuahuaPct, "% Chihuahua")
	fmt.Println(dachshundPct, "% Dachshund")
}
