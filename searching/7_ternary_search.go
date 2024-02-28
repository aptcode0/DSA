package main

import "fmt"

func ternarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}
		if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(ternarySearch(arr, 8))
	fmt.Println(ternarySearch(arr, 11))
}
