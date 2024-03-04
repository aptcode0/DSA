package main

import (
	"errors"
	"fmt"
	"math"
)

/* single-source shortest paths problem. handle negative weight cycle too. */

func bellmandFord(n int, edges [][]int, source int) ([]int, error) {
	d := make([]int, n)
	for i, _ := range d {
		d[i] = math.MaxInt
	}
	d[source] = 0

	for cnt := 0; cnt < n-1; cnt++ {
		for _, e := range edges {
			u, v, w := e[0], e[1], e[2]
			if d[u] != math.MaxInt && d[u]+w < d[v] {
				d[v] = d[u] + w
			}
		}
	}

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if d[u] != math.MaxInt && d[u]+w < d[v] {
			return nil, errors.New("negative weight cycle found")
		}
	}

	return d, nil
}

func main() {
	adjs := [][]int{{0, 1, 3}, {0, 2, -4}, {1, 2, 3}, {2, 1, -2}}
	fmt.Println(bellmandFord(3, adjs, 0))

	adjs = [][]int{{0, 1, 3}, {0, 2, -4}, {1, 2, 3}, {2, 1, -6}}
	fmt.Println(bellmandFord(3, adjs, 0))
}
