package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func getUserDuration() (time.Duration, error) {
	var inp string
	var days, hours, minutes, seconds int
	var err error
	fmt.Print("Number of days: ")
	fmt.Scan(&inp)
	days, err = strconv.Atoi(inp)
	if err != nil {
		return 0, errors.New("getUserDuration: days must be in integer format")
	}
	fmt.Print("Number of hours: ")
	fmt.Scan(&inp)
	hours, err = strconv.Atoi(inp)
	if err != nil {
		return 0, errors.New("getUserDuration: hours must be in integer format")
	}
	fmt.Print("Number of minutes: ")
	fmt.Scan(&inp)
	minutes, err = strconv.Atoi(inp)
	if err != nil {
		return 0, errors.New("getUserDuration: minutes must be in integer format")
	}
	fmt.Print("Number of seconds: ")
	fmt.Scan(&inp)
	seconds, err = strconv.Atoi(inp)
	if err != nil {
		return 0, errors.New("getUserDuration: seconds must be in integer format")
	}
	dur, err := time.ParseDuration(strconv.Itoa(((days * 24) + hours)) +
		"h" + strconv.Itoa(minutes) + "m" + strconv.Itoa(seconds) + "s")
	if err != nil {
		return 0, errors.New("getUserDuration: critical error parsing input to duration")
	}
	return dur, nil
}

func getUserDate(includeTime bool) (time.Time, error) {
	var inp string
	var year, month, day int
	var err error
	fmt.Print("Enter year: ")
	fmt.Scan(&inp)
	year, err = strconv.Atoi(inp)
	if err != nil {
		return time.Time{}, errors.New("getUserDate: year must be in integer format")
	}
	fmt.Print("Enter month: ")
	fmt.Scan(&inp)
	month, err = strconv.Atoi(inp)
	if err != nil {
		return time.Time{}, errors.New("getUserDate: month must be in integer format")
	}
	fmt.Print("Enter day: ")
	fmt.Scan(&inp)
	day, err = strconv.Atoi(inp)
	if err != nil {
		return time.Time{}, errors.New("getUserDate: day must be in integer format")
	}
	if includeTime {
		var hour, minute, second int
		fmt.Print("Enter hour: ")
		fmt.Scan(&inp)
		hour, err = strconv.Atoi(inp)
		if err != nil {
			return time.Time{}, errors.New("getUserDate: hour must be in integer format")
		}
		fmt.Print("Enter minute: ")
		fmt.Scan(&inp)
		minute, err = strconv.Atoi(inp)
		if err != nil {
			return time.Time{}, errors.New("getUserDate: minute must be in integer format")
		}
		fmt.Print("Enter second: ")
		fmt.Scan(&inp)
		second, err = strconv.Atoi(inp)
		if err != nil {
			return time.Time{}, errors.New("getUserDate: second must be in integer format")
		}
		return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local), nil
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func durationCalculator() {
	var inp string
	var dur1, dur2 time.Duration
	var err error
	for {
		fmt.Println("For duration 1, please enter:")
		dur1, err = getUserDuration()
		if err == nil {
			break
		}
	}
	for {
		fmt.Println("For duration 2, please enter:")
		dur2, err = getUserDuration()
		if err == nil {
			break
		}
	}
	for {
		fmt.Print("Addition or subtraction (add/sub): ")
		fmt.Scan(&inp)
		if inp == "add" {
			days := int((dur1.Hours() + dur2.Hours()) / 24)
			resHours := int(dur1.Hours()+dur2.Hours()) - (days * 24)
			hours := resHours % 24
			minutes := int(dur1.Minutes()+dur2.Minutes()) -
				(hours * 60) - (days * 24 * 60)
			seconds := int(dur1.Seconds()+dur2.Seconds()) -
				(minutes * 60) - (hours * 60 * 60) - (days * 24 * 60 * 60)
			fmt.Println(days, "days", hours, "hours", minutes, "minutes",
				seconds, "seconds")
			fmt.Println((dur1.Hours()+dur2.Hours())/24, "days")
			fmt.Println(dur1.Hours()+dur2.Hours(), "hours")
			fmt.Println(dur1.Minutes()+dur2.Minutes(), "minutes")
			fmt.Println(dur1.Seconds()+dur2.Seconds(), "seconds")
			break
		} else if inp == "sub" {
			days := int((dur1.Hours() - dur2.Hours()) / 24)
			resHours := int(dur1.Hours()-dur2.Hours()) - (days * 24)
			hours := resHours % 24
			minutes := int(dur1.Minutes()-dur2.Minutes()) -
				(hours * 60) - (days * 24 * 60)
			seconds := int(dur1.Seconds()-dur2.Seconds()) -
				(minutes * 60) - (hours * 60 * 60) - (days * 24 * 60 * 60)
			fmt.Println(days, "days", hours, "hours", minutes, "minutes",
				seconds, "seconds")
			fmt.Println((dur1.Hours()-dur2.Hours())/24, "days")
			fmt.Println(dur1.Hours()-dur2.Hours(), "hours")
			fmt.Println(dur1.Minutes()-dur2.Minutes(), "minutes")
			fmt.Println(dur1.Seconds()-dur2.Seconds(), "seconds")
			break
		} else {
			continue
		}
	}
}

func dateCalculator() {
	var inp string
	var usrDate time.Time
	var usrDur time.Duration
	var err error
	for {
		fmt.Println("Enter information for date:")
		usrDate, err = getUserDate(true)
		if err == nil {
			break
		}
	}
	for {
		fmt.Println("Enter information for duration:")
		usrDur, err = getUserDuration()
		if err == nil {
			break
		}
	}
	for {
		fmt.Print("Addition or subtraction (add/sub): ")
		fmt.Scan(&inp)
		if inp == "add" {
			newDate := usrDate.Add(usrDur)
			fmt.Println(newDate)
			break
		} else if inp == "sub" {
			newDate := usrDate.Add(-usrDur)
			fmt.Println(newDate)
			break
		} else {
			continue
		}
	}
}

func ageCalculator() {
	var usrDate1, usrDate2 time.Time
	var err error
	for {
		fmt.Println("Enter first date:")
		usrDate1, err = getUserDate(false)
		if err == nil {
			break
		}
	}
	for {
		fmt.Println("Enter second date:")
		usrDate2, err = getUserDate(false)
		if err == nil {
			break
		}
	}
	diff := usrDate2.Sub(usrDate1)
	diffYears := int(diff.Hours() / 24 / 365)
	diffPostYears := int(diff.Hours()) / 24 % 365
	fmt.Println()
	fmt.Println(diffYears, "years", int(diffPostYears/30), "months",
		diffPostYears%30, "days")
	fmt.Println(int(diff.Hours()/24/30), "months",
		int(diff.Hours()/24)%30, "days")
	fmt.Println(int(diff.Hours()/24/7), "weeks",
		int(diff.Hours()/24)%7, "days")
	fmt.Println(int(diff.Hours()/24), "days")
	fmt.Println(diff.Hours(), "hours")
	fmt.Println(diff.Minutes(), "minutes")
	fmt.Println(diff.Seconds(), "seconds")
}

func main() {
	var inp string
	for {
		fmt.Println()
		fmt.Println("1. Time duration calculator")
		fmt.Println("2. Date calculator")
		fmt.Println("3. Age calculator")
		fmt.Println("q. Quit")
		fmt.Print("Enter an option: ")
		fmt.Scan(&inp)
		switch inp {
		case "1":
			durationCalculator()
		case "2":
			dateCalculator()
		case "3":
			ageCalculator()
		case "q":
			break
		}
	}
}
