package main

import (
	"fmt"
	"slices"
)

const Digits = 10

func radixsort(arr []int) {
	maxv := slices.Max(arr)
	for curr := 1; curr < maxv; curr *= Digits {
		fmt.Println(arr, curr)
		countsort(arr, curr)
	}
}

func countsort(arr []int, div int) {
	counts := make([]int, Digits)
	for _, v := range arr {
		counts[(v/div)%Digits]++
	}

	start := 0
	for i, c := range counts {
		counts[i] = start
		start += c
	}

	sorted := make([]int, len(arr))
	for _, v := range arr {
		pos := (v / div) % Digits
		sorted[counts[pos]] = v
		counts[pos]++
	}
	copy(arr, sorted)
}

func main() {
	nums := []int{985, 1345, 3932, 3422, 204, 6784, 123}

	radixsort(nums)

	fmt.Println(nums)
}
