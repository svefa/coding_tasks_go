package main

func longestValidParentheses(s string) int {
	uf := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		uf[i] = -1
	}
	st := []int{}
	res := 0
	//  0 1 2 3 4 5
	// ") ( ) ( ) )"
	// ......
	// i  res    stack  uf
	// 1        [1]
	// 2    2   []     [,,1, ,,]
	// 3        [3]
	// 4    4   []   [O:,1:,2:1,3:,4:1,5:]
	// 5    ...
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			st = append(st, i)
		}

		if s[i] == ')' && len(st) > 0 {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			for j > 0 && uf[j-1] != -1 {
				j = uf[j-1]
			}
			uf[i] = j
			if res < i-j+1 {
				res = i - j + 1
			}
		}
	}
	return res
}
