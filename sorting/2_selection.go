package main

import "fmt"

func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i+1; j < n; j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		nums[minIdx], nums[i] = nums[i], nums[minIdx]
	}
}

func main() {
	nums := []int{9, 3, 1, 8, 2}

	selectionSort(nums)

	fmt.Println(nums)
}