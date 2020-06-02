// Activity 1
/* Code Block A
package main
import "fmt"
func main() {
	fmt.Println("Hello, world!")
} */
/* Code Block B
package main
import "fmt"
func main() {
	fmt.Println("Go is fun!"); fmt.Println("Go is also easy.");
} */
/* Code Block C
package main
import "fmt"
func main() {
	var _name1 string = "Rebecca"
	var _nameAmp string = "Roberts"
	fmt.Println(_name1)
	fmt.Println(_nameAmp)
} */
// Activity 2
/* package main
import "fmt"
func main() {
	fmt.Println("Go is fun!")
	//fmt.Println("Go is also easy!")
} */
// Activity 3
/* package main
import "fmt"
func main() {
	var output string = "I love Go!"
	fmt.Println(output)
} */
// Activity 4
package main

import "fmt"

func main() {
	fmt.Println("Enter the state in which you were born: ")
	var birthState string
	fmt.Scan(&birthState)
	fmt.Println("Enter the state in which you currently reside: ")
	var currState string
	fmt.Scan(&currState)
	fmt.Println("You entered:", birthState, "=>", currState)
}
