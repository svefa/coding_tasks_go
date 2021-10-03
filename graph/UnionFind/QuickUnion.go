package mainqu

//	Union-find Constructor	Find	Union	Connected
// Time complexity	O(N)	O(H)	O(H)	O(H)
// O(N) - but faster than quick find
func union(p []int, x int, y int) {
	rx := find(p, x)
	ry := find(p, y)
	if rx != ry {
		p[rx] = ry
	}
}

// O(N) or O(H) -height
// connect n elementov of tree overall O(N lnN )
func find(p []int, x int) int {
	if x != p[x] {
		// optimization balance
		p[x] = find(p, p[x])
	}
	return p[x]
}

// O(N)
func connected(p []int, x, y int) bool {
	return find(p, x) == find(p, y)
}

func range_n(n int) []int {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return p
}

/*********************************************/
func main() {
	p := range_n(10)
	// 1-2-5-6-7 3-8-9 4
	union(p, 1, 2)
	union(p, 2, 5)
	union(p, 5, 6)
	union(p, 6, 7)
	union(p, 3, 8)
	union(p, 8, 9)
	print(connected(p, 1, 5), "\n") // true
	print(connected(p, 5, 7), "\n") // true
	print(connected(p, 4, 9), "\n") // false
	for i := 0; i < 10; i++ {
		print(p[i])
	}
	print("\n")
	// 1-2-5-6-7 3-8-9-4
	union(p, 9, 4)
	print(connected(p, 4, 9), "\n") // true
}
