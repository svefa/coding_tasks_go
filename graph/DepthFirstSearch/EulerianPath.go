package main

import "fmt"

func findItinerary(tickets [][]string) []string {
	g := map[string][]string{} // graph
	d := map[string]int{}      // degree
	u := map[string]int{}      // used
	p := []string{}            // path

	for _, t := range tickets {
		g[t[0]] = append(g[t[0]], t[1])
		d[t[1]]--
		d[t[0]]++
		u[t[0]]++
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
	var dfs func(string)
	dfs = func(from string) {
		for u[from] > 0 {
			u[from]--
			to := g[from][u[from]]
			dfs(to)
		}
		p = append(p, from)
	}
	dfs(c)
	return p
}

/************************************************/

func main() {
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	itinerary := findItinerary(tickets)
	fmt.Println(itinerary)
}
