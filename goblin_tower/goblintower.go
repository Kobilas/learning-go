package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type hero struct {
	maxHealth  int
	health     int
	maxDefense int
	defense    int
	attack     int
	gold       int
	potions    [5]int
}

type goblin struct {
	health  int
	attack  int
	defense int
}

func generateHero(gold int) hero {
	var health, defense int = rand.Intn(11) + 20, rand.Intn(3) + 1
	return hero{health, health, defense, defense, rand.Intn(5) + 1, gold, [5]int{2, 2, 2, 2, 2}}
}

func generateGoblin() goblin {
	return goblin{rand.Intn(6) + 5, rand.Intn(2) + 2, rand.Intn(2) + 1}
}

func (myHero *hero) fight(gob *goblin) {
	var gobFirst, heroFirst bool = true, true
	gob.printStats()
	for myHero.health > 0 && gob.health > 0 {
		if gob.defense > 0 {
			if myHero.attack > gob.defense {
				fmt.Println("The goblin's defense takes", gob.defense, "damage.")
			} else {
				fmt.Println("The goblin's defense takes", myHero.attack, "damage.")
			}
			gob.defense -= myHero.attack
		}
		if gob.defense < 0 {
			gobFirst = false
			fmt.Println("You break the goblin's defenses and deal", int(math.Abs(float64(gob.defense))), "damage!")
			gob.health += gob.defense
			gob.defense = 0
		} else {
			if gobFirst {
				fmt.Println("The goblin's defense was broken!")
				gobFirst = false
			}
			fmt.Println("You deal", myHero.attack, "damage.")
			gob.health -= myHero.attack
		}
		if myHero.defense > 0 {
			if gob.attack > myHero.defense {
				fmt.Println("Your defense takes", myHero.defense, "damage.")
			} else {
				fmt.Println("Your defense takes", gob.attack, "damage.")
			}
			myHero.defense -= gob.attack
		}
		if myHero.defense < 0 {
			heroFirst = false
			fmt.Println("The goblin breaks your defenses and deals", int(math.Abs(float64(myHero.defense))), "damage!")
			myHero.health += myHero.defense
			myHero.defense = 0
		} else {
			if heroFirst {
				fmt.Println("Your defense was broken!")
				heroFirst = false
			}
			fmt.Println("You take", gob.attack, "damage.")
			myHero.health -= gob.attack
		}
	}
}

func (myHero hero) checkInv() int {
	var count int
	for _, v := range myHero.potions {
		if v == 2 {
			count++
		}
	}
	return count
}

func (myHero *hero) quaff() {
	if myHero.checkInv() == 0 {
		fmt.Println("Ran out of potions")
	} else {
		for i := 0; i < len(myHero.potions); i++ {
			if myHero.potions[i] == 2 {
				myHero.potions[i] = 0
				if myHero.health <= myHero.maxHealth-2 {
					myHero.health += 2
				} else if myHero.health == myHero.maxHealth-1 {
					myHero.health++
				}
				break
			}
		}
		fmt.Println("You chug a potion.")
	}
}

func (myHero *hero) shop() {
	var ansStr string
	var ansInt int
	fmt.Println("Welcome to the Potion Shop!")
	for {
		fmt.Print("How many potions would you like to buy? ")
		fmt.Scan(&ansStr)
		ansInt, _ = strconv.Atoi(ansStr)
		if ansInt == 0 {
			fmt.Println("Okay then.")
			return
		} else if ansInt <= (5-myHero.checkInv()) && (myHero.gold-(ansInt*4)) >= 0 {
			myHero.gold -= ansInt * 4
			for i := 0; i < ansInt; i++ {
				for j := 0; j < len(myHero.potions); j++ {
					if myHero.potions[j] == 0 {
						myHero.potions[j] += 2
						break
					}
				}
			}
			fmt.Println("Thanks for shopping!")
			return
		} else {
			if ansInt > (5 - myHero.checkInv()) {
				fmt.Println("It doesn't look like you have enough space...")
			} else if (myHero.gold - (ansInt * 4)) < 0 {
				fmt.Println("What do I look like a charity?")
			}
		}
	}
}

func (myHero hero) printStats() {
	fmt.Println("Hero health: " + strconv.Itoa(myHero.health) + "/" + strconv.Itoa(myHero.maxHealth))
	fmt.Println("Hero defense: " + strconv.Itoa(myHero.defense) + "/" + strconv.Itoa(myHero.maxDefense))
	fmt.Println("Hero attack:", myHero.attack)
}

func (gob goblin) printStats() {
	fmt.Println("Goblin health: " + strconv.Itoa(gob.health) + "/" + strconv.Itoa(gob.health))
	fmt.Println("Goblin defense: " + strconv.Itoa(gob.defense) + "/" + strconv.Itoa(gob.defense))
	fmt.Println("Goblin attack:", gob.attack)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var goblinsSlain, steps, level, maxLevel, reserves, encounter int
	var ans string
	for ans != "n" {
		fmt.Println("Welcome to the Goblin Tower")
		myHero := generateHero(reserves)
		myHero.printStats()
		steps = 0
		level = 0
		for myHero.health > 0 {
			ans = ""
			fmt.Println("You take a step.")
			steps++
			if steps == 10 {
				level++
				steps = 0
				if level > maxLevel {
					maxLevel = level
				}
				myHero.defense = myHero.maxDefense
				fmt.Println("You have reached the next level!")
				fmt.Println("All defense regained, you encounter a potion shop.")
				fmt.Println("Current gold:", myHero.gold, "; Current slots open:", 5-myHero.checkInv())
				fmt.Print("Buy potion(s) for 4 gold each? (y/n) ")
				fmt.Scan(&ans)
				if ans == "y" {
					myHero.shop()
				}
			} else {
				encounter = rand.Intn(2)
				if encounter == 1 {
					fmt.Println("A goblin! You draw your sword.")
					enemyGoblin := generateGoblin()
					myHero.fight(&enemyGoblin)
					if myHero.health > 0 {
						goblinsSlain++
						myHero.gold += 2
					}
				}
			}
			for ans != "n" && myHero.health > 0 && myHero.checkInv() > 0 {
				myHero.printStats()
				fmt.Println("Potions left:", myHero.checkInv())
				fmt.Print("Drink potion? (y/n) ")
				fmt.Scan(&ans)
				if ans == "y" {
					myHero.quaff()
				}
			}
		}
		fmt.Println("You have been slain...")
		fmt.Print("Play again? (y/n) ")
		fmt.Scan(&ans)
	}
	fmt.Println("Highest level reached:", maxLevel)
	fmt.Println("Goblins slain:", goblinsSlain)
}
