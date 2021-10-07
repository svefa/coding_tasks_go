package main

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
	i := k - 1
	wei := 0
	for wei < (1 << 32) {
		if !v[i] {
			v[i] = true
			if wei > ti {
				ti = wei
			}
			for j, t := range g[i] {
				if !v[j] {
					if wei+t < w[j] {
						w[j] = wei + t
					}
				}
			}
			wei = 1 << 32
			for k, _ := range w {
				if v[k] == false && w[k] < wei {
					i = k
					wei = w[k]
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
