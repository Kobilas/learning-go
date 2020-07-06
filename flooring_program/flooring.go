//Darius
//Matthew
//Jasmine

package main

import "fmt"

//Make a map with each state and a random tax
tax = map[string]float64{

//file name for orders is Orders_MMDDYYYY.txt.
//Example, January 1st, 2017 = Orders_01012017.txt



//Edit an order method



//Add an order method - Darius

func getMonth() int {
	var month int
	fmt.Println("What month will this be? Input will be as MM") //MM
	fmt.Scanln(&month)
	for month > 12 || month < 1 {
		fmt.Println("Wrong input. Try again.")
		fmt.Scanln(&month)
	}
	return month
}

func getDay() int {
	var day int
	fmt.Println("What day will this be? Input will be as DD") //DD
	fmt.Scanln(&day)
	for day > 31 || day < 1 {
		fmt.Println("Wrong input. Try again.")
		fmt.Scanln(&day)
	}
}

func getYear() int {
	var year int
	fmt.Println("What year will this be? Input will be as YYYY") //YYYY
	fmt.Scanln(&year)
	for year < 2019 {
		fmt.Println("Wrong input. Try again.")
		fmt.Scanln(&year)
	}
}

func getDate() time.Time {
	now := time.Now()

	month := getMonth()
	day := getDay()
	year := getYear()

	orderTime := time.Date(year, time.Month(month), day, 0, 0, 00, 0, time.Local)

	for !orderTime.After(now) {
		fmt.Println("Wrong inputs. The date for the order has to be after today. Try again")
		month = getMonth()
		day = getDay()
		year = getYear()

		orderTime = time.Date(year, time.Month(month), day, 0, 0, 00, 0, time.Local)
	}
	return orderTime
}

func getCustName() string {
	nameRegex := regexp.MustCompile("[a-zA-Z',.]+\s*[0-9]+")
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What is your name?")
	if scanner.Scan() {
		name := scanner.Text()
	}
	for !nameRegex.MatchString(name) {
		if scanner.Scan() {
			name = scanner.Text()
		}
	}
	return name
}

func getState() (string, float64) {
	file, err = os.Open("/taxes.txt")
	if err != nil {
		panic("File not found.")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var state string
	fmt.Println("What state are you looking for? Type the whole name")
	fmt.Scanln(&state)
	state = strings.ToUpper(state)
	var exist boolean = false

	for _, line := range lines {
		split := strings.Split(line,",")
		
		if(strings.ToUpper(split[1]) == state) {
			exist = true
			break
		}
	}
	
	var tax string
	for !exist {
		fmt.Println("State doesn't exist, try again.")
		fmt.Scanln(&state)
		state = strings.ToUpper(state)

		for _, line := range lines {
			split := strings.Split(line,",")
			
			if(strings.ToUpper(split[1]) == state) {
				tax = split[2]
				exist = true
				break
			}
		}
	}
	file.Close()
	return state, strconv.ParseFloat(tax, 64)
}

func getFloorType() (string, float64) {
	fmt.Println("Wood : $0.50")
	fmt.Println("Laminate : $0.70")
	fmt.Println("Vinyl : $1.00")
	fmt.Println("Tile : $2.00")
	fmt.Println("Carpet: $3.00")
	
	var flooring string
	var price float64
	for {
		fmt.Println("What kind of flooring do you want?")
		fmt.Scanln(&flooring)
		switch strings.ToUpper(flooring) {
		case "WOOD":
			price = .50
			break
		case "LAMINATE":
			price = .70
			break
		case "VINYL":
			price = 1.00
			break
		case "TILE":
			price = 2.00
			break
		case "CARPET":
			price = 3.00
			break
		default:
			fmt.Println("Wrong input. Try again.")
		}
	}
	return flooring, price
}

func getArea() float64 {

	var area float64 = 100


	for {
		fmt.Println("Enter the area required. Must be a minimum of 100 sq ft.")
		fmt.Scanln(&area)

		if (area >= 100) {
			break
		} else {
			fmt.Println("Wrong input. Try again.")
		}
	}
	return area
}
  
func addOrder() {
	orderTime := getDate()
	name := getCustName()
func addOrder() {
	values := []string{}
	orderTime := getDate()
	values = append(values, getCustName())
	state, tax := getState()
	values = append(values, state)
	values = append(values, strconv.FormatFloat(tax, 'f', 2, 64))
	flooring, price := getFloorType()
	values[3] = flooring
	area := getArea()
	values[4] = strconv.FormatFloat(area, 'f', 2, 64)

	costPerSqFt := area * price
	
	values[5] = strconv.FormatFloat(costPerSqFt, 'f', 2, 64)

	laborCostPerSqFt = //need products method

  results := //join together values with join 

}

}



//Display orders method



//User interface method


func main() {
	var input string
	for input != "5"{
		fmt.Println("Flooring Program:")
		fmt.Println("1. Display Orders")
		fmt.Println("2. Add an Order")
		fmt.Println("3. Edit an Order")
    	fmt.Println("4. Remove an Order")
   		fmt.Println("5. Quit")
  		fmt.Print("Choose a number 1-5: ")
    	fmt.Scan(&input)

    switch input {
		case "1":
			
	
		case "2":
			addOrder()
		
		case "3":
			editOrder()

		case "4":
			removeOrder()
			    
        case "5":
            fmt.Println("Thank you for using the Flooring Program")
	
		default:
			fmt.Println("Please choose a valid value between 1-4 or 5 to exit")
        
	}
}