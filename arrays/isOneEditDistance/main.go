package main

func isOneEditDistance(s string, t string) bool {
	// s="ab" t="acb"
	m, n := len(s), len(t)
	// m=2 n=3
	i, j := 0, 0
	for i < m && j < n && s[i] == t[j] {
		i++
		j++
	}
	// i=1 j=1
	println(i, j)
	if i == n && j == m {
		return false
	}

	if m > n {
		i++
	} else if m < n {
		j++
	} else {
		i++
		j++
	}
	// i=1; j=2
	println(i, j)
	for i < m && j < n && s[i] == t[j] {
		i++
		j++
	}
	// i =2; j=3
	println(i, j)
	return (i == m && j == n)
}
func main() {
	x := isOneEditDistance("ab", "acb")
	println(x)
}
