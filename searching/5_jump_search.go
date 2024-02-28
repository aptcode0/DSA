package main

import (
	"fmt"
	"math"
)

func jumpSearch(arr []int, target int) int {
	n := len(arr)
	jumpSize := int(math.Sqrt(float64(n)))

	prev := 0
	step := jumpSize
	for end := min(step, n) - 1; arr[end] < target; end = min(step, n) - 1 {
		prev = step
		step += jumpSize
		if prev >= n {
			return -1
		}
	}
	step = min(step+1, n)
	for prev < step && arr[prev] < target {
		prev++
	}
	if arr[prev] == target {
		return prev
	}

	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(jumpSearch(arr, 8))
	fmt.Println(jumpSearch(arr, 11))
}
