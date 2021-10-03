package main

import "fmt"

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	v := map[int]bool{}
	g := map[int][]int{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
	}
	fmt.Println(g, "\n")
	var dfs func(p int) bool
	dfs = func(p int) bool {
		if v[p] == true {
			fmt.Println("cicl", p, v)
			return false
		}
		v[p] = true
		if len(g[p]) == 0 && p != destination {
			fmt.Println("dend")
			return false
		}
		for i := 0; i < len(g[p]); i++ {
			if !dfs(g[p][i]) {
				return false
			}

		}
		v[p] = false
		return true
	}
	return dfs(source)
}

func main() {
	n := 2
	edges := [][]int{{0, 1}, {1, 1}}
	source := 0
	destination := 1

	n = 4
	edges = [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	source = 0
	destination = 3

	res := leadsToDestination(n, edges, source, destination)
	print(res, "\n")
}
