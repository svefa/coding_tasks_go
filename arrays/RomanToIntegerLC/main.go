package main

func romanToInt(s string) int {
	m := map[rune]int{'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}
	// result, prev, next
	r, p, n := 0, 0, 0
	for _, c := range s {
		n = m[c]
		r += n
		if p < n {
			r -= 2 * p
		}
	}
	return r
}

func main() {
	println(romanToInt("III")) //3
}
