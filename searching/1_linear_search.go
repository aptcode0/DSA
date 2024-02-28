package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 10, 12, 15}
	fmt.Println(linearSearch(arr, 8))
	fmt.Println(linearSearch(arr, 11))
}
