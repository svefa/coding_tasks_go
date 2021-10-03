// lc 547
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	p := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
	}
	count := n

	// conn: [[1,1,0],[1,1,0],[0,0,1]]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				// 0,1
				union(p, r, i, j, &count)

			}
		}
	}

	return count
}

func find(p []int, i int) int {
	if p[i] != i {
		// opt
		p[i] = find(p, p[i])
	}
	return p[i]
}

func union(p, r []int, i, j int, count *int) {
	//0,1
	pi := find(p, i)
	pj := find(p, j)
	// pi, pj = 0,1
	if pi != pj {
		if r[pi] < r[pj] {
			p[pi] = p[j]
		} else if r[pi] > r[pj] {
			p[pj] = pi
		} else {
			p[pj] = pi
			r[pi]++
		}
		*count--
	}
}

