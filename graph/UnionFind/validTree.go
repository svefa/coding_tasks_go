type DisjointUnion struct {
	root []int
}

func constructor(n int) DisjointUnion {
	du := DisjointUnion{make([]int, n)}
	for i := 0; i < n; i++ {
		du.root[i] = i
	}
	return du
}

func (du DisjointUnion) find(x int) int {
	for x != du.root[x] {
		x = du.root[x]
	}
	return x
}

func (du DisjointUnion) union(x, y int) bool {
	rootX := du.find(x)
	rootY := du.find(y)

	if rootX == rootY {
		return false
	}

	du.root[rootY] = rootX
	return true
}

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	du := constructor(n)
	for _, edge := range edges {
		if !du.union(edge[0], edge[1]) {
			return false
		}
	}

	return true
}