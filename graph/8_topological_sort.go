package main

import "fmt"

func topologicalSort(n int, edges [][]int) []int {
	adjs, ins := buildDependencies(edges)
	q := findStartNodes(n, ins)

	orderedNodes := order(adjs, ins, q)
	if len(orderedNodes) != n {
		return nil
	}
	return orderedNodes
}

func buildDependencies(edges [][]int) (map[int][]int, map[int]int) {
	adjs := map[int][]int{}
	ins := map[int]int{}
	for _, p := range edges {
		adjs[p[0]] = append(adjs[p[0]], p[1])
		ins[p[1]]++
	}

	return adjs, ins
}

func findStartNodes(n int, ins map[int]int) []int {
	var q []int
	for node := 0; node < n; node++ {
		if ins[node] == 0 {
			q = append(q, node)
		}
	}
	return q
}

func order(adjs map[int][]int, ins map[int]int, q []int) []int {
	var nodes []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		nodes = append(nodes, node)
		for _, adj := range adjs[node] {
			ins[adj]--
			if ins[adj] == 0 {
				q = append(q, adj)
			}
		}
	}
	return nodes
}

func main() {
	adjs := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	fmt.Println(topologicalSort(4, adjs))
}
