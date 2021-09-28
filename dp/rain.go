package main

func trap(height []int) int {
	n := len(height)
	l := make([]int, n)
	r := make([]int, n)
	m := 0
	for i, h := range height {
		m = max(h, m)
		l[i] = m
	}
	m = 0
	for i := n - 1; i >= 0; i-- {
		m = max(height[i], m)
		r[i] = m
	}
	s := 0
	for i := 0; i < n; i++ {
		s += min(l[i], r[i]) - height[i]
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	s := trap(height) // 6
	fmt.Println(s)
}
*/
