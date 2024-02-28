package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quicksort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quicksortHelper(nums, 0, len(nums)-1)
}

func quicksortHelper(nums []int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(nums, low, high)
	quicksortHelper(nums, low, pivot-1)
	quicksortHelper(nums, pivot+1, high)
}

func partition(nums []int, low, high int) int {
	randIdx := rand.Intn(high-low+1) + low

	nums[high], nums[randIdx] = nums[randIdx], nums[high]
	pivot := nums[high]

	j := low
	for i := low; i < high; i++ {
		if nums[i] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j], nums[high] = nums[high], nums[j]
	return j
}


func main() {
	nums := []int{9, 3, 1, 7, 8, 2, 6}

	quicksort(nums)

	fmt.Println(nums)
}