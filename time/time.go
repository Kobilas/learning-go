package main

import (
	"fmt"
	"strconv"
	"time"
)

func checkLeap() bool {
	var inpStr string
	var inp int
	fmt.Print("Enter a year: ")
	fmt.Scan(&inpStr)
	inp, _ = strconv.Atoi(inpStr)
	usrTime := time.Date(inp, 12, 31, 0, 0, 0, 0, time.Local)
	dayOfYear := usrTime.YearDay()
	if dayOfYear == 366 {
		return true
	}
	return false
}

func main() {
	now := time.Now()
	fmt.Println("Current date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	_, week := now.ISOWeek()
	fmt.Println("Current week:", week)
	fmt.Println("Current weekday:", now.Weekday())
	fmt.Println("Current day of the year:", now.YearDay())
	fmt.Println("Current day of the month:", now.Day())
	fmt.Println("Current day of the week:", (now.Day()+1)%7)

}
