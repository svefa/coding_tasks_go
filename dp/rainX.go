package main

func trapX(height []int) int {
	var total int
	right := 0
	left := 0
	start, end := 0, len(height)-1
	for start < end {
		lh := height[start]
		rh := height[end]
		if lh < rh {
			if lh > left {
				left = lh
			} else {
				total += left - lh
			}
			start++
		} else {
			if rh > right {
				right = rh
			} else {
				total += right - rh
			}
			end--
		}
	}
	return total
}

/*
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	s := trapX(height) // 6
	fmt.Println(s)
}
*/
