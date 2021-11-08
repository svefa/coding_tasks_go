package main

// slow because of no MEMO
import "fmt"

func isMatch(s string, p string) bool {
	fmt.Println(s, p, len(s), len(p))
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		fmt.Println("len(s) > 0 && len(p) == 0")
		return false
	}
	if len(s) == 0 && len(p) == 1 {
		return false
	}
	if len(s) == 0 && p[1] == '*' {
		return isMatch(s, p[2:])
	}
	if len(s) == 0 && len(p) > 0 {
		return false
	}
	if len(p) == 1 && p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}
	if len(p) == 1 && p[0] != '.' {
		return p[0] == s[0] && isMatch(s[1:], p[1:])
	}
	if p[:2] == ".*" {
		if isMatch(s[1:], p) {
			return true
		}
		if isMatch(s, p[2:]) {
			return true
		}
		return false
	}
	if p[1] == '*' {
		if isMatch(s, p[2:]) {
			return true
		}
		if s[0] == p[0] && isMatch(s[1:], p) {
			return true
		}
		return false
	}
	if p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}
	return p[0] == s[0] && isMatch(s[1:], p[1:])
}

func main() {
	fmt.Println(isMatch("ab", ".*.."))
	//fmt.Println(isMatch("aaa", "a*a"))
	//fmt.Println(isMatch("", "..a*"))
	//fmt.Println(isMatch("a", ".*..a*"))
}
