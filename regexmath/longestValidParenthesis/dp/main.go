package main

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	res := 0
	//  0 1 2 3 4 5
	// ") ( ) ( ) )"
	// dp 0 0 2 0 2+2 0
	// ......
	// i  res    stack  uf
	// 1        [1]
	// 2    2   []     [,,1, ,,]
	// 3        [3]
	// 4    4   []   [O:,1:,2:1,3:,4:1,5:]
	// 5    ...
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if i-1 >= 0 && s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]

				}
				if res < dp[i] {
					res = dp[i]
				}
			}
			if i-1 >= 0 && s[i-1] == ')' {
				j := i - 1 - dp[i-1]
				if j >= 0 && s[j] == '(' {
					dp[i] = dp[i-1] + 2
					if j-1 >= 0 {
						dp[i] += dp[j-1]
					}
					if res < dp[i] {
						res = dp[i]
					}
				}
			}
		}
		//fmt.Println(i,res, st, uf)
	}
	return res
}
