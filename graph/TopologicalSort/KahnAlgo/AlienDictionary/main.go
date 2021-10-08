package main

func alienOrder(words []string) string {
	// words = ["wrt","wrf","er","ett","rftt"]
	// word ["abc","ab"]
	in := map[byte]int{}
	g := map[byte][]byte{}
	q := []byte{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			in[words[i][j]] = 0
		}
	}
	n := len(in)
	for i := 1; i < len(words); i++ {
		w1, w2 := words[i-1], words[i]
		for j := 0; j < len(w1) && j < len(w2); j++ {
			a, b := w1[j], w2[j]
			if a != b {
				g[a] = append(g[a], b)
				in[b]++
				break
			}
			if (j == len(w1)-1 || j == len(w2)-1) && len(w1) > len(w2) {
				return ""
			}
		}
	}
	// g[w] = [e] g[t] =[f] g[r]=[t] g[e]=[r]
	// in = {w: 0, e:1, f:1, r:1, t:1}
	// w -> e -> r ->t -f
	for b, _ := range in {
		if in[b] == 0 {
			q = append(q, b)
		}
	}
	// q=[w,]
	i := 0
	for i < len(q) {
		// i=0 q=[w]
		a := q[i]
		// i =0 q = [w] a=w g[w]= [e]
		//. i=1  q= [w,e] a=e g[e] = [r]
		// i=2  q= [w,e,r] a=r g[r] = [t]
		// i=3  q= [w,e,r,t] a=t g[t] = [f]
		// i=4  q= [w,e,r,t,f] a=f g[f] = []
		for _, b := range g[a] {
			in[b]--
			if in[b] == 0 {
				// in {e: 0, r:0, t:0, f:0}
				q = append(q, b)
				// q =[w,e,r,t,f]
			}
		}
		i++
	}
	if len(q) == n {
		return string(q)
	}
	return ""
}
