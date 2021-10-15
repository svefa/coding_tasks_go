package main

func validPalindrome(s string) bool {
	//ax....yb
	var i, j, i1, j1 int
	for i, j = 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			break
		}
	}
	if i >= j {
		return true
	}
	for i1, j1 = i, j-1; i1 < j1; i1, j1 = i1+1, j1-1 {
		if s[i1] != s[j1] {
			break
		}
	}
	if i1 >= j1 {
		return true
	}
	for i1, j1 = i+1, j; i1 < j1; i1, j1 = i1+1, j1-1 {
		if s[i1] != s[j1] {
			break
		}
	}
	if i1 >= j1 {
		return true
	}
	return false
}
