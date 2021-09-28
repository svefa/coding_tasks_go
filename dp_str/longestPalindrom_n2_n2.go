package main

// babad
// 1
// 01
// 001
// 0001
// 00001
func longestPalindrome_n2_n2(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	l, r := 0, 0

	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if i > 0 && s[i-1] == s[i] {
			dp[i-1][i] = true
			l = i - 1
			r = i
		}
	}
	// i=2, j= 4
	// dp [3][3] true; s[2] != s[4] (2,4)
	for i := 2; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if dp[j-i+1][j-1] && s[j-i] == s[j] {
				dp[j-i][j] = true
				l = j - i
				r = j
			}
		}
	}
	return s[l : r+1]
}

/*
func main() {
	s := "babad"
	out := longestPalindrome(s)
	fmt.Println(out)
}
*/
