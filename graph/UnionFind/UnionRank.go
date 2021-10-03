//Union-find Constructor	Find	Union	Connected
//Time complexity	O(N)	O(logN)	O(logN)	O(logN)

package main

//parent references and rank of height of tree making more balanced
func union(p []int, x1, x2 int) {
	p1, r1 := find(p, x1)
	p2, r2 := find(p, x2)
	if p1 != p2 {
		if r1 > r2 {
			p[p2] = p1
		} else {
			p[p1] = p2
		}

	}
}

// ox - origin x; rx - path to origin
func find(p []int, x int) (int, int) {
	r:=0
	for x != p[x] {
		x = p[x]
		r++
	}
	return (p, r)
}


func connected(p []int, x, y int) bool {
	px, _ := find(p, x)
	py, _ := find(p, y)
	return px == py
}

func range_n(n int) []int {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return p
}

/*********************************************************/

func main() {
	n := 10
	p := range_n(n)
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
	// 1-2-5-6-7 3-8-9-4
	union(p, 9, 4)
	print(connected(p, 4, 9), "\n") // true
	for i := 0; i < n; i++ {
		print(p[i])
	}
}
