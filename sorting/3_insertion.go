package main

import "fmt"

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i-1
		key := nums[i]
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func main() {
	nums := []int{9, 3, 1, 8, 2}

	insertionSort(nums)

	fmt.Println(nums)
}