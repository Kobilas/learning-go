// Activity 1
/* package main
import "fmt"
func main() {
	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range alphabet {
		fmt.Print(string(c))
	}
	fmt.Println("")
} */

// Activity 2
/* package main
import "fmt"
func main() {
	for i := 1; i <= 100; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	for i := 2; i <= 100; i += 2 {
		fmt.Print(i, " ")
	}
} */

// Activity 3

// Activity 4
/* package main
import "fmt"
func main() {
	var sum int
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 0 - 100:", sum)
} */

// Activity 5
/* package main
import "fmt"
func main() {
	for i := 100; i <= 1000; i += 50 {
		fmt.Print(i, " ")
	}
	fmt.Println()
	var i uint = 100
	for {
		if i > 1000 {
			break
		}
		fmt.Print(i, " ")
		i += 50
	}
} */

// Activity 6
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var nStr string
	fmt.Print("Enter a number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	nSqr := n * n
	for i := 1; i <= nSqr; i++ {
		fmt.Print(n*i, " ")
	}
} */

// Activity 7
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var xStr, yStr string
	fmt.Print("Enter an X number: ")
	fmt.Scan(&xStr)
	x, _ := strconv.Atoi(xStr)
	fmt.Print("Enter a Y number: ")
	fmt.Scan(&yStr)
	y, _ := strconv.Atoi(yStr)
	var product int = 1
	for i := 1; i <= y; i++ {
		product *= x
	}
	println("X^Y:", product)
} */

// Activity 8
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var nStr string
	var sum, prevSum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	for i := 1; i <= n; i++ {
		prevSum += i
		sum += prevSum
	}
	fmt.Println("Sum of (1) + (1+2) + ...:", sum)
} */

// Activity 9
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var nStr string
	fmt.Print("Enter a number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	fmt.Println("Number of digits in number:", len(nStr))
	fmt.Println("First and last digits in number:", nStr[:1], nStr[len(nStr):])
	var sum, prod int = 0, 1
	for _, c := range nStr {
		num, _ := strconv.Atoi(string(c))
		sum += num
		prod *= num
	}
	fmt.Println("Sum of digits in number:", sum)
	fmt.Println("Product of digits in number:", prod)
	var divBool bool = false
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			divBool = true
		}
	}
	fmt.Println("Number is prime:", !divBool)
	var factorial int = 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("n! :", factorial)
} */

// Activity 10
/* package main
import (
	"fmt"
	"strconv"
)
func main() {
	var nStr string
	fmt.Print("Enter a number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
} */

// Activity 11
/* package main
import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	var mStr, nStr string
	fmt.Print("Enter a number: ")
	fmt.Scan(&mStr)
	m, _ := strconv.Atoi(mStr)
	fmt.Print("Enter another number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	var min int = int(math.Min(float64(m), float64(n)))
	var gcf int
	for i := 2; i <= min; i++ {
		if m%i == 0 && n%i == 0 {
			gcf = i
		}
	}
	fmt.Println("Greatest common factor:", gcf)
} */

// Activity 12
/* package main
import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	var mStr, nStr string
	fmt.Print("Enter a number: ")
	fmt.Scan(&mStr)
	m, _ := strconv.Atoi(mStr)
	fmt.Print("Enter a number: ")
	fmt.Scan(&nStr)
	n, _ := strconv.Atoi(nStr)
	min := math.Min(float64(m), float64(n))
	var lcm, i int = 1, int(min)
	for {
		if i%m == 0 && i%n == 0 {
			lcm = i
			break
		}
		i++
	}
	fmt.Println("Least common multiple:", lcm)
} */
