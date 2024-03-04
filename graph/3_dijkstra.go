package main

import (
	"fmt"
	"math"
)

/* single-source shortest paths problem. */

type edge struct {
	to     int
	weight int
}

func dijkstra(adjs [][]edge, source int) []int {
	n := len(adjs)
	visited := make([]bool, n)
	d := make([]int, n)
	for i, _ := range d {
		d[i] = math.MaxInt
	}
	d[source] = 0

	for cnt := 0; cnt < n-1; cnt++ {
		u := minDist(d, visited)
		if d[u] == math.MaxInt {
			break
		}
		visited[u] = true

		for _, adj := range adjs[u] {
			if !visited[adj.to] && d[u]+adj.weight < d[adj.to] {
				d[adj.to] = d[u] + adj.weight
			}
		}
	}
	return d
}

func minDist(dists []int, visited []bool) int {
	mind := math.MaxInt
	u := 0
	for i, d := range dists {
		if !visited[i] && d < mind {
			mind = d
			u = i
		}
	}
	return u
}

func main() {
	adjs := [][]edge{{{1, 3}, {2, 8}}, {{2, 3}}, {{1, 2}}}
	fmt.Println(dijkstra(adjs, 0))
}
