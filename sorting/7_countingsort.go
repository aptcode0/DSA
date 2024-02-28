package main

import (
	"fmt"
)

func countingsort(arr []int) {
	if len(arr) == 0 {
		return
	}
	minv, maxv := arr[0], arr[0]
	for _, v := range arr {
		minv = min(minv, v)
		maxv = max(maxv, v)
	}
	size := maxv - minv + 1
	counts := make([]int, size)
	for _, v := range arr {
		counts[v-minv]++
	}

	start := 0
	for i, c := range counts {
		counts[i] = start
		start += c
	}

	sorted := make([]int, len(arr))
	for _, v := range arr {
		sorted[counts[v-minv]] = v
		counts[v-minv]++
	}
	copy(arr, sorted)
}

func main() {
	nums := []int{-3, -3, 0, 8, 3, 1, 7, 8, 2, 6}

	countingsort(nums)

	fmt.Println(nums)
}
