package main

import "fmt"

type result struct {
	dist   int
	parent int
}

func bfs(adjs [][]int, s int) []result {
	results := make([]result, len(adjs))
	results[s].dist = 0
	results[s].parent = -1
	seen := make([]bool, len(adjs))
	seen[s] = true
	q := []int{s}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adjs[u] {
			if !seen[v] {
				seen[v] = true
				results[v].dist = results[u].dist + 1
				results[v].parent = u
				q = append(q, v)
			}
		}
	}
	return results
}

func main() {
	adjs := [][]int{{1}, {2}, {0, 1}}
	fmt.Println(bfs(adjs, 0))
}
