package mainT

// tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	g := map[string][]string{} // graph
	d := map[string]int{}      // degree
	u := map[string]bool{}     // used
	p := []string{}            // path

	for _, t := range tickets {
		g[t[0]] = append(g[t[0]], t[1])
		d[t[1]]--
		d[t[0]]++
		u[t[0]+t[1]] = false
		// g = {muc:[lhr], jfk:[muc], sfo:[sjc], lhr[sfo]
		// d={muc:0, lhr: 0, jfk: +1, sfo:0, sjc: -1}
		// u = {m_l_: false, j_m_:false, s_s:false, l_s:false}
	}
	// n= 4

	var c string
	for k, _ := range d {
		if d[k] == 1 || c == "" {
			c = k
		}
	}
	// c= jfk p=[]
	p = append(p, c)
	var dfs func(c string) bool
	dfs = func(from string) bool {
		if len(p) == n+1 {
			return true
		}
		for _, to := range g[from] {
			if !u[from+to] {
				u[from+to] = true
				p = append(p, to)
				// u={j_m, m_l, l_s, s_s} , p=[jfk,muc, lhr, sfo, sjc]
				if dfs(to) {
					return true
				}
				p = p[0 : len(p)-1]
				u[from+to] = false
			}
		}
		return false
	}
	dfs(c)
	return p
}
