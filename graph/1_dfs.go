package main

import "fmt"

type state struct {
	start int
	end   int
	color int
}

func dfs(adjs [][]int, state []state, node int, timer *int) {
	state[node].start = *timer
	*timer++

	state[node].color = 1
	for _, adj := range adjs[node] {
		if state[adj].color == 0 {
			dfs(adjs, state, adj, timer)
		}
	}
	state[node].color = 2
	state[node].end = *timer
	*timer++
}

func main() {
	adjs := [][]int{{1, 2}, {2}, {}}
	states := make([]state, 3)
	timer := 0
	dfs(adjs, states, 0, &timer)
	fmt.Println(states)
}
