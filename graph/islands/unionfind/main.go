package main

func find(i1 int, p []int) int {
	for p[i1] != i1 {
		i1 = p[i1]
	}
	return i1
}
func union(i1, i2 int, p []int) {
	p1, p2 := find(i1, p), find(i2, p)
	if p2 < p1 {
		p[p1] = p2
	} else {
		p[p2] = p1
	}
}
func numIslands(grid [][]byte) int {
	// [[1],[1]]
	// m = 2 n=1
	m, n := len(grid), len(grid[0])
	p := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		if grid[i/n][i%n] == '1' {
			p[i] = i
		} else {
			p[i] = -1
		}
	}
	// p = [0,1]
	for i := 0; i < m*n; i++ {
		if p[i] == -1 {
			continue
		}
		if i >= n && p[i-n] != -1 {
			union(i, i-n, p)
		}
		if i%n != 0 && p[i-1] != -1 {
			union(i, i-1, p)
		}
	}
	cnt := 0
	for i := 0; i < m*n; i++ {
		if p[i] == i {
			cnt++
		}
	}
	return cnt
}

/*
[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
[  ["1","1","0","0","0"],  ["1","1","0","0","0"],["0","0","1","0","0"],  ["0","0","0","1","1"]]
[["1"],["1"]]
*/
