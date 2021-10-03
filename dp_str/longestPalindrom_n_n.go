package main

// FAILED !!!

// longestPalindrome_n_n implements Manacher's algorithm
func longestPalindrome_n_n(s string) string {
	f := func(odd bool) string {
		n := len(s)
		radi := make([]int, n)
		i := 0
		c := 0
		r, j := 0, 0

		if odd {
			j = 1
		}
		for c+j < n {
			for c-(r+1) > 0 && c+j+(r+1) < n && s[c-(r+1)] == s[c+j+(r+1)] {
				r++
			}

			radi[c] = r
			if radi[i] < r {
				i = c
			}
			c++
			center := c
			radius := r

			for c+j < n && c < center+radius {
				mc := center - (c - center) + j
				mr := center + radius - c
				if radi[mc] < mr {
					radi[c+j] = radi[mc]
					c++
				} else if radi[mc] > mr {
					radi[c+j] = mr
					c++
				} else {
					r = mr
					break
				}
			}
		}
		return s[i-radi[i] : i+radi[i]+1-j]
	}
	s1 := f(true)
	s2 := f(false)
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
