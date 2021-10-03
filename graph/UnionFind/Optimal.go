// Union-find Constructor	Find	Union	Connected
//Time complexity	 O(N)	O(⍺(N))	O(⍺(N))	O(⍺(N))
//
//  N is the number of vertices in the graph.
// ⍺ refers to the Inverse Ackermann function.
// In practice, we assume it's a constant. In other words,
// O(⍺(N)) is regarded as O(1) on average.

// UNION RANK and COMPRESSED FIND

package main

func union(p []int, r []int, i1, i2 int) {
	p1 := find(p, i1)
	p2 := find(p, i2)
	if i1 == 5 && i2 == 6 {
		print_slice(p)
		print_slice(r)
		print(i1, i2, p1, p2, r[i1], r[i2], "\n")
	}

	if p1 != p2 {
		if r[p1] > r[p2] {
			p[i2] = i1
		} else if r[p1] < r[p2] {
			p[i1] = i2
		} else {
			p[i1] = i2
			r[i2] += 1
		}
	}
}
func find(p []int, i int) int {
	if p[i] != i {
		p[i] = find(p, p[i])
	}
	return p[i]
}

func connected(p []int, i, j int) bool {
	return find(p, i) == find(p, j)
}

/**************************************************/
func range_n(n int) ([]int, []int) {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return p, r
}
func print_slice(p []int) {
	for i := 0; i < len(p); i++ {
		print(p[i])
	}
	print("\n")

}

/******************************************************/

func main() {
	n := 10
	p, r := range_n(n)
	// 1-2-5-6-7 3-8-9 4
	union(p, r, 1, 2)
	union(p, r, 2, 5)
	print_slice(p)
	print_slice(r)
	union(p, r, 5, 6)
	print_slice(p)
	print_slice(r)

	union(p, r, 6, 7)
	union(p, r, 3, 8)
	union(p, r, 8, 9)
	print(connected(p, 1, 5), "\n") // true
	print(connected(p, 5, 7), "\n") // true
	print(connected(p, 4, 9), "\n") // false
	// 1-2-5-6-7 3-8-9-4
	union(p, r, 9, 4)
	print(connected(p, 4, 9), "\n") // true

	print_slice(p)
	print_slice(r)
}
