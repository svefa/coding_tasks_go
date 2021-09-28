package main

// babad
// 1
// 01
// 001
// 0001
// 00001
func longestPalindrome_n2_c(s string) string {
	n := len(s)
	o := ""

	for i := 0; i < n; i++ {
		// odd
		l := i
		r := i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(o) {
			o = s[l+1 : r]
		}
		// even
		l = i
		r = i + 1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(o) {
			o = s[l+1 : r]
		}
	}
	return o
}

/*
func main() {
	s := "babad"
	out := longestPalindrome(s)
	fmt.Println(out)
}
*/
