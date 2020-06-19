package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dish struct {
	name       string
	morsels    int
	beingEaten bool
}

type person struct {
	name   string
	eating bool
}

func genDish(name string) dish {
	return dish{name, rand.Intn(6) + 5, false}
}

func eatFood(member *person, food *dish) {
	fmt.Println(member.name, "is enjoying some", food.name)
	food.morsels--
	food.beingEaten = true
	member.eating = true
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	food.beingEaten = false
	member.eating = false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var dinner = [5]dish{
		genDish("tacos"),
		genDish("sushi"),
		genDish("pizza"),
		genDish("quesadillas"),
		genDish("crepes"),
	}

	var family = [3]person{
		person{"Kai", false},
		person{"Prithvi", false},
		person{"Matt", false},
	}

	var food = make(chan dish)

	var pers = make(chan person)

	fmt.Println("Bon appétit!")

	for dinner[0].morsels != 0 && dinner[1].morsels != 0 && dinner[2].morsels != 0 && dinner[3].morsels != 0 && dinner[4].morsels != 0 {

		eatFood(&family[0], &dinner[rand.Intn(5)])
		eatFood(&family[1], &dinner[rand.Intn(5)])
		eatFood(&family[2], &dinner[rand.Intn(5)])
	}

	fmt.Println("That was delicious!")
}
