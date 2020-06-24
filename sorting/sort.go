package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func selectionSort(A *[]int) {
	n := len(*A)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (*A)[j] < (*A)[minIdx] {
				minIdx = j
			}
		}
		swap(A, minIdx, i)
	}
}

func mergeSort(A *[]int, l, r int) {
	if l < r {
		m := int(math.Floor(float64((l + r) / 2)))
		mergeSort(A, l, m)
		mergeSort(A, m+1, r)
		merge(A, l, m, r)
	}
}

func merge(A *[]int, l, m, r int) {
	n1, n2 := m-l+1, r-m
	L, R := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = (*A)[l+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = (*A)[m+1+i]
	}
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			(*A)[k] = L[i]
			i++
		} else {
			(*A)[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		(*A)[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		(*A)[k] = R[j]
		j++
		k++
	}
}

func quickSortLomuto(A *[]int, lo, hi int) {
	if lo < hi {
		p := lomutoPartition(A, lo, hi)
		quickSortLomuto(A, lo, p-1)
		quickSortLomuto(A, p+1, hi)
	}
}

func quickSortHoare(A *[]int, lo, hi int) {
	if lo < hi {
		p := hoarePartition(A, lo, hi)
		quickSortHoare(A, lo, p-1)
		quickSortHoare(A, p+1, hi)
	}
}

func lomutoPartition(A *[]int, lo, hi int) int {
	pivot := (*A)[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if (*A)[j] < pivot {
			swap(A, i, j)
			i++
		}
	}
	swap(A, i, hi)
	return i
}

func hoarePartition(A *[]int, lo, hi int) int {
	pivot := (*A)[int(math.Floor(float64((hi+lo)/2)))]
	i := lo
	j := hi
	for {
		for (*A)[i] < pivot {
			i++
		}
		for (*A)[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		swap(A, i, j)
		i++
		j--
	}
}

func insertionSort(A *[]int) {
	n := len(*A)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && (*A)[j-1] > (*A)[j]; j-- {
			swap(A, j, j-1)
		}
	}
}

func bubbleSort(A *[]int) {
	n := len(*A)
	for i := 0; i < n; i++ {
		for j := 0; j < (n - i - 1); j++ {
			if (*A)[j] > (*A)[j+1] {
				swap(A, j, j+1)
			}
		}
	}
}

func swap(A *[]int, idx1, idx2 int) {
	tmp := (*A)[idx1]
	(*A)[idx1] = (*A)[idx2]
	(*A)[idx2] = tmp
}

func generateUnsorted(n int) []int {
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = rand.Intn(101)
	}
	return A
}

func main() {
	rand.Seed(time.Now().UnixNano())
	A := generateUnsorted(20)
	fmt.Println(A)
	bubbleSort(&A)
	fmt.Println(A)
	B := generateUnsorted(20)
	fmt.Println(B)
	insertionSort(&B)
	fmt.Println(B)
	C := generateUnsorted(20)
	fmt.Println(C)
	quickSortLomuto(&C, 0, len(C)-1)
	fmt.Println(C)
	D := generateUnsorted(20)
	fmt.Println(D)
	quickSortHoare(&D, 0, len(D)-1)
	fmt.Println(D)
	E := generateUnsorted(20)
	fmt.Println(E)
	mergeSort(&E, 0, len(E)-1)
	fmt.Println(E)
	F := generateUnsorted(20)
	fmt.Println(F)
	selectionSort(&F)
	fmt.Println(F)
}
