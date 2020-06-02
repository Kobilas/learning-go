// Activity 1
/* package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)
	var ansStr string
	var ans, correct, incorrect, wrongStreak, total int
	var num1, num2, op int
	for i := 0; i < 10; i++ {
		num1 = gen.Intn(21)
		num2 = gen.Intn(21)
		op = gen.Intn(3)
		// 0 -- addition
		// 1 -- subtraction
		// 2 -- modulo
		wrongStreak = 0
		switch op {
		case 0:
			fmt.Print(num1, "+", num2, "= ? ")
			fmt.Scan(&ansStr)
			ans, _ = strconv.Atoi(ansStr)
			if ans == num1+num2 {
				fmt.Println("Correct")
				wrongStreak = 0
				correct++
			} else {
				fmt.Println("Incorrect, answer:", num1+num2)
				wrongStreak++
				incorrect++
			}
		case 1:
			fmt.Print(num1, "-", num2, "= ? ")
			fmt.Scan(&ansStr)
			ans, _ = strconv.Atoi(ansStr)
			if ans == num1-num2 {
				fmt.Println("Correct")
				wrongStreak = 0
				correct++
			} else {
				fmt.Println("Incorrect, answer:", num1-num2)
				wrongStreak++
				incorrect++
			}
		case 2:
			fmt.Print(num1, "%", num2, "= ? ")
			fmt.Scan(&ansStr)
			ans, _ = strconv.Atoi(ansStr)
			if ans == num1%num2 {
				fmt.Println("Correct")
				wrongStreak = 0
				correct++
			} else {
				fmt.Println("Incorrect, answer:", num1%num2)
				wrongStreak++
				incorrect++
			}
		}
		total++
		fmt.Println("Questions answered:", total)
		fmt.Println("Accuracy rate:", math.Trunc((float64(correct)/float64(total))*100))
		if wrongStreak >= 3 {
			fmt.Println("You lost...")
			return
		}
	}
	fmt.Println("You win!")
} */

// Activity 2
/* package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)
func main() {
	var toShorten, ret string
	var comLoc int
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)
	fmt.Print("Enter url to shorten: ")
	fmt.Scan(&toShorten)
	ret = "http://shrt.com/"
	comLoc = strings.Index(toShorten, ".com")
	ret += toShorten[comLoc+5 : comLoc+6]
	ret += toShorten[len(toShorten)-1 : len(toShorten)]
	ret += strconv.Itoa(gen.Intn(100))
	fmt.Print("Your shortened URL:", ret)
} */

// Activity 3
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var phoneStr string
	fmt.Print("Enter phone number: ")
	fmt.Scan(&phoneStr)
	_, err := strconv.Atoi(phoneStr)
	if err != nil {
		if len(phoneStr) != 12 {
			fmt.Println(phoneStr, "is an invalid phone number")
			return
		}
		if phoneStr[3] == '-' && phoneStr[7] == '-' {
			fmt.Println(phoneStr, "is a valid phone number")
			return
		} else {
			fmt.Println(phoneStr, "is an invalid phone number")
			return
		}
	} else {
		fmt.Println(phoneStr[:3]+"-"+phoneStr[3:6]+"-"+phoneStr[6:], "is a valid phone number")
	}
} */

// Activity 4
/* package main

import (
	"fmt"
	"strings"
)

func main() {
	var emailStr string
	for {
		fmt.Print("Enter email address: ")
		fmt.Scan(&emailStr)
		atLoc := strings.Index(emailStr, "@")
		comLoc := strings.Index(emailStr, ".com")
		if atLoc == -1 || comLoc == -1 {
			fmt.Println("Valid format: False\nTry again")
			continue
		} else {
			fmt.Println("Valid format: True")
			fmt.Println("Domain:", emailStr[atLoc+1:comLoc])
			fmt.Println("Identifier:", emailStr[:atLoc])
			break
		}
	}
} */

// Activity 5
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ccard0 int
	var ucard0 int
	var cPoints, shownPoints, uPoints int
	var handNum int = 3
	var hitstay string
	ns := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(ns)
	fmt.Println("Hand 1 :")
	ccard0 = gen.Intn(10) + 1
	cPoints += ccard0
	ucard0 = gen.Intn(10) + 1
	uPoints += ucard0
	fmt.Println("Your hand:", uPoints)
	fmt.Println("Hand 2 :")
	ccard0 = gen.Intn(10) + 1
	cPoints += ccard0
	shownPoints += ccard0
	fmt.Println("Computer hand:", "X +", shownPoints)
	ucard0 = gen.Intn(10) + 1
	uPoints += ucard0
	fmt.Println("Your hand:", uPoints)
	for {
		fmt.Println("Hand", handNum, ": ")
		fmt.Println("Your hand:", uPoints)
		fmt.Println("Hit or stay? ")
		fmt.Scan(&hitstay)
		if hitstay == "hit" {
			ucard0 = gen.Intn(10) + 1
			uPoints += ucard0
			if uPoints > 21 {
				break
			} else if uPoints == 21 {
				fmt.Println("Computer hand:", cPoints)
				fmt.Println("Your hand:", uPoints)
				if uPoints == cPoints {
					fmt.Println("Tie game!")
					return
				} else {
					fmt.Println("Player wins!")
					return
				}
			} else {
				continue
			}
		} else if hitstay == "stay" {
			for {
				if cPoints < 17 {
					ccard0 = gen.Intn(10) + 1
					cPoints += ccard0
					fmt.Println("Computer hand:", "X+", shownPoints)

				} else {
					break
				}
			}
			break
		} else {
			continue
		}
	}
	fmt.Println("upoints:", uPoints, "; cpoints:", cPoints)
	if uPoints > 21 {
		fmt.Println("Player loses!")
	}
	if cPoints > 21 {
		fmt.Println("Computer loses!")
	}
	if uPoints == cPoints {
		fmt.Println("Tie game!")
	}
	if uPoints == 21 {
		fmt.Println("Player wins!")
	} else if cPoints == 21 {
		fmt.Println("Computer wins!")
	} else {
		if uPoints > cPoints {
			fmt.Println("Player wins!")
		} else {
			fmt.Println("Computer wins!")
		}
	}
}
