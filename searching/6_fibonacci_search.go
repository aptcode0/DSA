package main

import "fmt"

func fibonacciSearch(arr []int, target int) int {
	prev, curr := 0, 1
	nxt := curr + prev
	n := len(arr)

	for nxt < n {
		prev = curr
		curr = nxt
		nxt = prev + curr
	}
	offset := -1
	for curr > 1 {
		i := min(offset+prev, n-1)
		if arr[i] == target {
			return i
		}
		if arr[i] < target {
			offset = i // skip lesser values
			// one down
			nxt = curr
			curr = prev
			prev = nxt - curr
		} else {
			// two down	- shrinking incrementing size more as we are ahead of target
			nxt = prev
			curr = curr - prev
			prev = nxt - curr
		}
	}
	if curr == 1 && arr[offset+1] == target {
		return offset + 1
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(fibonacciSearch(arr, 8))
	fmt.Println(fibonacciSearch(arr, 11))
}
