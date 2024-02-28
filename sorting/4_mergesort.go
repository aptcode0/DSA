package main

import "fmt"

func mergeSort(nums []int) {
	tmp := make([]int, len(nums))
	divideAndMerge(nums, tmp, 0, len(nums)-1)
}

func divideAndMerge(nums, tmp []int, left, right int) {
	if right <= left {
		return
	}
	mid := left + (right - left) / 2
	divideAndMerge(nums, tmp, left, mid)
	divideAndMerge(nums, tmp, mid+1, right)
	merge(nums, tmp, left, mid, right)
}

func merge(nums, tmp []int, left, mid, right int) {
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	for i := left; i <= right; i++ {
		nums[i] = tmp[i]
	}
}


func main() {
	nums := []int{9, 3, 1, 8, 2}

	mergeSort(nums)

	fmt.Println(nums)
}