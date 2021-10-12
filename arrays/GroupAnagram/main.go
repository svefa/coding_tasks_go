package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]int)
	for i := 0; i < len(strs); i++ {
		key := []byte(strs[i])
		sort.Slice(key, func(a, b int) bool { return key[a] < key[b] })
		m[string(key)] = append(m[string(key)], i)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		s := make([]string, 0, len(v))
		for _, j := range v {
			s = append(s, strs[j])
		}
		res = append(res, s)
	}
	return res
}
