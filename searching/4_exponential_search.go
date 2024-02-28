package main

import (
	"fmt"
	"slices"
)

func exponentialSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[0] == target {
		return 0
	}
	pos := 1
	for pos < len(arr) && arr[pos] <= target {
		pos *= 2
	}

	idx, isFound := slices.BinarySearch(arr[pos/2:min(len(arr)-1, pos)], target)
	if !isFound {
		return -1
	}
	return pos/2 + idx
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(exponentialSearch(arr, 8))
	fmt.Println(exponentialSearch(arr, 11))
}
