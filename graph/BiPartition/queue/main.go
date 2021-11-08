package main

func isBipartite(graph [][]int) bool {
	v := make([]int, len(graph))
	var q []int
	for i := 0; i < len(graph); i++ {
		if v[i] != 0 {
			continue
		}
		v[i] = 1
		q = []int{i}
		for len(q) > 0 {
			i = q[0]
			q = q[1:]
			for _, j := range graph[i] {
				if v[j] != 0 && v[j] == v[i] {
					return false
				}
				if v[j] == 0 {
					v[j] = -v[i]
					q = append(q, j)
				}
			}
		}
	}
	return true

}
