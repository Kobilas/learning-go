package main

import (
	"fmt"
	"strconv"
)

// Activity 1
/* package main
import "fmt"
func main() {
	fmt.Println(((5 + 3) % 2) * 9)
} */
// Activity 2
/* package main
import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	fmt.Println("Enter a float: ")
	var fltStr string
	fmt.Scan(&fltStr)
	var flt float64
	flt, _ = strconv.ParseFloat(fltStr, 64)
	fmt.Println("Your float, truncated:", math.Trunc(flt))
} */
// Activity 3
/* package main
import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	var pStr, rStr, nStr, tStr string
	var p, r, n, t float64
	fmt.Print("Enter your initial deposit: ")
	fmt.Scan(&pStr)
	p, _ = strconv.ParseFloat(pStr, 64)
	fmt.Print("Enter your interest rate: ")
	fmt.Scan(&rStr)
	r, _ = strconv.ParseFloat(rStr, 64)
	fmt.Print("Enter how many times interest calculated annually: ")
	fmt.Scan(&nStr)
	n, _ = strconv.ParseFloat(nStr, 64)
	fmt.Print("Enter number of years since initial deposit: ")
	fmt.Scan(&tStr)
	t, _ = strconv.ParseFloat(tStr, 64)
	fmt.Println("V = P(1 + r/n)^nt , where ...")
	fmt.Println("V -- value = ", p*math.Pow((1+r/n), n*t))
	fmt.Println("P -- initial deposit = ", p)
	fmt.Println("r -- interest rate = ", r)
	fmt.Println("n -- # of times per year interest is calculated = ", n)
	fmt.Println("t -- # of years since initial deposit = ", t)
} */
// Activity 4
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var principalStr, rateStr, daysStr string
	var principal, rate, days float64
	fmt.Print("Enter principal amount: ")
	fmt.Scan(&principalStr)
	principal, _ = strconv.ParseFloat(principalStr, 64)
	fmt.Print("Enter rate of interest: ")
	fmt.Scan(&rateStr)
	rate, _ = strconv.ParseFloat(rateStr, 64)
	fmt.Print("Enter number of days for a loan: ")
	fmt.Scan(&daysStr)
	days, _ = strconv.ParseFloat(daysStr, 64)
	fmt.Println("Your interest is :", principal*rate*days/float64(365))
} */
// Activity 5
/* package main
import "fmt"
func main() {
	var a, b, c, d, e int = 1, 3, 5, 7, 11
	var tf bool = e > d
	fmt.Println("e > d:", tf)
	tf = c != a
	fmt.Println("c != a:", tf)
	tf = 3%1 == 0
	fmt.Println("3 % 1 == 0:", tf)
	tf = a > b
	fmt.Println("a > b:", tf)
	tf = c == d
	fmt.Println("c == d:", tf)
	tf = 11%2 == 42
	fmt.Println("11 % 2 == 42:", tf)
} */
// Activity 6
/* package main
import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	var num string
	var numNum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	numNum, _ = strconv.Atoi(num)
	fmt.Println("You selected:", numNum)
	fmt.Println("The boolean of your number is:", numNum != 0)
	fmt.Println("The binary equivalent of your number is:", strconv.FormatInt(int64(numNum), 2))
	fmt.Println("The square root of your number is:", strconv.FormatFloat(math.Sqrt(float64(numNum)), 'f', 3, 64))
} */
// Activity 7
func main() {
	var strNum0, strNum1, strNum2, strNum3, strNum4 string
	var num0, num1, num2, num3, num4 int
	fmt.Print("Enter 5 numbers: ")
	fmt.Scan(&strNum0)
	num0, _ = strconv.Atoi(strNum0)
	fmt.Scan(&strNum1)
	num1, _ = strconv.Atoi(strNum1)
	fmt.Scan(&strNum2)
	num2, _ = strconv.Atoi(strNum2)
	fmt.Scan(&strNum3)
	num3, _ = strconv.Atoi(strNum3)
	fmt.Scan(&strNum4)
	num4, _ = strconv.Atoi(strNum4)
	fmt.Println("Values entered:", num0, ",", num1, ",", num2, ",", num3, ",", num4)
	fmt.Println("Product of ints:", num0*num1*num2*num3*num4)
	fmt.Println("Sum of ints:", num0+num1+num2+num3+num4)
}
