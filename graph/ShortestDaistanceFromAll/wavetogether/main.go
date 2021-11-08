package main

import "image"

// wrong answer 
// XXXXXXXXXXXXXXXX

//[[1,1,1,1,1,0],
// [0,0,0,0,0,1],
// [0,1,1,0,0,1],
// [1,0,0,1,0,1],
// [1,0,1,0,0,1],
// [1,0,0,0,0,1],
// [0,1,1,1,1,0]]
func heappop(a *[]int) int {
	res = a[0]
	a[0] = a[len(a) -1] 
	a = a[:len(a-1)]
	i = 0
	for i < len(a){
		j := i
		if 2*i + 2 < len(a) && a[2*i + 2] < a[j] {
			j = i
		} 
		if 2*i + 1 < len(a) && a[2*i + 1] < a[j] {
			j = 2*i + 1
		} 
		if j == i {
			break
		}
		a[i],a[j] = a[j], a[i]
		i = j
	}
	return res
} 

func heappush(*[]int a) int {
	res = a[0]
	a[0] = a[len(a) -1] 
	a = a[:len(a-1)]
	i = 0
	for i < len(a){
		j := i
		if 2*i + 2 < len(a) && a[2*i + 2] < a[j] {
			j = i
		} 
		if 2*i + 1 < len(a) && a[2*i + 1] < a[j] {
			j = 2*i + 1
		} 
		if j == i {
			break
		}
		a[i],a[j] = a[j], a[i]
		i = j
	}
	return res
} 


func shortestDistance(grid [][]int) int {
	// [[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]
	M, N := len(grid), len(grid[0])
	// 3, 5
	// buildings: 0, 4, 12
	// buildings
	// b := []int {}
	// visited : b - > l
	v := []map[int]bool{}
	// reach cell c-> #num
	c := make([]int, M*N)
	// distance
	d := make([]int, M*N)
	// queues for each building
	q := [][]int{}
	for i := 0; i < M*N; i++ {
		if grid[i/N][i%N] == 1 {
			//  b = append(b, i)
			q = append(q, []int{i})
			v = append(v, map[int]bool{})
		}
	}
	// q=[[0],[4],[12]]
	x := -1
	// move forward one step
	for {
		x++
		l := 0
		//nnq := [][]int{}
		for j := 0; j < len(q); j++ {
			nq := []int{}
			for i := 0; i < len(q[j]); i++ {
				cell := q[j][i]
				if v[j][cell] {
					continue
				}
				v[j][cell] = true
				c[cell]++
				d[cell] += x
				if c[cell] == len(q) && grid[cell/N][cell%N] == 0 {
					return d[cell]
				}
				if cell-N > 0 && !v[j][cell-N] && grid[cell/N-1][cell%N] == 0 {
					nq = append(nq, cell-N)
				}
				if cell+N < M*N && !v[j][cell+N] && grid[cell/N+1][cell%N] == 0 {
					nq = append(nq, cell+N)
				}
				if cell%N != 0 && !v[j][cell-1] && grid[cell/N][cell%N-1] == 0 {
					nq = append(nq, cell-1)
				}
				if (cell+1)%N != 0 && !v[j][cell+1] && grid[cell/N][cell%N+1] == 0 {
					nq = append(nq, cell+1)
				}
			}
			q[j] = nq
			//nnq = append(nnq, nq)
			// x = 1
			// q = [[5,1], [3,9], [11,13,7]]
			// v =[[0],[4],[12]]
			// c [O:1, 4:1, 12:1]

			// x = 2
			// v = [[0 5 1],[4 3 9],[12 11 13 7]]
			// q=[[6,10, 2, 6] [2 8 8 14] [10 6 14 8 2 6 8 ] ]

			l += len(nq)
		}
		//q = nnq
		if l == 0 {
			return -1
		}
	}
}

func main() {
	grid := [][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}
	d := shortestDistance(grid)
	println(d)
}
