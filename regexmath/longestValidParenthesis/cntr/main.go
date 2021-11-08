package main

func longestValidParentheses(s string) int {
	cnt := 0
	res := 0
	m := 0
	l, r := 0, 0
	for l < len(s) {
		r = l
		cnt = 0
		m = 0
		for cnt >= 0 && r < len(s) {
			if s[r] == '(' {
				cnt++
				m += 2
			}
			if s[r] == ')' {
				cnt--
			}
			if cnt == 0 && res < m {
				res = m
			}
			r++
		}
		l = r
	}
	l, r = len(s)-1, len(s)-1
	for r >= 0 {
		r = l
		cnt = 0
		m = 0
		for cnt >= 0 && l >= 0 {
			if s[l] == '(' {
				cnt--
				m += 2
			}
			if s[l] == ')' {
				cnt++
			}
			if cnt == 0 && res < m {
				res = m
			}
			l--
		}
		r = l
	}

	return res
}

func main() {
	///  ((xxxx)(www)(xxxx)
	s := ") ( ) ( ) )"
	// )  1 1 2 2 3 4
	// (  0 1 1 2 2 2
	n := longestValidParentheses(s)
	println(n)
}
