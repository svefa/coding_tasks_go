package main

import (
	"fmt"
	"sort"
)

func rec(s string, res *[]string) {
	cnt := 0
	tmp := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			tmp = append(tmp, i)
			cnt--
			if cnt == -1 {
				for _, j := range tmp {
					if j == 0 || s[j] != s[j-1] {
						rec(s[:j]+s[j+1:], res)
					}
				}
				return
			}
		}
	}
	cnt = 0
	tmp = tmp[:0]
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			tmp = append(tmp, i)
			cnt--
			if cnt == -1 {
				for _, j := range tmp {
					if j == len(s)-1 || s[j] != s[j+1] {
						rec(s[:j]+s[j+1:], res)
					}
				}
				return
			}

		}
		if s[i] == ')' {
			cnt++
		}
	}
	*res = append(*res, s)

}
func removeInvalidParentheses(s string) (res []string) {
	x := []string{}
	rec(s, &x)
	sort.Strings(x)
	for i := 0; i < len(x); i++ {
		if len(res) == 0 || res[len(res)-1] != x[i] {
			res = append(res, x[i])
		}
	}
	return
}

func main() {
	// s := "("
	s := "(()"
	res := removeInvalidParentheses(s)
	fmt.Println(res)

}
