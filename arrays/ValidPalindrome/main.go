package main

import "strings"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	s = strings.ToLower(s)
	for i < j {
		for i < j && (('0' <= s[i] && s[i] <= '9') || ('a' <= s[i] && s[i] <= 'z')) == false {
			i++
		}
		for i < j && (('0' <= s[j] && s[j] <= '9') || ('a' <= s[j] && s[j] <= 'z')) == false {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
