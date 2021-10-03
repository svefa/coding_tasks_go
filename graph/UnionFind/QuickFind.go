// Union-find Constructor	Find	Union	Connected
// Time complexity	O(N)	O(1)	O(N)	O(1)

package mainqf

import (
	"fmt"
)

type UF struct {
	root []int
}

// O(N)
//  connect n elementov of tree overall O(N*N)
func (uf *UF) Union(x int, y int) {
	rtx := uf.Find(x)
	rty := uf.Find(y)
	if rtx != rty {
		for i := 0; i < len(uf.root); i++ {
			if uf.root[i] == rty {
				uf.root[i] = rtx
			}
		}
	}
}

// O(1)
func (uf *UF) Find(x int) int {
	return uf.root[x]
}

// O(1)
func (uf *UF) Connected(x int, y int) bool {
	return uf.root[x] == uf.root[y]
}

func InitUF(n int) *UF {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return &UF{root}
}

/**************************************************/

func test() {
	// 1-2-5-6-7 3-8-9 4
	uf := InitUF(10)
	uf.Union(1, 2)
	uf.Union(2, 5)
	uf.Union(7, 6)
	uf.Union(6, 5)

	uf.Union(3, 8)
	uf.Union(8, 9)

	fmt.Println(uf.Connected(1, 5)) // true
	fmt.Println(uf.Connected(5, 7)) // true
	fmt.Println(uf.Connected(4, 9)) // false

	//1-2-5-6-7 3-8-9-4
	uf.Union(9, 4)
	fmt.Println(uf.Connected(4, 9)) // true

}
