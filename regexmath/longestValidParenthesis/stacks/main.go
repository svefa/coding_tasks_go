func longestValidParentheses(s string) int {
	// intervals ()
	res := 0
	stack := []int{}
	intervals := [][]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			// stack = [0]
		}
		if s[i] == ')' && len(stack) > 0 {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// j = 0 stack=[]
			intervals = append(intervals, []int{j, i})
		}
	}
	if len(intervals) == 0 {
		return 0
	}

	// intervals as stack
	for len(intervals) > 1 {
		i, j := intervals[len(intervals)-1][0], intervals[len(intervals)-1][1]
		intervals = intervals[:len(intervals)-1]
		if res < j-i+1 {
			res = j - i + 1
		}

		if intervals[len(intervals)-1][0] > i {
			// eat
			intervals[len(intervals)-1][0], intervals[len(intervals)-1][1] = i, j
		} else if intervals[len(intervals)-1][1] == i-1 {
			// merge
			intervals[len(intervals)-1][1] = j
		}
	}
	i, j := intervals[0][0], intervals[0][1]
	if res < j-i+1 {
		res = j - i + 1
	}
	return res
}