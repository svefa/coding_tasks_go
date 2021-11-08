package main

// TLE
func longestValidParentheses(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n && cnt >= 0; j++ {
			if s[j] == '(' {
				cnt++
			}
			if s[j] == ')' {
				cnt--
				if cnt == 0 && res < j-i+1 {
					res = j - i + 1
				}
			}
		}
	}
	return res
	// ")()())"
	// expand (....) :  ()...... : ...()
	// ----____
	// ")()())"

}
