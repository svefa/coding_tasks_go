package main

import "sort"

func union(i, j int, p []int) {
	pi, pj := find(i, p), find(j, p)
	p[pj] = pi
}

func find(i int, p []int) int {
	for i != p[i] {
		i = p[i]
	}
	return i
}

func accountsMerge(accounts [][]string) [][]string {
	p := make([]int, len(accounts))
	for i := 0; i < len(p); i++ {
		p[i] = i
	}

	m := map[string]int{}

	for j, a := range accounts {
		for i := 1; i < len(a); i++ {
			if _, ok := m[a[i]]; !ok {
				m[a[i]] = j

			} else {
				union(j, m[a[i]], p)
			}
		}
	}

	r := map[int][]string{}
	for i := 0; i < len(p); i++ {
		p[i] = find(i, p)
		if p[i] == i {
			r[i] = []string{}
		}
	}

	for e, i := range m {
		r[p[i]] = append(r[p[i]], e)
	}
	res := [][]string{}
	for i, es := range r {
		sort.Strings(es)
		res = append(res, append([]string{accounts[i][0]}, es...))
	}
	return res

}
