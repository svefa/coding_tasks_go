func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})
	for _, pip := range pipes {
		union(p, wells, pip[0]-1, pip[1]-1, pip[2])
	}
	s := 0
	for i := 0; i < n; i++ {
		s += wells[i]
	}
	return s
}

func union(p, r []int, i, j, v int) {
	pi, pj := find(p, i), find(p, j)
	if pi == pj {
		return
	}
	if v > r[pi] && v > r[pj] {
		return
	}
	if r[pi] < r[pj] {
		p[pj] = p[pi]
		r[pj] = v
	} else {
		p[pi] = p[pj]
		r[pi] = v
	}
}

func find(p []int, i int) int {
	if i != p[i] {
		i = find(p, p[i])
	}
	return i
}