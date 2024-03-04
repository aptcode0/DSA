package main

import "fmt"

// Find all strongly connected components, single DFS traversal
func tarjan(n int, edges [][]int) [][]int {
	adjs := createGraph[int](n, edges)
	fmt.Println(adjs)
	visited, ins, low, instack := make([]bool, n), make([]int, n), make([]int, n), make([]bool, n)
	var stack []int
	var res [][]int
	time := 0

	var dfs func(int)
	dfs = func(v int) {
		stack = append(stack, v)
		visited[v], instack[v] = true, true
		ins[v], low[v] = time, time
		time++
		for _, adj := range adjs[v] {
			if !visited[adj] {
				dfs(adj)
				low[v] = min(low[v], low[adj])
				continue
			}
			if instack[adj] {
				low[v] = min(low[v], ins[adj])
			}
		}

		if ins[v] == low[v] {
			var u int
			var comp []int
			for {
				u, stack = stack[len(stack)-1], stack[:len(stack)-1]
				instack[u] = false
				comp = append(comp, u)
				if u == v {
					break
				}
			}
			res = append(res, comp)
		}
	}
	for v, _ := range adjs {
		if !visited[v] {
			dfs(v)
		}
	}
	return res
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}}
	// [[0 1 2] [3] [4]]
	fmt.Println(tarjan(5, edges))
}
func createGraph[T comparable](n int, edges [][]T) map[T][]T {
	g := make(map[T][]T, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
	}

	return g
}
