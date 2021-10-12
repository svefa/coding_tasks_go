package main

func addBinary(a string, b string) string {
	r := ""
	c := 0
	la, lb := len(a), len(b)
	i := 0
	for i = 1; i <= la && i <= lb; i++ {
		s := c
		if a[la-i] == '1' {
			s++
		}
		if b[lb-i] == '1' {
			s++
		}
		if s%2 == 0 {
			r = "0" + r
		} else {
			r = "1" + r
		}
		c = 0
		if s/2 == 1 {
			c = 1
		}
	}
	for ; i <= la; i++ {
		s := c
		if a[la-i] == '1' {
			s++
		}
		if s%2 == 0 {
			r = "0" + r
		} else {
			r = "1" + r
		}
		c = 0
		if s/2 == 1 {
			c = 1
		}
	}
	for ; i <= lb; i++ {
		s := c
		if b[lb-i] == '1' {
			s++
		}
		if s%2 == 0 {
			r = "0" + r
		} else {
			r = "1" + r
		}
		c = 0
		if s/2 == 1 {
			c = 1
		}
	}
	if c == 1 {
		r = "1" + r
	}
	if r[0] == '0' && len(r) > 1 {
		r = r[1:]
	}
	return r

}
