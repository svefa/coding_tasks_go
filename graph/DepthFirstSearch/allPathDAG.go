// 797. Given a directed acyclic graph (DAG) of n nodes labeled
// from 0 to n - 1, find all possible paths from node 0 to node n - 1
// and return them in any order.

//    0 -> 1
//    |    |
//    v    v
//    2 -> 3
// Input: graph = [[1,2],[3],[3],[]]
// Output: [[0,1,3],[0,2,3]]

package mainA

import "fmt"

func allPathsSourceTarget(graph [][]int) (res [][]int) {
	p := []int{0}
	var dfs func(s, t int, p []int)
	dfs = func(s, t int, p []int) {
		if s == t {
			res = append(res, append([]int{}, p...))
			return
		}
		for _, i := range graph[s] {
			p = append(p, i)
			dfs(i, t, p)
			p = p[0 : len(p)-1]
		}
	}
	dfs(0, len(graph)-1, p)
	return
}

/********************************/

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)

}
