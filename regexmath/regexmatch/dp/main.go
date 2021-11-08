package main

import "fmt"

func isMatch(s string, p string) bool {
	// "ab"
	px := []string{}
	j := len(p) - 1
	for j >= 0 {
		if p[j] == '*' {
			px = append([]string{p[j:j+1] + p[j-1:j]}, px...)
			j -= 2
		} else {
			px = append([]string{p[j : j+1]}, px...)
			j--
		}
	}
	n, m := len(s), len(px)
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for j := 1; j < m+1; j++ {
		if px[j-1][0] == '*' && dp[0][j-1] {
			dp[0][j] = true
		} else if px[j-1] == "." && dp[0][j-1] {
			dp[0][j] = true
		} else {
			break
		}
	}
	// aa a*a
	//  n\m     ''    '*a'   '.'      'a'
	// ''      true   true   false
	// 'a'     false  true
	// 'a'     false

	// ab .*
	//        ""      "*."
	//   ''   true    true
	//   'a'  false   true
	//   'b'  false
	//
	//

	// aab c*a*b
	//        ""      "*c"    "*a"    "b"
	//   ''   true    true     true    false
	//   'a'  false   false    true    false
	//   'a'  false   false    true         true
	//   'b'  false   false
	//

	fmt.Println(s, px)
	for k, _ := range dp {
		fmt.Println(dp[k])
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Println(i, j, (s[i-1 : i]), px[j-1])
			//for k, _ := range dp {

			//fmt.Println(dp[k])
			//}
			if px[j-1] == "*." {
				if dp[i-1][j] || dp[i-1][j-1] {
					dp[i][j] = true
				}
			} else if px[j-1][0] == '*' {
				fmt.Println("*", px[j-1], dp[i-1][j], dp[i-1][j-1], s[i-1], px[j-1][1])
				if dp[i][j-1] || (dp[i-1][j-1] && s[i-1] == px[j-1][1]) {
					dp[i][j] = true
				}

			} else if px[j-1] == "." {
				if dp[i-1][j-1] {
					dp[i][j] = true
				}
			} else {
				fmt.Println("else", s[i-1], px[j-1][0], dp[i-1][j-1])
				if s[i-1] == px[j-1][0] && dp[i-1][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[n][m]

}

func main() {
	//fmt.Println(isMatch("aa", "a*"))
	//fmt.Println(isMatch("aaa", "ab*ac*a"))
	fmt.Println(isMatch("aab", "c*a*b"))
	//fmt.Println(isMatch("ab", ".*.."))
	//fmt.Println(isMatch("aaa", "a*a"))
	//fmt.Println(isMatch("", "..a*"))
	//fmt.Println(isMatch("a", ".*..a*"))
}
