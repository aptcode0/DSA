package main

import (
	"fmt"
	"slices"
)

func bucketsort(arr []int, bcnt int) {
	shift := slices.Min(arr)
	maxv := slices.Max(arr) - shift
	bsize := maxv / bcnt
	if maxv%bcnt != 0 {
		bsize++
	}

	buckets := make([][]int, bcnt)
	for _, v := range arr {
		idx := (v - shift) / bsize
		buckets[idx] = append(buckets[idx], v)
	}

	for i, _ := range buckets {
		slices.Sort(buckets[i])
	}

	fmt.Println(buckets)
	var sorted []int
	for _, b := range buckets {
		for _, v := range b {
			sorted = append(sorted, v)
		}
	}
	copy(arr, sorted)
}

func main() {
	nums := []int{22, 50, -11, 9, 32, 28, 41, 12}

	bucketsort(nums, 5)

	fmt.Println(nums)
}
