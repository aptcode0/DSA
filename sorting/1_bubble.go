package main

import "fmt"

func bubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			} 
		}
		if !swapped {
			break
		}
	}
}

func main() {
	nums := []int{9, 3, 1, 8, 2}

	bubbleSort(nums)

	fmt.Println(nums)
}