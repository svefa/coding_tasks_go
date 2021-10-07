// 797. Given a directed acyclic graph (DAG) of n nodes labeled
// from 0 to n - 1, find all possible paths from node 0 to node n - 1
// and return them in any order.

//    0 -> 1
//    |    |
//    v    v
//    2 -> 3
// Input: graph = [[1,2],[3],[3],[]]
// Output: [[0,1,3],[0,2,3]]

package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)

	paths := [][]int{}
	s, t := 0, n-1
	path := []int{s}
	//dfs(s,t,graph,path,&paths)
	bfs(s, t, graph, path, &paths)
	return paths
}

func bfs(s, t int, graph [][]int, path []int, paths *[][]int) {
	q := [][]int{{s}}

	for len(q) > 0 {
		path, q = q[0], q[1:]
		i := path[len(path)-1]
		if i == t {
			*paths = append(*paths, append([]int{}, path...))
		}
		for _, n := range graph[i] {
			q = append(q, append([]int{}, append(path, n)...))
		}
	}
}

/********************************/

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)

}
