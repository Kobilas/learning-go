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

func create(f *os.File) {
	var inp string
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
	tmp := bird{birdID, name, scientificName, wingspan, region, diet}
	bw := bufio.NewWriter(f)
	bw.WriteString(tmp.toString())
	bw.Flush()
	birdID++
	birdCount++
}

func read(f *os.File) {
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func update(f *os.File) {
	var idToUpdate string
	lines := make([]string, birdCount)
	fmt.Print("Enter ID of bird to update: ")
	fmt.Scan(&idToUpdate)
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if line[0] != idToUpdate {
			lines = append(lines, strings.Join(line, ",")+"\n")
		} else {
			fmt.Println("Updating", strings.Join(line, ", "))
			var tmp string
			fmt.Print("Update name? (" + line[1] + ") ")
			fmt.Scan(&tmp)
			if tmp != "" {
				line[1] = tmp
			}
			fmt.Print("Update scientific name? (" + line[2] + ") ")
			fmt.Scan(&tmp)
			if tmp != "" {
				line[2] = tmp
			}
			fmt.Print("Update wingspan? (" + line[3] + ") ")
			fmt.Scan(&tmp)
			if tmp != "" {
				line[3] = tmp
			}
			fmt.Print("Update region? (" + line[4] + ") ")
			fmt.Scan(&tmp)
			if tmp != "" {
				line[4] = tmp
			}
			fmt.Print("Update diet? (" + line[5] + ") ")
			fmt.Scan(&tmp)
			if tmp != "" {
				line[5] = tmp
			}
			lines = append(lines, strings.Join(line, ",")+"\n")
			fmt.Print("Updated to", strings.Join(line, ", "))
		}
	}
	rewrite(f, lines)
}

func destroy(f *os.File) {
	var idToDelete string
	lines := make([]string, birdCount-1)
	fmt.Print("Enter ID of bird to delete: ")
	fmt.Scan(&idToDelete)
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if line[0] != idToDelete {
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
			create(f)
		case "2":
			read(f)
		case "3":
			update(f)
		case "4":
			destroy(f)
		case "q":
			os.Exit(0)
		}
	}
}
