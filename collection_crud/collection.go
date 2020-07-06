package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bird struct {
	id             int
	name           string
	scientificName string
	wingspan       int
	region         string
	diet           string
}

var birdID int
var birdCount int

func (b bird) toString() string {
	return fmt.Sprintf("%d,%s,%s,%d,%s,%s\n",
		b.id, b.name, b.scientificName, b.wingspan, b.region, b.diet)
}

func setCurrentIDAndGetLength(f *os.File) {
	var greatest, count int
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		idStr := strings.Split(scanner.Text(), ",")[0]
		id, _ := strconv.Atoi(idStr)
		if id > greatest {
			greatest = id
		}
		count++
	}
	birdID = greatest + 1
	birdCount = count
}

func Create(f *os.File, name, scientificName, region, diet string, wingspan int) {
	tmp := bird{birdID, name, scientificName, wingspan, region, diet}
	bw := bufio.NewWriter(f)
	bw.WriteString(tmp.toString())
	bw.Flush()
	birdID++
	birdCount++
}

func Read(f *os.File) {
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func Update(f *os.File, id, name, scientificName, wingspan, region, diet string) {
	lines := make([]string, birdCount)
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if line[0] != id {
			lines = append(lines, strings.Join(line, ",")+"\n")
		} else {
			fmt.Println("Updating", strings.Join(line, ", "))
			if name != "" {
				line[1] = name
			}
			if scientificName != "" {
				line[2] = scientificName
			}
			if wingspan != "" {
				line[3] = wingspan
			}
			if region != "" {
				line[4] = region
			}
			if diet != "" {
				line[5] = diet
			}
			lines = append(lines, strings.Join(line, ",")+"\n")
			fmt.Print("Updated to", strings.Join(line, ", "))
		}
	}
	rewrite(f, lines)
}

func Destroy(f *os.File, id string) {
	lines := make([]string, birdCount-1)
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if line[0] != id {
			lines = append(lines, strings.Join(line, ",")+"\n")
		} else {
			fmt.Println("Deleting", strings.Join(line, ", "))
		}
	}
	rewrite(f, lines)
	birdCount--
}
func rewrite(f *os.File, lines []string) {
	// Truncate changes size of file, does not change I/O offset, hence the seek
	f.Truncate(0)
	f.Seek(0, 0)
	bw := bufio.NewWriter(f)
	for _, val := range lines {
		bw.WriteString(val)
	}
	bw.Flush()
}

func main() {
	var inp string
	var f *os.File
	var err error
	f, err = os.OpenFile("collection.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("error: when opening or creating file in read/write mode")
	}
	defer f.Close()
	setCurrentIDAndGetLength(f)
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
			var name, scientificName, region, diet string
			var wingspan int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter scientific name: ")
			fmt.Scan(&scientificName)
			fmt.Print("Enter wingspan: ")
			fmt.Scan(&inp)
			wingspan, _ = strconv.Atoi(inp)
			fmt.Print("Enter region: ")
			fmt.Scan(&region)
			fmt.Print("Enter diet: ")
			fmt.Scan(&diet)
			Create(f, name, scientificName, region, diet, wingspan)
		case "2":
			Read(f)
		case "3":
			fmt.Print("Enter ID of bird to delete: ")
			fmt.Scan(&inp)
			var name, scientificName, wingspan, region, diet string
			fmt.Print("Update name? ")
			fmt.Scan(&name)
			fmt.Print("Update scientific name? ")
			fmt.Scan(&scientificName)
			fmt.Print("Update wingspan? ")
			fmt.Scan(&wingspan)
			fmt.Print("Update region? ")
			fmt.Scan(&region)
			fmt.Print("Update diet? ")
			fmt.Scan(&diet)
			Update(f, inp, name, scientificName, wingspan, region, diet)
		case "4":
			fmt.Print("Enter ID of bird to delete: ")
			fmt.Scan(&inp)
			Destroy(f, inp)
		case "q":
			os.Exit(0)
		}
	}
}
