package main

import (
	"cmp"
	"fmt"
	"slices"
)

// Minimum spanning tree (MST)

func kruskal(n int, edges [][]int) [][]int {
	slices.SortFunc(edges, func(e1, e2 []int) int {
		return cmp.Compare(e1[2], e2[2])
	})

	ds := NewDisjointSet(n)
	var mst [][]int
	for _, edge := range edges {
		if ds.Union(edge[0], edge[1]) {
			mst = append(mst, edge)
		}
	}
	return mst
}

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i, _ := range parent {
		parent[i] = i
	}
	return &DisjointSet{parent, rank}
}

func (s DisjointSet) Union(x int, y int) bool {
	xp, yp := s.Find(x), s.Find(y)
	if xp == yp {
		return false
	}
	if s.rank[xp] > s.rank[yp] {
		s.parent[yp] = xp
	} else if s.rank[yp] > s.rank[xp] {
		s.parent[xp] = yp
	} else {
		s.parent[yp] = xp
		s.rank[xp]++
	}
	return true
}

func (s DisjointSet) Find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

func main() {
	edges := [][]int{{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4}}
	fmt.Println(kruskal(4, edges))
}
