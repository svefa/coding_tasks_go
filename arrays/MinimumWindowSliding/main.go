package main

func minWindow(s string, t string) string {
	// s = 'ADOBECODEBANC'
	// t = 'ABC'
	r := ""
	mt := make(map[byte]int)
	ms := make(map[byte]int)
	nt := len(t)
	for i := range t {
		mt[t[i]]++
	}
	// mt {'A':1,'B':1,'C':1}
	// nt = 3
	j := 0
	for i := range s {
		ms[s[i]]++
		// { ms = {'A':1 } }
		if mt[s[i]] >= ms[s[i]] {
			nt--
		}
		// mt[A]=1  >= ms[A=1]
		// i=0 s[i]='A' nt=2
		// i=3 s[i]=B nt=1
		// i=5 s[i]=C nt=0
		if nt == 0 {
			for j <= i && nt == 0 {
				// j=0 s[j] = A
				ms[s[j]]--
				// ms={A:0}
				if mt[s[j]] > 0 && ms[s[j]] < mt[s[j]] {
					// 1 > 0  0 < 1
					if len(r) == 0 || i-j < len(r) {
						//5-0 r=s[0:6]=:ADOBEC
						r = s[j : i+1]
					}
					nt++
					// nt=1
					j++
					break
				}
				j++
			}
		}
	}
	return r

}
