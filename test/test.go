package main

import "fmt"

func replaceZeroWithHighestEven(arr *[]int) {
	var highest int
	for _, val := range *arr {
		if val > highest && val%2 == 0 {
			highest = val
		}
	}
	for i, val := range *arr {
		if val == 0 {
			(*arr)[i] = highest
		}
	}
}

func main() {
	var disArray = []int{3, 5, 6, 0, 2, 3, 5, 0}
	fmt.Println(disArray)
	replaceZeroWithHighestEven(&disArray)
	fmt.Println(disArray)
}
