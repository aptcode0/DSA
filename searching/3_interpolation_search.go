package main

import "fmt"

func interpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		pos := low + (target-arr[low])/(arr[high]-arr[low])*(high-low)
		if arr[pos] == target {
			return pos
		}
		if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(interpolationSearch(arr, 8))
	fmt.Println(interpolationSearch(arr, 11))
}
