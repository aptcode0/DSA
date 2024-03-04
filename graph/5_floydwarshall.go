package main

import (
	"fmt"
	"math"
)

/* All pairs shortest path (APSP) */

func floydwarshall(graph [][]int) [][]int {
	n := len(graph)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		copy(dp[i], graph[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][k] != math.MaxInt && dp[k][j] != math.MaxInt && dp[i][k]+dp[k][j] < dp[i][j] {
					dp[i][j] = dp[i][k] + dp[k][j]
				}
			}
		}
	}
	return dp
}

func main() {
	graph := [][]int{
		{0, 5, math.MaxInt64, 10},
		{math.MaxInt64, 0, 3, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, 0, 1},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, 0},
	}
	fmt.Println(floydwarshall(graph))
}
