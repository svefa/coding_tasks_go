package main

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, w := range wordDict {
		if len(w) <= len(s) && w == s[:len(w)] && wordBreak(s[len(w):], wordDict) {
			return true
		}
	}
	return false
}

// TLE
// "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
// ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
