package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	j, l := 0, 0
	for i, ch := range s {
		if m[ch] > j {
			j = m[ch] + 1
		}
		if l < i-j+1 {
			l = i - j
		}
		m[ch] = i
	}
	return l
}
func main() {
	l := lengthOfLongestSubstring("ttmmrhbvmt")
	println(l)
}
