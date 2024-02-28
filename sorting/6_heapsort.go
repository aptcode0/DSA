package main

import "fmt"

func heapsort(nums []int) {
	n := len(nums)
	for i := n/2-1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n-1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

func heapify(nums []int, n, parent int) {
	left := 2 * parent + 1
	right := 2 * parent + 2

	pos := parent
	if left < n && nums[pos] < nums[left] {
		pos = left
	}
	if right < n && nums[pos] < nums[right] {
		pos = right
	}
	if pos == parent {
		return
	}
	nums[pos], nums[parent] = nums[parent], nums[pos]
	heapify(nums, n, pos)
}


func main() {
	nums := []int{9, 3, 1, 7, 8, 2, 6}

	heapsort(nums)

	fmt.Println(nums)
}