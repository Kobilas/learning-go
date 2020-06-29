package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type book struct {
	id             int
	title          string
	author         string
	pages          int
	publishingYear int
	franchise      bool
}

func (b book) String() string {
	return fmt.Sprintf("{%s %s %d %d %t}",
		b.title, b.author, b.pages, b.publishingYear, b.franchise)
}

func create(f *os.File) {
	var inp string
	var title, author string
	var pages, publishingYear int
	var franchise bool
	fmt.Print("Enter title: ")
	fmt.Scan(&title)
	fmt.Print("Enter author: ")
	fmt.Scan(&author)
	fmt.Print("Enter number of pages: ")
	fmt.Scan(&inp)
	pages, _ = strconv.Atoi(inp)
	fmt.Print("Enter publishing year: ")
	fmt.Scan(&inp)
	publishingYear, _ = strconv.Atoi(inp)
	fmt.Print("Part of franchise? (y/n) ")
	fmt.Scan(&inp)
	franchise = inp == "y"
	tmp := book{title, author, pages, publishingYear, franchise}
	bw := bufio.NewWriter(f)
	bw.WriteString(tmp.String())
}

func read(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func update(f *os.File) {

}

func main() {
	var inp string
	var f *os.File
	var err error
	var highest int
	f, err = os.Open("collection.txt")
	if err != nil {
		err = nil
		f, err = os.Create("collection.txt")
		if err != nil {
			fmt.Println("Error creating file after attempting to open")
		}
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tmp := scanner.Text()

	}
	for {
		fmt.Println()
		fmt.Println("Main Menu")
		fmt.Println("1 - Create entry")
		fmt.Println("2 - Read entries")
		fmt.Println("3 - Update entry")
		fmt.Println("4 - Destroy entry")
		fmt.Println("q - quit")
		fmt.Print("Select an option: ")
		fmt.Scan(&inp)
		switch inp {
		case "1":
			create(f)
		case "2":
			read(f)
		case "3":
		case "4":
		case "q":
			break
		}
	}
}
