package main

import "fmt"

type burger struct {
	bun     bool
	price   float64
	dressed bool
}

type drink struct {
	price float64
	name  string
}

type side struct {
	price float64
	name  string
}

type combo struct {
	comboBurger burger
	comboDrink  drink
	comboSide   side
	price       float64
}

func userInputBurger() burger {
	var bun, dressed string
	ret := new(burger)
	fmt.Print("Bun on burger? (y/n) ")
	fmt.Scan(&bun)
	fmt.Print("Dress the burger? (y/n) ")
	fmt.Scan(&dressed)
	if bun == "y" {
		ret.bun = true
		ret.price = 5.0
	} else {
		ret.bun = false
		ret.price = 4.0
	}
	if dressed == "y" {
		ret.dressed = true
	} else {
		ret.dressed = false
	}
	return *ret
}

func userInputDrink() drink {
	var name string
	ret := new(drink)
	fmt.Print("dr.pep, oj, seltzer, or water? ")
	fmt.Scan(&name)
	if name == "water" {
		ret.price = 1.0
	} else {
		ret.price = 2.0
	}
	ret.name = name
	return *ret
}

func userInputSide() side {
	var name string
	ret := new(side)
	fmt.Print("fries, onionrings, salad, or coleslaw? ")
	fmt.Scan(&name)
	if name == "fries" {
		ret.price = 1.0
	} else {
		ret.price = 2.0
	}
	ret.name = name
	return *ret
}

func userInputCombo() combo {
	burg := userInputBurger()
	drnk := userInputDrink()
	sd := userInputSide()
	ret := new(combo)
	ret.comboBurger = burg
	ret.comboDrink = drnk
	ret.comboSide = sd
	ret.price = (burg.price + drnk.price + sd.price) * 0.8
	return *ret
}

func takeOrderFromUser() {
	var burgers []burger
	var drinks []drink
	var sides []side
	var combos []combo
	var answer string
	for {
		fmt.Print("Add a burger/drink/side/combo or quit? ")
		fmt.Scan(&answer)
		if answer == "quit" {
			break
		} else if answer == "burger" {
			burgers = append(burgers, userInputBurger())
		} else if answer == "drink" {
			drinks = append(drinks, userInputDrink())
		} else if answer == "side" {
			sides = append(sides, userInputSide())
		} else {
			combos = append(combos, userInputCombo())
		}
	}
	fmt.Println("burgers: ", burgers)
	fmt.Println("drinks: ", drinks)
	fmt.Println("sides: ", sides)
	fmt.Println("combos: ", combos)
	fmt.Println("Thanks!")
}

func main() {
	takeOrderFromUser()
}
