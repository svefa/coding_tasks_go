package main

// works

func wordBreak(s string, wordDict []string) bool {
	m := map[int]bool{}
	m[0] = true
	for i := 0; i < len(s); i++ {
		if !m[i] {
			continue
		}
		for _, w := range wordDict {
			if i+len(w) <= len(s) && !m[i+len(w)] && w == s[i:i+len(w)] {
				m[i+len(w)] = true
			}
		}
	}
	return m[len(s)]
}
