package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	intVar := 7
	fmt.Printf("intVar type: %T\n", intVar)
	fltVar := 66.6
	fmt.Printf("fltVar type: %T\n", fltVar)
	strVar := "string_here"
	fmt.Printf("strVar type: %T\n", strVar)

	var mathClass complex64 = 7 + 7i
	var investmentAccountBalance float32 = 7070707070707070
	var partsPerBillion float64 = 0.000000000000000000007
	var investmentInsurancePlan float32 = 888888888
	var numberOfPlantsOnPlanet float64 = 99999999999999999999999
	fmt.Println("math problem from class:", mathClass)
	fmt.Println("customer investment account balance:", investmentAccountBalance)
	fmt.Println("parts per billion of nitrogen monoxide:", partsPerBillion)
	fmt.Println("insurance for investment account:", investmentInsurancePlan)
	fmt.Println("number of plants on Earth 2:", numberOfPlantsOnPlanet)

	mathClass += complex64(1)
	fmt.Println("math problem from class + 2:", mathClass)
	investmentAccountBalance -= float32(666.0)
	fmt.Println("account balance after poor investment:", investmentAccountBalance)
	partsPerBillion += float64(3.5)
	fmt.Println("parts per billion after pollution:", partsPerBillion)
	investmentInsurancePlan -= investmentInsurancePlan
	fmt.Println("insurance gone after bad investment :( :", investmentInsurancePlan)
	numberOfPlantsOnPlanet -= float64(42)
	fmt.Println("plants per pollution:", numberOfPlantsOnPlanet)

	{
		a, b, c := 7, 42, 69
		fmt.Println("c plus a, minus b:", c+a-b)
		fmt.Println("c minus a, times b:", (c-a)*b)
		fmt.Println("c divided by a, with the result divided by b:", c/a/b)
		fmt.Println("add b to the remainder of c divided by a:", c%a+b)
	}

	{
		var d int8 = 100
		var e int = 70
		fmt.Println("d =", d)
		fmt.Println("e =", e)
		fmt.Println("d == e:", int(d) == e)
		fmt.Println("d != e:", int(d) != e)
	}

	{
		a, b, c := 3.73, 6.9, 42.7
		fmt.Println("ceiling of a multiplied by c:", math.Ceil(a*c))
		fmt.Println("floor of a divided by b:", math.Floor(a/b))
		fmt.Println("square root of b, rounded to three places:", math.Floor(math.Sqrt(b)*1000)/1000)
		fmt.Println("ceiling of  c plus b, divided by a:", math.Ceil((c+b)/a))
		fmt.Println("c squared:", math.Pow(c, 2))
		fmt.Println("highest value of a, b, and c:", math.Max(a, math.Max(b, c)))
		fmt.Println("lowest value of a, b, and c:", math.Min(a, math.Min(b, c)))
	}

	fmt.Println("rand.Intn(100):", rand.Intn(100))
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	fmt.Println("rand.Intn(100) with seed:", generator.Intn(100))
	fmt.Println("rand.Intn(100) with seed 2: ", generator.Intn(100))

	{
		var m, n, p int = 3, 7, 11
		a := m > n
		fmt.Println("m > n:", a)
		b := n <= p
		fmt.Println("n <= p:", b)
		c := m == p
		fmt.Println("m == p:", c)
		d := m != p
		fmt.Println("m != p:", d)
	}
}
