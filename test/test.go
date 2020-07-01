package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func swap(A *[]int, idxOne, idxTwo int) {
	tmp := (*A)[idxOne]
	(*A)[idxOne] = (*A)[idxTwo]
	(*A)[idxTwo] = tmp
}

func evenBeforeOdd(A *[]int) {
	var even, odd int = 0, len(*A) - 1
	for i := 0; i < len(*A)-1 && even < odd; i++ {
		if (*A)[i]%2 == 1 {
			for j := len(*A) - 1; j > even; j-- {
				odd--
				if (*A)[j]%2 == 0 {
					swap(A, i, j)
					break
				}
			}
		}
		even++
	}
}

func evenBeforeOddOutOfPlace(A []int) []int {
	B := make([]int, len(A))
	var even, odd = 0, len(A) - 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			B[even] = A[i]
			even++
		} else {
			B[odd] = A[i]
			odd--
		}
	}
	return B
}

func main() {
	/* var disArray = []int{3, 5, 6, 0, 2, 3, 5, 0}
	fmt.Println(disArray)
	replaceZeroWithHighestEven(&disArray)
	fmt.Println(disArray) */
	rand.Seed(time.Now().UnixNano())
	C := make([]int, 200)
	for i := 0; i < 200; i++ {
		C[i] = rand.Intn(501)
	}
	start := time.Now()
	D := evenBeforeOddOutOfPlace(C)
	fmt.Println(time.Since(start).Nanoseconds())
	fmt.Println(D)
	start = time.Now()
	evenBeforeOdd(&C)
	fmt.Println(time.Since(start).Nanoseconds())
	fmt.Println(C)
}
