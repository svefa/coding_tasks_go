package main

import "fmt"

func networkDelayTime(times [][]int, n int, k int) int {
	g := make([]map[int]int, n)
	for i, _ := range g {
		g[i] = make(map[int]int)
	}

	v := make([]bool, n)
	w := make([]int, n)
	for _, t := range times {
		g[t[0]-1][t[1]-1] = t[2]
		//g[t[1]-1][t[0]-1] = t[2]
	}
	for i, _ := range w {
		w[i] = 1 << 32
	}

	w[k-1] = 0
	ti := 0
	q := [][]int{{k - 1, 0}}
	for len(q) > 0 {
		var i, wei int
		q, i, wei = heappop(q)
		if !v[i] {
			v[i] = true
			if wei > ti {
				ti = wei
			}
			for j, t := range g[i] {
				if !v[j] {
					if wei+t < w[j] {
						w[j] = wei + t
						q = heappush(q, j, w[j])
					}
				}
			}
		}
	}

	for i, _ := range w {
		if w[i] == 1<<32 {
			return -1
		}
	}

	return ti
}

/***************/
//  [ind, weight] -element
func heappush(q [][]int, j, wei int) [][]int {
	q = append(q, []int{j, wei})

	for i := len(q) - 1; i > 0; i = (i - 1) / 2 {
		if q[i][1] > q[(i-1)/2][1] {
			break
		}
		q[i], q[(i-1)/2] = q[(i-1)/2], q[i]
	}
	return q
}

func heappop(q [][]int) ([][]int, int, int) {
	fmt.Println("heappop start:", q)
	n := len(q)
	q[0], q[n-1] = q[n-1], q[0]

	for i := 0; i < len(q)-1; {
		fmt.Println("heappop swap:", q)
		if 2*i+2 < len(q)-1 {
			if q[i][1] <= q[2*i+2][1] && q[i][1] <= q[2*i+1][1] {
				break
			}
			if q[2*i+2][1] > q[2*i+1][1] {
				q[i], q[2*i+1] = q[2*i+1], q[i]
				i = 2*i + 1
			} else {
				q[i], q[2*i+2] = q[2*i+2], q[i]
				i = 2*i + 2
			}
		} else if 2*i+2 < len(q)-1 {
			if q[i][1] <= q[2*i+1][1] {
				break
			}
			q[i], q[2*i+1] = q[2*i+1], q[i]
			i = 2*i + 1
		} else {
			break
		}

	}
	fmt.Println(q[n-1][0], q[n-1][1], "heappop end:", q)
	return q[:n-1], q[n-1][0], q[n-1][1]
}

/****************

func heappoppush(q [][]int, e []int) []int {
    e, q[0] = q[0], e
    for i := 0; i < len(q); {
        if 2*i+2 < len(q) {
            if q[i] <= q[2*i+2] && q[i] <= q[2*i+1] {
                return
            }
            if q[2*i+2] < q[2*i+1] {
                q[i], q[2*i+1] = q[2*i+1], q[i]
                i = 2*i+1
            } else {
                q[i], q[2*i+2] = q[2*i+2], q[i]
                i = 2*i+2
            }
        } else if 2*i+2 < len(q) {
            if q[i] <= q[2*i+1] {
                return
            }
            q[i],  q[2*i+1] = q[2*i+1], q[i]
            i = 2*i+1
        }
    }
    return e
}
***/

func main() {
	test3()
}

/*****************************************/
func test1() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	t := networkDelayTime(times, n, k)
	println("res:", t)

}
func test2() {
	times := [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}
	n := 3
	k := 1
	t := networkDelayTime(times, n, k)
	println("res:", t)
}

func test3() {
	times := [][]int{{1, 5, 66}, {3, 5, 55}, {4, 3, 29}, {1, 2, 9}, {3, 4, 10}, {3, 1, 3}, {2, 3, 78}, {1, 4, 98}, {4, 5, 21}, {5, 2, 19}, {5, 1, 76}, {4, 1, 65}, {3, 2, 27}, {5, 3, 23}, {5, 4, 12}, {2, 1, 36}, {4, 2, 75}, {2, 4, 11}, {1, 3, 30}, {2, 5, 8}}
	n := 5
	k := 1
	t := networkDelayTime(times, n, k)
	println("res:", t)
}
