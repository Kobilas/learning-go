package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	value int
	suit  string
}

type deck struct {
	cards []card
}

type player struct {
	points     int
	hidden     int
	aces       int
	hiddenAces int
}

func generatePlayer(pointValue, hiddenValue int) *player {
	newPlayer := new(player)
	newPlayer.points = pointValue
	newPlayer.hidden = hiddenValue
	return newPlayer
}

func generateCard(value int, suit string) card {
	return card{value, suit}
}

func generateDeck() *deck {
	var suits = [4]string{"spade", "heart", "diamond", "club"}
	ret := new(deck)
	for _, v := range suits {
		for j := 0; j < 13; j++ {
			ret.cards = append(ret.cards, generateCard(j+1, v))
		}
	}
	return ret
}

func (d *deck) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *deck) dealCard() card {
	var ret card
	ret, d.cards = d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
	return ret
}

func (p *player) hit(fromDeck *deck, hide bool) {
	newCard := fromDeck.dealCard()
	if !hide {
		if newCard.value == 1 {
			p.aces++
		} else if newCard.value > 10 {
			p.points += 10
		} else {
			p.points += newCard.value
		}
	} else {
		if newCard.value == 1 {
			p.hiddenAces++
		} else if newCard.value > 10 {
			p.hidden += 10
		} else {
			p.hidden += newCard.value
		}
	}
}

func (p player) checkPoints() int {
	var ret int = p.points + p.hidden
	if ret < 11 && p.aces < 1 {
		ret += (p.aces + p.hiddenAces) * 11
	} else if p.aces > 1 {
		ret += p.aces + p.hiddenAces
	}
	return ret
}

func checkWinner(playerOne, playerTwo player, stay bool) bool {
	var checkFlag bool = false
	var one, two int = playerOne.checkPoints(), playerTwo.checkPoints()
	fmt.Println("Final totals: Player One:", one, "; Player Two:", two)
	if one > 21 {
		fmt.Println("You busted")
		return true
	} else if one == 21 {
		fmt.Println("You hit blackjack")
		return true
	} else {
		checkFlag = true
	}
	if two > 21 {
		fmt.Println("Player Two busted")
		return true
	} else if two == 21 {
		fmt.Println("Player Two hit blackjack")
		return true
	}

	if stay && checkFlag {
		if one > two {
			fmt.Println("Player One wins!")
			return true
		} else if one < two {
			fmt.Println("Player Two wins!")
			return true
		} else if one == two {
			fmt.Println("Tie!")
			return true
		}
	}
	return false
}

func main() {
	var ans string
	var playerOne, dealer *player = generatePlayer(0, -1), generatePlayer(0, 0)

	for {
		fmt.Println("The name of the game is BLACKJACK")
		var myDeck *deck = generateDeck()
		myDeck.shuffleDeck()
		playerOne.hit(myDeck, false)
		dealer.hit(myDeck, false)
		playerOne.hit(myDeck, false)
		dealer.hit(myDeck, true)
		fmt.Println("Dealer totals:", dealer.points, "points;", dealer.aces, "aces")
		for ans != "stay" {
			fmt.Println("Your totals:", playerOne.points, "points;", playerOne.aces, "aces")
			fmt.Println("Hit or stay? ")
			fmt.Scan(&ans)
			if ans == "hit" {
				playerOne.hit(myDeck, false)
				if checkWinner(*playerOne, *dealer, false) == true {
					break
				}
			}
		}
		for dealer.checkPoints() < 17 {
			if dealer.checkPoints() < 17 {
				dealer.hit(myDeck, true)
				if checkWinner(*playerOne, *dealer, false) == true {
					break
				}
			}
		}
		checkWinner(*playerOne, *dealer, true)
		fmt.Print("Play again? (y/n) ")
		fmt.Scan(&ans)
		if ans == "n" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
