package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func printBoard(boardState [3][3]int) {
	var tmp string
	for i := 0; i < len(boardState); i++ {
		for j := 0; j < len(boardState[i]); j++ {
			if j != 2 {
				tmp = strconv.Itoa(boardState[i][j])
				fmt.Print(tmp + "|")
			} else {
				fmt.Println(boardState[i][j])
			}
		}
		if i != 2 {
			fmt.Println("-----")
		}
	}
	fmt.Println()
}

func printIndices() {
	fmt.Println("0|1|2")
	fmt.Println("-----")
	fmt.Println("3|4|5")
	fmt.Println("-----")
	fmt.Println("6|7|8")
	fmt.Println()
}

func takeAnswer(player bool) int {
	// player one = false
	// player two = true
	var playerAns string
	var ans int
	if player {
		fmt.Print("Player two, where would you like to move? (enter an index) ")
	} else {
		fmt.Print("Player one, where would you like to move? (enter an index) ")
	}
	fmt.Scan(&playerAns)
	ans, _ = strconv.Atoi(playerAns)
	return ans
}

func checkWin(boardState [3][3]int, player bool) bool {
	// player one = false
	// player two = true
	var row, colOne, colTwo, colThree, rightDiagonal, leftDiagonal int
	// check for win in the columns
	for i := 0; i < len(boardState); i++ {
		if player {
			if boardState[i][0] == 2 {
				colOne++
			}
			if boardState[i][1] == 2 {
				colTwo++
			}
			if boardState[i][2] == 2 {
				colThree++
			}
		} else {
			if boardState[i][0] == 1 {
				colOne++
			}
			if boardState[i][1] == 1 {
				colTwo++
			}
			if boardState[i][2] == 1 {
				colThree++
			}
		}
	}
	if colOne == 3 || colTwo == 3 || colThree == 3 {
		return true
	}
	// check for win in the rows
	for i := 0; i < len(boardState); i++ {
		for j := 0; j < len(boardState[i]); j++ {
			if player {
				if boardState[i][j] == 2 {
					row++
				}
			} else {
				if boardState[i][j] == 1 {
					row++
				}
			}
		}
		if row == 3 {
			return true
		}
		row = 0
	}
	// check for win in the diagonals
	for i := 0; i < len(boardState); i++ {
		if player {
			if boardState[i][i] == 2 {
				rightDiagonal++
			}
			if boardState[i][2-i] == 2 {
				leftDiagonal++
			}
		} else {
			if boardState[i][i] == 1 {
				rightDiagonal++
			}
			if boardState[i][2-i] == 1 {
				leftDiagonal++
			}
		}
	}
	if rightDiagonal == 3 || leftDiagonal == 3 {
		return true
	}
	return false
}

func checkBlock(boardState [3][3]int) int {
	// player one = false
	// player two = true
	var row, colOne, colTwo, colThree, rightDiagonal, leftDiagonal int
	var poss, tmp []int
	// check for win in the columns
	for i := 0; i < len(boardState); i++ {
		if boardState[i][0] == 1 {
			colOne++
			tmp = append(tmp, boardState[i][0])
		}
		if boardState[i][1] == 1 {
			colTwo++
		}
		if boardState[i][2] == 1 {
			colThree++
		}
	}
	if colOne == 2 || colTwo == 2 || colThree == 2 {
		return true
	}
	// check for win in the rows
	for i := 0; i < len(boardState); i++ {
		for j := 0; j < len(boardState[i]); j++ {
			if boardState[i][j] == 1 {
				row++
			}
		}
		if row == 2 {
			return true
		}
		row = 0
	}
	// check for win in the diagonals
	for i := 0; i < len(boardState); i++ {
		if boardState[i][i] == 1 {
			rightDiagonal++
		}
		if boardState[i][2-i] == 1 {
			leftDiagonal++
		}
	}
	if rightDiagonal == 2 || leftDiagonal == 2 {
		return true
	}
	return false
}

func checkAvailability(boardState [3][3]int, target int) bool {
	var x, y int
	x = int(math.Trunc(float64(target) / 3.0))
	y = target % 3
	if boardState[x][y] == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Welcome to tic-tac-toe")
	fmt.Print("Would you like to play vs player or computer? ")
	ns := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(ns)
	var board [3][3]int
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var ansOne, ansTwo, x, y int
	var player string
	var oneWins, twoWins bool
	fmt.Scan(&player)
	printIndices()
	if player == "player" {
		for i := 0; ; i++ {
			printBoard(board)
			ansOne = takeAnswer(false)
			ansTwo = takeAnswer(true)
			x = int(math.Trunc(float64(ansOne) / 3.0))
			y = ansOne % 3
			board[x][y] = 1
			x = int(math.Trunc(float64(ansTwo) / 3.0))
			y = ansTwo % 3
			board[x][y] = 2
			if i >= 2 {
				oneWins = checkWin(board, false)
				twoWins = checkWin(board, true)
				if oneWins {
					printBoard(board)
					fmt.Println("PLAYER ONE WINS!")
					break
				} else if twoWins {
					printBoard(board)
					fmt.Println("PLAYER TWO WINS!")
					break
				}
			}
			if i >= 8 {
				if !oneWins && !twoWins {
					printBoard(board)
					fmt.Println("TIE!")
					break
				}
			}
		}
	} else {
		for i := 0; ; i++ {
			printBoard(board)
			for {
				ansOne = takeAnswer(false)
				if !checkAvailability(board, ansOne) {
					fmt.Println("Please enter available space")
					continue
				} else {
					break
				}
			}
			x = int(math.Trunc(float64(ansOne) / 3.0))
			y = ansOne % 3
			board[x][y] = 1
			/*
				a[ansOne] = a[len(a)-1]
				a = a[:len(a)-1]
			*/
			for {
				ansTwo = rng.Intn(len(a))
				if !checkAvailability(board, ansTwo) {
					continue
				} else {
					break
				}
			}
			/*
				a[ansTwo] = a[len(a)-1]
				a = a[:len(a)-1]
			*/
			x = int(math.Trunc(float64(ansTwo) / 3.0))
			y = ansTwo % 3
			board[x][y] = 2
			if i >= 2 {
				oneWins = checkWin(board, false)
				twoWins = checkWin(board, true)
				if oneWins {
					printBoard(board)
					fmt.Println("PLAYER ONE WINS!")
					break
				} else if twoWins {
					printBoard(board)
					fmt.Println("COMPUTER WINS!")
					break
				}
			}
			if i >= 8 {
				if !oneWins && !twoWins {
					printBoard(board)
					fmt.Println("TIE!")
					break
				}
			}
		}
	}
}
