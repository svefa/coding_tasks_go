package main

/* 1631. Path With Minimum Effort
“The Shortest Path Faster Algorithm” (SPFA algorithm).
Graph wit negative edges, but no negative cycles.
Improved variation of the Bellman-Ford algorithm by using a queue.

Time complecity:  O(V*E)
Space complexity: O(V)
*/

func minimumEffortPath(heights [][]int) int {
	//[[1,2,2],
	//[3,8,2],
	//[5,3,5]]

	//  0 1 1  012     q = 0 13 3042 0426 42613
	//  2   1  345
	//  2 2 2  678
	r, c := len(heights), len(heights[0])
	// h = [[1,10,6]]
	// tmp = [0 inf]
	temp := make([][]int, r)
	for i := 0; i < r; i++ {
		temp[i] = make([]int, c)
		for j := 0; j < c; j++ {
			temp[i][j] = 1 << 32
		}
	}
	temp[0][0] = 0

	// queue and visited for shortes path faster algo
	q := []int{0}
	v := make([]bool, c*r)
	v[q[0]] = true

	for len(q) > 0 {
		//fmt.Println(v, q, temp)
		v[q[0]] = false
		i, j := q[0]/c, q[0]%c
		q = q[1:]
		//for i := 0; i < r; i++ {
		//	for j := 0; j < c; j++ {
		// i,j,h,t = 0,0,1,0   i,j,h,t = 0,1,10,9 i,j,h,t = 0,2,6,inf
		h := heights[i][j]
		t := temp[i][j]
		// h=1 h = 10 h= 6
		for _, ij := range [][]int{{i - 1, j}, {i, j - 1}, {i, j + 1}, {i + 1, j}} {
			x, y := ij[0], ij[1]
			// x,y= 0,1
			// x,y = 0,0  0,2 h = 10 t = 9
			if 0 <= x && x < r && 0 <= y && y < c {
				// dh = 9
				// dh = 9, 4
				dh := h - heights[x][y]
				if dh < 0 {
					dh = -dh
				}
				if dh < t {
					dh = t
				}
				// dh = 9  9
				// add to queue only updated
				if temp[x][y] > dh {
					temp[x][y] = dh
					if !v[x*c+y] {
						q = append(q, x*c+y)
						v[x*c+y] = true
					}
				}
				// temp = [[0,9,inf ...]]
				// temp =[0,9,9]
				// q =  0, 2

				// temp = 0,9
			}

		}
		//	}
		//}
	}

	return temp[r-1][c-1]

}

func main() {
	heights := [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	z := minimumEffortPath(heights)
	println(z)
}
