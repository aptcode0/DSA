package main

import "fmt"

// Strongly connected components Time: O(V+E) Space: O(V)
func kosaraju[T comparable](n int, edges [][]T) [][]T {
	adjs := createGraph(n, edges)
	adjsT := createGraph(n, reverse(edges))

	visited := make(map[T]bool, n)
	stack := NewStack[T]()
	for node, _ := range adjs {
		if !visited[node] {
			dfs(adjs, node, visited, stack)
		}
	}

	clear(visited)

	var res [][]T
	for !stack.IsEmpty() {
		v := stack.Pop()
		component := NewStack[T]()
		if !visited[v] {
			dfs(adjsT, v, visited, component)
			res = append(res, component.ToArray())
		}
	}
	return res
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}}
	fmt.Println(kosaraju(5, edges))

	flights := [][]string{{"A", "B"}, {"B", "C"}, {"C", "A"}, {"B", "D"}, {"D", "E"}}
	fmt.Println(kosaraju(5, flights))
}

func createGraph[T comparable](n int, edges [][]T) map[T][]T {
	g := make(map[T][]T, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
	}
	return g
}

func reverse[T comparable](edges [][]T) [][]T {
	res := make([][]T, len(edges))
	for i, e := range edges {
		res[i] = []T{e[1], e[0]}
	}
	return res
}

func dfs[T comparable](adjs map[T][]T, curr T, visited map[T]bool, stack *Stack[T]) {
	visited[curr] = true
	for _, adj := range adjs[curr] {
		if !visited[adj] {
			dfs(adjs, adj, visited, stack)
		}
	}
	stack.Push(curr)
}

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

func (s *Stack[T]) Pop() T {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}

func (s Stack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	}
	return s.values[len(s.values)-1]
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s Stack[T]) ToArray() []T {
	return s.values
}
