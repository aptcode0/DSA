package main

import (
	"fmt"
	"math"
)

type edge struct {
	to     int
	weight int
}

func prims(adjs [][]edge) ([]edge, int) {
	n := len(adjs)
	visited := make([]bool, n)
	d := make([]edge, n)
	tot := 0
	for i, _ := range d {
		d[i] = edge{-1, math.MaxInt}
	}
	d[0].weight = 0

	for cnt := 0; cnt < n; cnt++ {
		u := minDist(d, visited)
		if d[u].weight == math.MaxInt {
			break
		}
		tot += d[u].weight
		visited[u] = true

		for _, adj := range adjs[u] {
			if !visited[adj.to] && adj.weight < d[adj.to].weight {
				d[adj.to].weight = adj.weight
				d[adj.to].to = u
			}
		}
	}
	return d, tot
}

func minDist(dists []edge, visited []bool) int {
	mind := math.MaxInt
	u := 0
	for i, d := range dists {
		if !visited[i] && d.weight < mind {
			mind = d.weight
			u = i
		}
	}
	return u
}

// Prims works for undirected graph only, as it assumes that every node connected to every other node
// In directed graph, even though graph is connected, but might not be strongly connected like undirected.
func main() {
	adjs := [][]edge{{{1, 10}, {2, 6}, {3, 5}},
		{{0, 10}, {3, 15}}, {{0, 6}, {3, 4}}, {{0, 5},
			{1, 15}, {2, 4}}}

	// [[2 3 4] [0 3 5] [0 1 10]]
	fmt.Println(prims(adjs))
}
