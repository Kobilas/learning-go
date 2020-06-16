package main

import (
	"math/rand"
	"time"
)

func nRandInts(n int) []int {
	ns := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(ns)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = rng.Intn(201) - 100
	}
	return ret
}

func maxInt(arr [100]int) int {
	var max int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func maxIdx(arr [100]int) int {
	var maxIdx, max int
	for i, v := range arr {
		if v > max {
			max = v
			maxIdx = i
		}
	}
	return maxIdx
}

func minInt(arr [100]int) int {
	var min int
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func minIdx(arr [100]int) int {
	var minIdx, min int
	for i, v := range arr {
		if v < min {
			min = v
			minIdx = i
		}
	}
	return minIdx
}

func main() {

}
